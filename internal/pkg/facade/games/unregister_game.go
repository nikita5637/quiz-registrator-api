package games

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	"github.com/nikita5637/quiz-registrator-api/pkg/pb/registrator"
)

// UnregisterGame ...
func (f *Facade) UnregisterGame(ctx context.Context, gameID int32) (model.UnregisterGameStatus, error) {
	game, err := f.gameStorage.GetGameByID(ctx, gameID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.UnregisterGameStatusInvalid, fmt.Errorf("unregister game error: %w", model.ErrGameNotFound)
		}

		return model.UnregisterGameStatusInvalid, fmt.Errorf("unregister game error: %w", err)
	}

	if !game.IsActive() {
		return model.UnregisterGameStatusInvalid, fmt.Errorf("unregister game error: %w", model.ErrGameNotFound)
	}

	if !game.Registered {
		return model.UnregisterGameStatusNotRegistered, nil
	}

	game.Registered = false
	game.Payment = int32(registrator.Payment_PAYMENT_INVALID)

	err = f.gameStorage.Update(ctx, game)
	if err != nil {
		return model.UnregisterGameStatusInvalid, fmt.Errorf("unregister game error: %w", err)
	}

	return model.UnregisterGameStatusOK, nil
}
