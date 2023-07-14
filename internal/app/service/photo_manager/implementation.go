//go:generate mockery --case underscore --name GamePhotosFacade --with-expecter

package photomanager

import (
	"context"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	photomanagerpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/photo_manager"
)

// GamePhotosFacade ...
type GamePhotosFacade interface {
	AddGamePhotos(ctx context.Context, gameID int32, urls []string) error
	GetGamesWithPhotos(ctx context.Context, limit, offset uint32) ([]model.Game, error)
	GetNumberOfGamesWithPhotos(ctx context.Context) (uint32, error)
	GetPhotosByGameID(ctx context.Context, gameID int32) ([]string, error)
}

// Implementation ...
type Implementation struct {
	gamePhotosFacade GamePhotosFacade

	photomanagerpb.UnimplementedServiceServer
}

// Config ...
type Config struct {
	GamePhotosFacade GamePhotosFacade
}

// New ...
func New(cfg Config) *Implementation {
	return &Implementation{
		gamePhotosFacade: cfg.GamePhotosFacade,
	}
}
