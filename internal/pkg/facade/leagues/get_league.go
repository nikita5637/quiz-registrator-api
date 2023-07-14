package leagues

import (
	"context"
	"database/sql"
	"errors"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
)

// GetLeagueByID ...
func (f *Facade) GetLeagueByID(ctx context.Context, leagueID int32) (model.League, error) {
	league, err := f.leagueStorage.GetLeagueByID(ctx, leagueID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.League{}, ErrLeagueNotFound
		}

		return model.League{}, err
	}

	return league, nil
}
