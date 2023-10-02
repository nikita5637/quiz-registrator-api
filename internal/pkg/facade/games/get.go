package games

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/go-xorm/builder"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	timeutils "github.com/nikita5637/quiz-registrator-api/utils/time"
)

// GetGameByID ...
func (f *Facade) GetGameByID(ctx context.Context, id int32) (model.Game, error) {
	modelGame := model.NewGame()
	err := f.db.RunTX(ctx, "GetGameByID", func(context.Context) error {
		dbGame, err := f.gameStorage.GetGameByID(ctx, int(id))
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return fmt.Errorf("get game by ID error: %w", ErrGameNotFound)
			}

			return fmt.Errorf("get game by ID error: %w", err)
		}

		if dbGame.DeletedAt.Valid {
			return ErrGameNotFound
		}

		modelGame = convertDBGameToModelGame(*dbGame)
		modelGame.HasPassed = gameHasPassed(modelGame)

		return nil
	})
	if err != nil {
		return model.NewGame(), fmt.Errorf("GetGameByID error: %w", err)
	}

	return modelGame, nil
}

// GetGamesByIDs ...
func (f *Facade) GetGamesByIDs(ctx context.Context, ids []int32) ([]model.Game, error) {
	var modelGames []model.Game
	err := f.db.RunTX(ctx, "GetGameByID", func(context.Context) error {
		dbGames, err := f.gameStorage.Find(ctx, builder.And(
			builder.In("id", ids),
			builder.IsNull{
				"deleted_at",
			},
		), "")
		if err != nil {
			return fmt.Errorf("get games by IDs error: %w", err)
		}

		modelGames = make([]model.Game, 0, len(dbGames))
		for _, dbGame := range dbGames {
			modelGame := convertDBGameToModelGame(dbGame)
			modelGame.HasPassed = gameHasPassed(modelGame)

			modelGames = append(modelGames, modelGame)
		}

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("GetGameByID error: %w", err)
	}

	return modelGames, nil
}

// GetTodaysGames ...
func (f *Facade) GetTodaysGames(ctx context.Context) ([]model.Game, error) {
	var modelGames []model.Game
	timeNow := timeutils.TimeNow()
	dateExpr := fmt.Sprintf("date LIKE \"%04d-%02d-%02d%%\"", timeNow.Year(), timeNow.Month(), timeNow.Day())
	err := f.db.RunTX(ctx, "GetTodaysGames", func(ctx context.Context) error {
		dbGames, err := f.gameStorage.Find(ctx, builder.NewCond().And(
			builder.Eq{
				"registered": true,
			},
			builder.Expr(dateExpr),
		), "")
		if err != nil {
			return fmt.Errorf("find error: %w", err)
		}

		modelGames = make([]model.Game, 0, len(dbGames))
		for _, dbGame := range dbGames {
			modelGame := convertDBGameToModelGame(dbGame)
			modelGame.HasPassed = gameHasPassed(modelGame)

			modelGames = append(modelGames, modelGame)
		}

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("GetTodaysGames error: %w", err)
	}

	return modelGames, nil
}
