package league

import (
	"context"
	"errors"
	"fmt"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/leagues"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/i18n"
	leaguepb "github.com/nikita5637/quiz-registrator-api/pkg/pb/league"
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

// GetLeague ...
func (i *Implementation) GetLeague(ctx context.Context, req *leaguepb.GetLeagueRequest) (*leaguepb.League, error) {
	league, err := i.leaguesFacade.GetLeagueByID(ctx, req.GetId())
	if err != nil {
		st := status.New(codes.Internal, err.Error())

		if errors.Is(err, leagues.ErrLeagueNotFound) {
			st = status.New(codes.NotFound, err.Error())
			ei := &errdetails.ErrorInfo{
				Reason: fmt.Sprintf("league not found"),
			}
			lm := &errdetails.LocalizedMessage{
				Locale:  i18n.GetLangFromContext(ctx),
				Message: i18n.GetTranslator(leagueNotFoundLexeme)(ctx),
			}
			st, err = st.WithDetails(ei, lm)
			if err != nil {
				panic(err)
			}
		}

		return nil, st.Err()
	}

	return &leaguepb.League{
		Id:        league.ID,
		Name:      league.Name,
		ShortName: league.ShortName,
		LogoLink:  league.LogoLink,
		WebSite:   league.WebSite,
	}, err
}
