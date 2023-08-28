package games

import (
	"context"
	"fmt"

	"github.com/go-xorm/builder"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/logger"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	database "github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
	time_utils "github.com/nikita5637/quiz-registrator-api/utils/time"
)

// AddGames ...
func (f *Facade) AddGames(ctx context.Context, games []model.Game) error {
	leagues := make([]int32, 0)
	for _, game := range games {
		leagues = append(leagues, game.LeagueID)
	}

	err := f.db.RunTX(ctx, "Add games", func(ctx context.Context) error {
		var err error
		var notDeletedDBGames []database.Game
		notDeletedDBGames, err = f.gameStorage.Find(ctx, builder.NewCond().And(
			builder.In(
				"league_id", leagues,
			),
			builder.IsNull{
				"deleted_at",
			},
		), "")
		if err != nil {
			return fmt.Errorf("add games error: %w", err)
		}

		activeModelGames := make([]model.Game, 0)
		for _, notDeletedDBGame := range notDeletedDBGames {
			notDeletedModelGame := convertDBGameToModelGame(notDeletedDBGame)
			if notDeletedModelGame.IsActive() {
				activeModelGames = append(activeModelGames, notDeletedModelGame)
			}
		}

		for i, game := range games {
			var dbGames []database.Game
			dbGames, err = f.gameStorage.Find(ctx, builder.NewCond().And(
				builder.Eq{
					"league_id": game.LeagueID,
					"type":      game.Type,
					"number":    game.Number,
					"place_id":  game.PlaceID,
					"date":      game.DateTime().AsTime(),
				},
				builder.IsNull{
					"deleted_at",
				},
			), "")
			if err != nil {
				return fmt.Errorf("add games error: %w", err)
			}

			if len(dbGames) == 0 {
				logger.InfoKV(ctx, "inserting new game", "fields", game)
				var gameID int
				gameID, err = f.gameStorage.Insert(ctx, convertModelGameToDBGame(game))
				if err != nil {
					return fmt.Errorf("add games error: %w", err)
				}

				games[i].ID = int32(gameID)
			} else {
				games[i].ID = int32(dbGames[0].ID)
			}
		}

		gameIDsForDelete := getGameIDsForDelete(games, activeModelGames)
		for _, id := range gameIDsForDelete {
			logger.InfoKV(ctx, "deleting game that not contains in master", "game ID", id)
			err = f.gameStorage.Delete(ctx, int(id))
			if err != nil {
				return fmt.Errorf("delete game error: %w", err)
			}
		}

		outdatedGames := make([]model.Game, 0)
		for _, notDeletedDBGame := range notDeletedDBGames {
			modelGame := convertDBGameToModelGame(notDeletedDBGame)
			if !modelGame.IsActive() {
				outdatedGames = append(outdatedGames, modelGame)
			}
		}

		for _, outdatedGame := range outdatedGames {
			logger.InfoKV(ctx, "deleting outdated game", "game ID", outdatedGame.ID)
			err := f.gameStorage.Delete(ctx, int(outdatedGame.ID))
			if err != nil {
				return fmt.Errorf("delete outdated game error: %w", err)
			}
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("add games error: %w", err)
	}

	return nil
}

func getGameIDsForDelete(games, activeGames []model.Game) []int32 {
	gameIDsForDelete := make([]int32, 0)
	for _, activeGame := range activeGames {
		// current time is in active lag time window after game datetime
		if time_utils.TimeNow().After(activeGame.Date.AsTime()) {
			continue
		}

		found := false

		for _, game := range games {
			if game.ID == activeGame.ID {
				found = true
				break
			}
		}

		if !found {
			gameIDsForDelete = append(gameIDsForDelete, activeGame.ID)
		}
	}

	return gameIDsForDelete
}
