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
func (f *Facade) PatchGame(ctx context.Context, game model.Game) (model.Game, error) {
	err := f.db.RunTX(ctx, "PatchGame", func(ctx context.Context) error {
		var builderCond builder.Cond
		if externalID, isPresent := game.ExternalID.Get(); isPresent {
			builderCond = builder.And(
				builder.Neq{
					"id": game.ID,
				},
				builder.Eq{
					"external_id": externalID,
					"league_id":   game.LeagueID,
					"place_id":    game.PlaceID,
					"number":      game.Number,
					"date":        game.DateTime().AsTime(),
				},
			)
		} else {
			builderCond = builder.And(
				builder.Neq{
					"id": game.ID,
				},
				builder.IsNull{
					"external_id",
				},
				builder.Eq{
					"league_id": game.LeagueID,
					"place_id":  game.PlaceID,
					"number":    game.Number,
					"date":      game.DateTime().AsTime(),
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

		patchedDBGame := convertModelGameToDBGame(game)
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

		game.HasPassed = !game.IsActive()

		return nil
	})
	if err != nil {
		return model.NewGame(), fmt.Errorf("PatchGame error: %w", err)
	}

	return game, nil
}
