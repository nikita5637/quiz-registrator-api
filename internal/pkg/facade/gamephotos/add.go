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
func (f *Facade) AddGamePhotos(ctx context.Context, id int32, urls []string) error {
	err := f.db.RunTX(ctx, "AddGamePhotos", func(ctx context.Context) error {
		if _, err := f.gameStorage.GetGameByID(ctx, int(id)); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return fmt.Errorf("get game by ID error: %w", games.ErrGameNotFound)
			}

			return fmt.Errorf("get game by ID error: %w", err)
		}

		for _, url := range urls {
			gamePhoto := model.GamePhoto{
				FkGameID: id,
				URL:      url,
			}

			if _, err := f.gamePhotoStorage.Insert(ctx, gamePhoto); err != nil {
				return fmt.Errorf("insert error: %w", err)
			}

			logger.DebugKV(ctx, "added new game photo", "game id", id, "url", url)
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("AddGamePhotos: %w", err)
	}

	return nil
}
