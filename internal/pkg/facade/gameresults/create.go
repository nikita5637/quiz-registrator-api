package gameresults

import (
	"context"
	"fmt"

	"github.com/go-sql-driver/mysql"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/games"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
)

// CreateGameResult ...
func (f *Facade) CreateGameResult(ctx context.Context, gameResult model.GameResult) (model.GameResult, error) {
	createdModelGameResult := model.GameResult{}
	err := f.db.RunTX(ctx, "CreateGameResult", func(ctx context.Context) error {
		newDBGameResult := convertModelGameResultToDBGameResult(gameResult)
		id, err := f.gameResultStorage.CreateGameResult(ctx, newDBGameResult)
		if err != nil {
			if err, ok := err.(*mysql.MySQLError); ok {
				if err.Number == 1062 {
					return fmt.Errorf("create game result error: %w", ErrGameResultAlreadyExists)
				} else if err.Number == 1452 {
					return fmt.Errorf("create game result error: %w", games.ErrGameNotFound)
				}
			}

			return fmt.Errorf("create game result error: %w", err)
		}

		newDBGameResult.ID = id
		createdModelGameResult = convertDBGameResultToModelGameResult(newDBGameResult)

		return nil
	})
	if err != nil {
		return model.GameResult{}, fmt.Errorf("CreateGameResult error: %w", err)
	}

	return createdModelGameResult, nil
}
