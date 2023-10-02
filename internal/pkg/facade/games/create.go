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

// CreateGame ...
func (f *Facade) CreateGame(ctx context.Context, game model.Game) (model.Game, error) {
	createdModelGame := model.NewGame()
	err := f.db.RunTX(ctx, "CreateGame", func(ctx context.Context) error {
		var builderCond builder.Cond
		if externalID, isPresent := game.ExternalID.Get(); isPresent {
			builderCond = builder.And(
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

		newDBGame := convertModelGameToDBGame(game)
		id, err := f.gameStorage.CreateGame(ctx, newDBGame)
		if err != nil {
			if err, ok := err.(*mysql.MySQLError); ok {
				if err.Number == 1452 {
					if i := strings.Index(err.Message, leagueIBFK1ConstraintName); i != -1 {
						return fmt.Errorf("create game error: %w", leagues.ErrLeagueNotFound)
					}

					return fmt.Errorf("create game error: %w", places.ErrPlaceNotFound)
				}
			}

			return fmt.Errorf("create game error: %w", err)
		}

		newDBGame.ID = id
		createdModelGame = convertDBGameToModelGame(newDBGame)

		return nil
	})
	if err != nil {
		return model.NewGame(), fmt.Errorf("CreateGame error: %w", err)
	}

	return createdModelGame, nil
}
