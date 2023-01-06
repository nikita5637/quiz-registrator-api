package gamephotos

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
)

// GetGamesWithPhotos ...
func (f *Facade) GetGamesWithPhotos(ctx context.Context, limit, offset uint32) ([]model.Game, error) {
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

	games := make([]model.Game, 0, limit)

	for i := uint32(0); i < limit; i++ {
		game, err := f.gameStorage.GetGameByID(ctx, gameIDsWithPhotos[offset+i])
		if err != nil {
			return nil, fmt.Errorf("get games with photos error: %w", err)
		}

		gameResult, err := f.gameResultStorage.GetGameResultByFkGameID(ctx, game.ID)
		if err != nil {
			return nil, fmt.Errorf("get games with photos error: %w", err)
		}

		game.ResultPlace = gameResult.Place

		games = append(games, game)
	}

	return games, nil
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
func (f *Facade) GetPhotosByGameID(ctx context.Context, gameID int32) ([]string, error) {
	_, err := f.gameStorage.GetGameByID(ctx, gameID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, model.ErrGameNotFound
		}

		return nil, fmt.Errorf("get photos by game id error: %w", err)
	}

	gamePhotos, err := f.gamePhotoStorage.GetGamePhotosByGameID(ctx, gameID)
	if err != nil {
		return nil, fmt.Errorf("get photos by game id error: %w", err)
	}

	urls := make([]string, 0, len(gamePhotos))
	for _, gamePhoto := range gamePhotos {
		urls = append(urls, gamePhoto.URL)
	}

	return urls, nil
}
