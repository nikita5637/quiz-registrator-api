package games

import (
	"context"
	"fmt"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/go-xorm/builder"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/leagues"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/places"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
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
			)
		}

		dbGames, err := f.gameStorage.Find(ctx, builderCond, "")
		if err != nil {
			return fmt.Errorf("find error: %w", err)
		}

		if len(dbGames) > 0 {
			return ErrGameAlreadyExists
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

		return nil
	})
	if err != nil {
		return model.NewGame(), fmt.Errorf("PatchGame error: %w", err)
	}

	return modelGame, nil
}
