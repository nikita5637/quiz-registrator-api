package places

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/go-xorm/builder"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
)

// GetPlaceByID ...
func (f *Facade) GetPlaceByID(ctx context.Context, placeID int32) (model.Place, error) {
	place, err := f.placeStorage.GetPlaceByID(ctx, placeID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.Place{}, model.ErrPlaceNotFound
		}

		return model.Place{}, err
	}

	return place, nil
}

// GetPlaceByNameAndAddress ...
func (f *Facade) GetPlaceByNameAndAddress(ctx context.Context, name, address string) (model.Place, error) {
	places, err := f.placeStorage.Find(ctx, builder.Eq{
		"name":    name,
		"address": address,
	}, "")
	if err != nil {
		return model.Place{}, fmt.Errorf("get place by name and address error: %w", err)
	}

	if len(places) == 0 {
		return model.Place{}, fmt.Errorf("get place by name and address error: %w", model.ErrPlaceNotFound)
	}

	return places[0], nil
}
