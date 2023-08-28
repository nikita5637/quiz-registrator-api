package games

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	"github.com/nikita5637/quiz-registrator-api/pkg/ics"
)

// UnregisterGame ...
func (f *Facade) UnregisterGame(ctx context.Context, id int32) (model.UnregisterGameStatus, error) {
	unregisterGameStatus := model.UnregisterGameStatusOK
	err := f.db.RunTX(ctx, "UnregisterGame", func(ctx context.Context) error {
		dbGame, err := f.gameStorage.GetGameByID(ctx, int(id))
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return fmt.Errorf("get game by ID error: %w", ErrGameNotFound)
			}

			return fmt.Errorf("get game by ID error: %w", err)
		}

		modelGame := convertDBGameToModelGame(*dbGame)

		if !modelGame.IsActive() {
			return ErrGameNotFound
		}

		if !dbGame.Registered {
			unregisterGameStatus = model.UnregisterGameStatusNotRegistered
			return nil
		}

		dbGame.Registered = false
		dbGame.Payment = sql.NullInt64{
			Int64: int64(model.PaymentInvalid),
			Valid: false,
		}

		if err = f.gameStorage.Update(ctx, *dbGame); err != nil {
			return fmt.Errorf("update game error: %w", err)
		}

		if err := f.rabbitMQProducer.Send(ctx, ics.Event{
			GameID: int32(dbGame.ID),
			Event:  ics.EventUnregistered,
		}); err != nil {
			return fmt.Errorf("send message to rabbitMQ error: %w", err)
		}

		return nil
	})
	if err != nil {
		return model.UnregisterGameStatusInvalid, fmt.Errorf("unregister game error: %w", err)
	}

	return unregisterGameStatus, nil
}
