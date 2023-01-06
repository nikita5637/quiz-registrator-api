package games

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	"github.com/nikita5637/quiz-registrator-api/pkg/pb/registrator"
)

// RegisterGame ...
func (f *Facade) RegisterGame(ctx context.Context, gameID int32) (model.RegisterGameStatus, error) {
	game, err := f.gameStorage.GetGameByID(ctx, gameID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.RegisterGameStatusInvalid, fmt.Errorf("register game error: %w", model.ErrGameNotFound)
		}

		return model.RegisterGameStatusInvalid, fmt.Errorf("register game error: %w", err)
	}

	if !game.IsActive() {
		return model.RegisterGameStatusInvalid, fmt.Errorf("register game error: %w", model.ErrGameNotFound)
	}

	if game.Registered {
		return model.RegisterGameStatusAlreadyRegistered, nil
	}

	game.Registered = true
	game.Payment = int32(registrator.Payment_PAYMENT_CASH)

	err = f.gameStorage.Update(ctx, game)
	if err != nil {
		return model.RegisterGameStatusInvalid, fmt.Errorf("register game error: %w", err)
	}

	return model.RegisterGameStatusOK, nil
}
