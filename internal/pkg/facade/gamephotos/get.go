package gamephotos

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/games"
)

// GetPhotosByGameID ...
func (f *Facade) GetPhotosByGameID(ctx context.Context, id int32) ([]string, error) {
	_, err := f.gameStorage.GetGameByID(ctx, int(id))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, games.ErrGameNotFound
		}

		return nil, fmt.Errorf("get photos by game id error: %w", err)
	}

	gamePhotos, err := f.gamePhotoStorage.GetGamePhotosByGameID(ctx, int(id))
	if err != nil {
		return nil, fmt.Errorf("get photos by game id error: %w", err)
	}

	urls := make([]string, 0, len(gamePhotos))
	for _, gamePhoto := range gamePhotos {
		urls = append(urls, gamePhoto.URL)
	}

	return urls, nil
}
