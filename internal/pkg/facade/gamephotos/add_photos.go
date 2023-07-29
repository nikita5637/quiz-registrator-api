package gamephotos

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/games"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/logger"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
)

// AddGamePhotos ...
func (f *Facade) AddGamePhotos(ctx context.Context, gameID int32, urls []string) error {
	_, err := f.gameStorage.GetGameByID(ctx, gameID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("add game photos error: %w", games.ErrGameNotFound)
		}

		return fmt.Errorf("add game photos error: %w", err)
	}

	err = f.db.RunTX(ctx, "add game photos", func(ctx context.Context) error {
		for _, url := range urls {
			gamePhoto := model.GamePhoto{
				FkGameID: gameID,
				URL:      url,
			}

			if _, err2 := f.gamePhotoStorage.Insert(ctx, gamePhoto); err2 != nil {
				return err2
			}

			logger.DebugKV(ctx, "added new game photo", "game id", gameID, "url", url)
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("add game photos error: %w", err)
	}

	return nil
}
