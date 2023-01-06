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
	leagueNotFoundLexeme = i18n.Lexeme{
		Key:      "league_not_found",
		FallBack: "League not found",
	}
)

// GetLeagueByID ...
func (r *Registrator) GetLeagueByID(ctx context.Context, req *registrator.GetLeagueByIDRequest) (*registrator.GetLeagueByIDResponse, error) {
	l, err := r.leaguesFacade.GetLeagueByID(ctx, req.GetId())
	if err != nil {
		st := status.New(codes.Internal, err.Error())

		if errors.Is(err, model.ErrLeagueNotFound) {
			st = status.New(codes.NotFound, err.Error())
			ei := &errdetails.ErrorInfo{
				Reason: fmt.Sprintf("league with id %d not found", req.GetId()),
			}
			lm := &errdetails.LocalizedMessage{
				Locale:  i18n.GetLangFromContext(ctx),
				Message: getTranslator(leagueNotFoundLexeme)(ctx),
			}
			st, err = st.WithDetails(ei, lm)
			if err != nil {
				panic(err)
			}
		}

		return nil, st.Err()
	}

	return &registrator.GetLeagueByIDResponse{
		League: &registrator.League{
			Id:        l.ID,
			Name:      l.Name,
			ShortName: l.ShortName,
			LogoLink:  l.LogoLink,
			WebSite:   l.WebSite,
		},
	}, err
}
