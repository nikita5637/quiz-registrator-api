package gamephotos

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/games"
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

	modelGames := make([]model.Game, 0, limit)
	for i := uint32(0); i < limit; i++ {
		dbGame, err := f.gameStorage.GetGameByID(ctx, int(gameIDsWithPhotos[offset+i]))
		if err != nil {
			return nil, fmt.Errorf("get game by ID error: %w", err)
		}

		gameResult, err := f.gameResultStorage.GetGameResultByFkGameID(ctx, dbGame.ID)
		if err != nil {
			return nil, fmt.Errorf("get game result error: %w", err)
		}

		modelGames = append(modelGames, model.Game{
			ID:          int32(dbGame.ID),
			ExternalID:  int32(dbGame.ExternalID.Int64),
			LeagueID:    int32(dbGame.LeagueID),
			Type:        model.GameType(dbGame.Type),
			Number:      dbGame.Number,
			Name:        dbGame.Name.String,
			PlaceID:     int32(dbGame.PlaceID),
			Date:        model.DateTime(dbGame.Date),
			Price:       uint32(dbGame.Price),
			PaymentType: string(dbGame.PaymentType),
			MaxPlayers:  uint32(dbGame.MaxPlayers),
			Payment:     int32(dbGame.Payment.Int64),
			Registered:  dbGame.Registered,
			CreatedAt:   model.DateTime(dbGame.CreatedAt.Time),
			UpdatedAt:   model.DateTime(dbGame.UpdatedAt.Time),
			DeletedAt:   model.DateTime(dbGame.DeletedAt.Time),
			// additional
			ResultPlace: gameResult.ResultPlace,
		})
	}

	return modelGames, nil
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

	gamePhotos, err := f.gamePhotoStorage.GetGamePhotosByGameID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("get photos by game id error: %w", err)
	}

	urls := make([]string, 0, len(gamePhotos))
	for _, gamePhoto := range gamePhotos {
		urls = append(urls, gamePhoto.URL)
	}

	return urls, nil
}
