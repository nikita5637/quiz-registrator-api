package mysql

import (
	"context"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/tx"
)

// PlaceStorageAdapter ...
type PlaceStorageAdapter struct {
	placeStorage *PlaceStorage
}

// NewPlaceStorageAdapter ...
func NewPlaceStorageAdapter(txManager *tx.Manager) *PlaceStorageAdapter {
	return &PlaceStorageAdapter{
		placeStorage: NewPlaceStorage(txManager),
	}
}

// GetPlaceByID ...
func (a *PlaceStorageAdapter) GetPlaceByID(ctx context.Context, id int32) (model.Place, error) {
	placeDB, err := a.placeStorage.GetPlaceByID(ctx, int(id))
	if err != nil {
		return model.Place{}, err
	}

	return convertDBPlaceToModelPlace(*placeDB), nil
}

func convertDBPlaceToModelPlace(place Place) model.Place {
	return model.Place{
		ID:        int32(place.ID),
		Name:      place.Name,
		Address:   place.Address,
		ShortName: place.ShortName.String,
		Latitude:  float32(place.Latitude.Float64),
		Longitude: float32(place.Longitude.Float64),
		MenuLink:  place.MenuLink.String,
	}
}
