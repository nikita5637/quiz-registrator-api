package mysql

import (
	"context"
	"database/sql"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
)

// GamePhotoStorageAdapter ...
type GamePhotoStorageAdapter struct {
	gamePhotoStorage *GamePhotoStorage
}

// NewGamePhotoStorageAdapter ...
func NewGamePhotoStorageAdapter(db *sql.DB) *GamePhotoStorageAdapter {
	return &GamePhotoStorageAdapter{
		gamePhotoStorage: NewGamePhotoStorage(db),
	}
}

// Insert ...
func (a *GamePhotoStorageAdapter) Insert(ctx context.Context, gamePhoto model.GamePhoto) (int32, error) {
	id, err := a.gamePhotoStorage.Insert(ctx, convertModelGamePhotoToDBGamePhoto(gamePhoto))
	if err != nil {
		return 0, err
	}

	return int32(id), nil
}

// GetGamePhotosByGameID ...
func (a *GamePhotoStorageAdapter) GetGamePhotosByGameID(ctx context.Context, gameID int32) ([]model.GamePhoto, error) {
	dbGamePhotos, err := a.gamePhotoStorage.GetGamePhotoByFkGameID(ctx, int(gameID))
	if err != nil {
		return nil, err
	}

	modelGamePhotos := make([]model.GamePhoto, 0, len(dbGamePhotos))
	for _, dbGamePhoto := range dbGamePhotos {
		modelGamePhotos = append(modelGamePhotos, convertDBGamePhotoToModelGamePhoto(*dbGamePhoto))
	}

	return modelGamePhotos, nil
}

// GetGameIDsWithPhotos ...
func (a *GamePhotoStorageAdapter) GetGameIDsWithPhotos(ctx context.Context, limit uint32) ([]int32, error) {
	return a.gamePhotoStorage.GetGameIDsWithPhotos(ctx, limit)
}

// GetGameIDsWithPhotos ...
func (s *GamePhotoStorage) GetGameIDsWithPhotos(ctx context.Context, limit uint32) ([]int32, error) {
	query := `SELECT DISTINCT(g.id) FROM game_photo AS gp JOIN game AS g ON gp.fk_game_id = g.id ORDER BY date DESC`

	var args []interface{}
	if limit > 0 {
		args = append(args, limit)
		query += " LIMIT ?"
	}

	rows, err := s.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var gameIDs []int32

	for rows.Next() {
		var gameID int32
		if err := rows.Scan(
			&gameID,
		); err != nil {
			return nil, err
		}

		gameIDs = append(gameIDs, gameID)
	}

	return gameIDs, nil
}

func convertDBGamePhotoToModelGamePhoto(game GamePhoto) model.GamePhoto {
	return model.GamePhoto{
		ID:       int32(game.ID),
		FkGameID: int32(game.FkGameID),
		URL:      game.URL,
	}
}

func convertModelGamePhotoToDBGamePhoto(game model.GamePhoto) GamePhoto {
	return GamePhoto{
		ID:       int(game.ID),
		FkGameID: int(game.FkGameID),
		URL:      game.URL,
	}
}
