package gamephotos

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/logger"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
)

// AddGamePhotos ...
func (f *Facade) AddGamePhotos(ctx context.Context, gameID int32, urls []string) error {
	_, err := f.gameStorage.GetGameByID(ctx, gameID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.ErrGameNotFound
		}

		return fmt.Errorf("add game photos error: %w", err)
	}

	for _, url := range urls {
		gamePhoto := model.GamePhoto{
			FkGameID: gameID,
			URL:      url,
		}
		_, err := f.gamePhotoStorage.Insert(ctx, gamePhoto)
		if err != nil {
			return fmt.Errorf("add game photos error: %w", err)
		}

		logger.DebugKV(ctx, "added new game photo", "game id", gameID, "url", url)
	}

	return nil
}
