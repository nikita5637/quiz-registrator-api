//go:generate mockery --case underscore --name GamePhotosFacade --with-expecter

package photomanager

import (
	"context"

	photomanagerpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/photo_manager"
)

// GamePhotosFacade ...
type GamePhotosFacade interface {
	AddGamePhotos(ctx context.Context, gameID int32, urls []string) error
	IsGameHasPhotos(ctx context.Context, gameID int32) (bool, error)
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
