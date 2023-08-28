package games

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
)

// UpdatePayment ...
func (f *Facade) UpdatePayment(ctx context.Context, id int32, payment model.PaymentType) error {
	dbGame, err := f.gameStorage.GetGameByID(ctx, int(id))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("update game payment error: %w", ErrGameNotFound)
		}

		return fmt.Errorf("update game payment error: %w", err)
	}

	modelGame := convertDBGameToModelGame(*dbGame)

	if !modelGame.IsActive() {
		return fmt.Errorf("update game payment error: %w", ErrGameNotFound)
	}

	if payment > 0 {
		dbGame.Payment = sql.NullInt64{
			Int64: int64(payment),
			Valid: true,
		}
	}

	err = f.gameStorage.Update(ctx, *dbGame)
	if err != nil {
		return fmt.Errorf("update game payment error: %w", err)
	}

	return nil
}
