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
	p, err := r.placesFacade.GetPlaceByID(ctx, req.GetId())
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
		Place: &registrator.Place{
			Id:        p.ID,
			Address:   p.Address,
			Name:      p.Name,
			ShortName: p.ShortName,
			Longitude: p.Longitude,
			Latitude:  p.Latitude,
			MenuLink:  p.MenuLink,
		},
	}, err
}
