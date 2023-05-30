package registrator

import (
	"context"
	"errors"
	"fmt"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/i18n"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	"github.com/nikita5637/quiz-registrator-api/pkg/pb/registrator"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	placeNotFoundLexeme = i18n.Lexeme{
		Key:      "place_not_found",
		FallBack: "Place not found",
	}
)

// GetPlaceByID ...
func (r *Registrator) GetPlaceByID(ctx context.Context, req *registrator.GetPlaceByIDRequest) (*registrator.GetPlaceByIDResponse, error) {
	place, err := r.placesFacade.GetPlaceByID(ctx, req.GetId())
	if err != nil {
		st := status.New(codes.Internal, err.Error())

		if errors.Is(err, model.ErrPlaceNotFound) {
			st = status.New(codes.NotFound, err.Error())
			ei := &errdetails.ErrorInfo{
				Reason: fmt.Sprintf("place with id %d not found", req.GetId()),
			}
			lm := &errdetails.LocalizedMessage{
				Locale:  i18n.GetLangFromContext(ctx),
				Message: getTranslator(placeNotFoundLexeme)(ctx),
			}
			st, err = st.WithDetails(ei, lm)
			if err != nil {
				panic(err)
			}
		}

		return nil, st.Err()
	}

	return &registrator.GetPlaceByIDResponse{
		Place: convertModelPlaceToPBPlace(place),
	}, err
}

// GetPlaceByNameAndAddress ...
func (r *Registrator) GetPlaceByNameAndAddress(ctx context.Context, req *registrator.GetPlaceByNameAndAddressRequest) (*registrator.GetPlaceByNameAndAddressResponse, error) {
	place, err := r.placesFacade.GetPlaceByNameAndAddress(ctx, req.GetName(), req.GetAddress())
	if err != nil {
		st := status.New(codes.Internal, err.Error())

		if errors.Is(err, model.ErrPlaceNotFound) {
			st = status.New(codes.NotFound, err.Error())
			ei := &errdetails.ErrorInfo{
				Reason: fmt.Sprintf("place with name %s and address %s not found", req.GetName(), req.GetAddress()),
			}
			lm := &errdetails.LocalizedMessage{
				Locale:  i18n.GetLangFromContext(ctx),
				Message: getTranslator(placeNotFoundLexeme)(ctx),
			}
			st, err = st.WithDetails(ei, lm)
			if err != nil {
				panic(err)
			}
		}

		return nil, st.Err()
	}

	return &registrator.GetPlaceByNameAndAddressResponse{
		Place: convertModelPlaceToPBPlace(place),
	}, nil
}

func convertModelPlaceToPBPlace(place model.Place) *registrator.Place {
	return &registrator.Place{
		Id:        place.ID,
		Address:   place.Address,
		Name:      place.Name,
		ShortName: place.ShortName,
		Latitude:  place.Latitude,
		Longitude: place.Longitude,
		MenuLink:  place.MenuLink,
	}
}
