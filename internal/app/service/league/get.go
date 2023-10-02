package league

import (
	"context"
	"errors"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/leagues"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/i18n"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	leaguepb "github.com/nikita5637/quiz-registrator-api/pkg/pb/league"
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
			st = model.GetStatus(ctx, codes.NotFound, err.Error(), "league not found", nil, leagueNotFoundLexeme)
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
