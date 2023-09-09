package mysql

import (
	"context"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/tx"
)

// GamePhotoStorageAdapter ...
type GamePhotoStorageAdapter struct {
	gamePhotoStorage *GamePhotoStorage
}

// NewGamePhotoStorageAdapter ...
func NewGamePhotoStorageAdapter(txManager *tx.Manager) *GamePhotoStorageAdapter {
	return &GamePhotoStorageAdapter{
		gamePhotoStorage: NewGamePhotoStorage(txManager),
	}
}

// Insert ...
func (a *GamePhotoStorageAdapter) Insert(ctx context.Context, gamePhoto GamePhoto) (int, error) {
	return a.gamePhotoStorage.Insert(ctx, gamePhoto)
}

// GetGameIDsWithPhotos ...
func (a *GamePhotoStorageAdapter) GetGameIDsWithPhotos(ctx context.Context, limit uint32) ([]int, error) {
	return a.gamePhotoStorage.GetGameIDsWithPhotos(ctx, limit)
}

// GetGamePhotosByGameID ...
func (a *GamePhotoStorageAdapter) GetGamePhotosByGameID(ctx context.Context, gameID int) ([]*GamePhoto, error) {
	return a.gamePhotoStorage.GetGamePhotoByFkGameID(ctx, gameID)
}

// GetGameIDsWithPhotos ...
func (s *GamePhotoStorage) GetGameIDsWithPhotos(ctx context.Context, limit uint32) ([]int, error) {
	query := `SELECT DISTINCT(g.id) FROM game_photo AS gp JOIN game AS g ON gp.fk_game_id = g.id ORDER BY date DESC`

	var args []interface{}
	if limit > 0 {
		args = append(args, limit)
		query += " LIMIT ?"
	}

	rows, err := s.db.Sync(ctx).QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var gameIDs []int

	for rows.Next() {
		var gameID int
		if err := rows.Scan(
			&gameID,
		); err != nil {
			return nil, err
		}

		gameIDs = append(gameIDs, gameID)
	}

	return gameIDs, nil
}
