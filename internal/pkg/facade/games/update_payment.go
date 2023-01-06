package games

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
)

// UpdatePayment ...
func (f *Facade) UpdatePayment(ctx context.Context, gameID int32, payment int32) error {
	game, err := f.gameStorage.GetGameByID(ctx, gameID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("update game payment error: %w", model.ErrGameNotFound)
		}

		return fmt.Errorf("update game payment error: %w", err)
	}

	if !game.IsActive() {
		return fmt.Errorf("update game payment error: %w", model.ErrGameNotFound)
	}

	game.Payment = payment

	err = f.gameStorage.Update(ctx, game)
	if err != nil {
		return fmt.Errorf("update game payment error: %w", err)
	}

	return nil
}
