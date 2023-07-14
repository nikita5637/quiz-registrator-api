package place

import (
	"context"
	"errors"
	"fmt"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/places"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/i18n"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	placepb "github.com/nikita5637/quiz-registrator-api/pkg/pb/place"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	placeNotFoundLexeme = i18n.Lexeme{
		Key:      "place_not_found",
		FallBack: "Place not found",
	}
)

// GetPlace ...
func (i *Implementation) GetPlace(ctx context.Context, req *placepb.GetPlaceRequest) (*placepb.Place, error) {
	place, err := i.placesFacade.GetPlaceByID(ctx, req.GetId())
	if err != nil {
		st := status.New(codes.Internal, err.Error())
		if errors.Is(err, places.ErrPlaceNotFound) {
			reason := fmt.Sprintf("place not found")
			st = model.GetStatus(ctx, codes.NotFound, err, reason, placeNotFoundLexeme)
		}

		return nil, st.Err()
	}

	return &placepb.Place{
		Id:        place.ID,
		Address:   place.Address,
		Name:      place.Name,
		ShortName: place.ShortName,
		Latitude:  place.Latitude,
		Longitude: place.Longitude,
		MenuLink:  place.MenuLink,
	}, nil
}
