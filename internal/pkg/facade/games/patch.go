package games

import (
	"context"
	"fmt"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/go-xorm/builder"
	"github.com/mono83/maybe"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/leagues"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/places"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	quizlogger "github.com/nikita5637/quiz-registrator-api/internal/pkg/quiz_logger"
	usersutils "github.com/nikita5637/quiz-registrator-api/utils/users"
)

// PatchGame ...
func (f *Facade) PatchGame(ctx context.Context, modelGame model.Game) (model.Game, error) {
	err := f.db.RunTX(ctx, "PatchGame", func(ctx context.Context) error {
		var builderCond builder.Cond
		if externalID, isPresent := modelGame.ExternalID.Get(); isPresent {
			builderCond = builder.And(
				builder.Neq{
					"id": modelGame.ID,
				},
				builder.Eq{
					"external_id": externalID,
					"league_id":   modelGame.LeagueID,
					"place_id":    modelGame.PlaceID,
					"number":      modelGame.Number,
					"date":        modelGame.DateTime().AsTime(),
				},
				builder.IsNull{
					"deleted_at",
				},
			)
		} else {
			builderCond = builder.And(
				builder.Neq{
					"id": modelGame.ID,
				},
				builder.IsNull{
					"external_id",
				},
				builder.Eq{
					"league_id": modelGame.LeagueID,
					"place_id":  modelGame.PlaceID,
					"number":    modelGame.Number,
					"date":      modelGame.DateTime().AsTime(),
				},
				builder.IsNull{
					"deleted_at",
				},
			)
		}

		dbGames, err := f.gameStorage.Find(ctx, builderCond, "")
		if err != nil {
			return fmt.Errorf("find error: %w", err)
		}

		if len(dbGames) > 0 {
			return ErrGameAlreadyExists
		}

		originalDBGame, err := f.gameStorage.GetGameByID(ctx, int(modelGame.ID))
		if err != nil {
			return fmt.Errorf("getting game by game ID error: %w", err)
		}

		patchedDBGame := convertModelGameToDBGame(modelGame)
		if err := f.gameStorage.PatchGame(ctx, patchedDBGame); err != nil {
			if err, ok := err.(*mysql.MySQLError); ok {
				if err.Number == 1452 {
					if i := strings.Index(err.Message, leagueIBFK1ConstraintName); i != -1 {
						return fmt.Errorf("patch game error: %w", leagues.ErrLeagueNotFound)
					}

					return fmt.Errorf("patch game error: %w", places.ErrPlaceNotFound)
				}
			}

			return fmt.Errorf("patch game error: %w", err)
		}

		modelGame.HasPassed = gameHasPassed(modelGame)

		if gameLink := getGameLink(modelGame); gameLink != "" {
			modelGame.GameLink = maybe.Just(gameLink)
		}

		logs := make([]quizlogger.Params, 0)

		userID := maybe.Nothing[int32]()
		userFromContext := usersutils.UserFromContext(ctx)
		if userFromContext != nil {
			userID = maybe.Just(userFromContext.ID)
		}

		if originalDBGame.Registered != modelGame.Registered {
			messageID := quizlogger.GameRegistered
			if !modelGame.Registered {
				messageID = quizlogger.GameUnregistered
			}

			logs = append(logs, quizlogger.Params{
				UserID:     userID,
				ActionID:   quizlogger.UpdatingActionID,
				MessageID:  messageID,
				ObjectType: maybe.Just(quizlogger.ObjectTypeGame),
				ObjectID:   maybe.Just(modelGame.ID),
				Metadata:   nil,
			})
		}

		if model.Payment(originalDBGame.Payment.Int64) != modelGame.Payment.Value() {
			logs = append(logs, quizlogger.Params{
				UserID:     userID,
				ActionID:   quizlogger.UpdatingActionID,
				MessageID:  quizlogger.GamePaymentChanged,
				ObjectType: maybe.Just(quizlogger.ObjectTypeGame),
				ObjectID:   maybe.Just(modelGame.ID),
				Metadata: quizlogger.GamePaymentChangedMetadata{
					OldPayment: model.Payment(originalDBGame.Payment.Int64),
					NewPayment: modelGame.Payment.Value(),
				},
			})
		}

		if len(logs) > 0 {
			if err := f.quizLogger.WriteBatch(ctx, logs); err != nil {
				return fmt.Errorf("write logs error: %w", err)
			}
		}

		return nil
	})
	if err != nil {
		return model.NewGame(), fmt.Errorf("PatchGame error: %w", err)
	}

	return modelGame, nil
}
