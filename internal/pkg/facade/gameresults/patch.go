package gameresults

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/go-sql-driver/mysql"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
)

const (
	fieldNameGameID      = "game_id"
	fieldNameResultPlace = "result_place"
	fieldNameRoundPoints = "round_points"
)

// PatchGameResult ...
func (f *Facade) PatchGameResult(ctx context.Context, gameResult model.GameResult, paths []string) (model.GameResult, error) {
	patchedGameResult := model.GameResult{}
	err := f.db.RunTX(ctx, "PatchGameResult", func(ctx context.Context) error {
		originalDBGameResult, err := f.gameResultStorage.GetGameResultByID(ctx, int(gameResult.ID))
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return fmt.Errorf("get original game result error: %w", model.ErrGameResultNotFound)
			}

			return fmt.Errorf("get original game result error: %w", err)
		}

		patchedDBGameResult := *originalDBGameResult
		for _, path := range paths {
			switch path {
			case fieldNameGameID:
				patchedDBGameResult.FkGameID = int(gameResult.FkGameID)
			case fieldNameResultPlace:
				patchedDBGameResult.Place = uint8(gameResult.ResultPlace)
			case fieldNameRoundPoints:
				patchedDBGameResult.Points = sql.NullString{
					Valid:  gameResult.RoundPoints.Valid,
					String: gameResult.RoundPoints.Value,
				}
			}
		}

		err = f.gameResultStorage.PatchGameResult(ctx, patchedDBGameResult)
		if err != nil {
			if err, ok := err.(*mysql.MySQLError); ok {
				if err.Number == 1452 {
					return fmt.Errorf("patch game result error: %w", model.ErrGameNotFound)
				} else if err.Number == 1062 {
					return fmt.Errorf("patch game result error: %w", model.ErrGameResultAlreadyExists)
				}

				return fmt.Errorf("patch game result error: %w", err)
			}

			return fmt.Errorf("patch game result error: %w", err)
		}

		patchedGameResult = convertDBGameResultToModelGameResult(patchedDBGameResult)

		return nil
	})
	if err != nil {
		return model.GameResult{}, fmt.Errorf("patch game result error: %w", err)
	}

	return patchedGameResult, nil
}
