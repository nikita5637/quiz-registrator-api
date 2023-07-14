package places

import (
	"context"
	"database/sql"
	"errors"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
)

// GetPlaceByID ...
func (f *Facade) GetPlaceByID(ctx context.Context, placeID int32) (model.Place, error) {
	place, err := f.placeStorage.GetPlaceByID(ctx, placeID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.Place{}, ErrPlaceNotFound
		}

		return model.Place{}, err
	}

	return place, nil
}
