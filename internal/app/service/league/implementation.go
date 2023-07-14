//go:generate mockery --case underscore --name LeaguesFacade --with-expecter

package league

import (
	"context"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	leaguepb "github.com/nikita5637/quiz-registrator-api/pkg/pb/league"
)

// LeaguesFacade ...
type LeaguesFacade interface {
	GetLeagueByID(ctx context.Context, id int32) (model.League, error)
}

// Implementation ...
type Implementation struct {
	leaguesFacade LeaguesFacade

	leaguepb.UnimplementedServiceServer
}

// Config ...
type Config struct {
	LeaguesFacade LeaguesFacade
}

// New ...
func New(cfg Config) *Implementation {
	return &Implementation{
		leaguesFacade: cfg.LeaguesFacade,
	}
}
