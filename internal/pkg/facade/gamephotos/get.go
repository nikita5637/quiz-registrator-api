package gamephotos

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/games"
)

// GetGameWithPhotosIDs ...
func (f *Facade) GetGameWithPhotosIDs(ctx context.Context, limit, offset uint32) ([]int32, error) {
	gameIDsWithPhotos, err := f.gamePhotoStorage.GetGameIDsWithPhotos(ctx, 0)
	if err != nil {
		return nil, fmt.Errorf("get games with photos error: %w", err)
	}

	numberOfGamesWithPhotos := uint32(len(gameIDsWithPhotos))
	if len(gameIDsWithPhotos) == 0 {
		return nil, nil
	}

	if offset >= numberOfGamesWithPhotos {
		return nil, nil
	}

	if offset+limit > numberOfGamesWithPhotos {
		limit = numberOfGamesWithPhotos - offset
	}

	ids := make([]int32, 0, limit)
	for i := uint32(0); i < limit; i++ {
		ids = append(ids, int32(gameIDsWithPhotos[offset+i]))
	}

	return ids, nil
}

// GetNumberOfGamesWithPhotos ...
func (f *Facade) GetNumberOfGamesWithPhotos(ctx context.Context) (uint32, error) {
	gameIDsWithPhotos, err := f.gamePhotoStorage.GetGameIDsWithPhotos(ctx, 0)
	if err != nil {
		return 0, fmt.Errorf("get number of games with photos error: %w", err)
	}

	return uint32(len(gameIDsWithPhotos)), nil
}

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
