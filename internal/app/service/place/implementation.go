//go:generate mockery --case underscore --name PlacesFacade --with-expecter

package place

import (
	"context"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	placepb "github.com/nikita5637/quiz-registrator-api/pkg/pb/place"
)

// PlacesFacade ...
type PlacesFacade interface {
	GetPlaceByID(ctx context.Context, placeID int32) (model.Place, error)
}

// Implementation ...
type Implementation struct {
	placesFacade PlacesFacade

	placepb.UnimplementedServiceServer
}

// Config ...
type Config struct {
	PlacesFacade PlacesFacade
}

// New ...
func New(cfg Config) *Implementation {
	return &Implementation{
		placesFacade: cfg.PlacesFacade,
	}
}
