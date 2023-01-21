package mysql

import (
	"context"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/tx"
)

// LeagueStorageAdapter ...
type LeagueStorageAdapter struct {
	leagueStorage *LeagueStorage
}

// NewLeagueStorageAdapter ...
func NewLeagueStorageAdapter(txManager *tx.Manager) *LeagueStorageAdapter {
	return &LeagueStorageAdapter{
		leagueStorage: NewLeagueStorage(txManager),
	}
}

// GetLeagueByID ...
func (a *LeagueStorageAdapter) GetLeagueByID(ctx context.Context, id int32) (model.League, error) {
	leagueDB, err := a.leagueStorage.GetLeagueByID(ctx, int(id))
	if err != nil {
		return model.League{}, err
	}

	return convertDBLeagueToModelLeague(*leagueDB), nil
}

func convertDBLeagueToModelLeague(league League) model.League {
	return model.League{
		ID:        int32(league.ID),
		Name:      league.Name,
		ShortName: league.ShortName.String,
		LogoLink:  league.LogoLink.String,
		WebSite:   league.WebSite.String,
	}
}
