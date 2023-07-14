package games

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	"github.com/nikita5637/quiz-registrator-api/pkg/ics"
	commonpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/common"
)

// UnregisterGame ...
func (f *Facade) UnregisterGame(ctx context.Context, gameID int32) (model.UnregisterGameStatus, error) {
	unregisterGameStatus := model.UnregisterGameStatusOK
	err := f.db.RunTX(ctx, "UnregisterGame", func(ctx context.Context) error {
		game, err := f.gameStorage.GetGameByID(ctx, gameID)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return fmt.Errorf("get game by ID error: %w", model.ErrGameNotFound)
			}

			return fmt.Errorf("get game by ID error: %w", err)
		}

		if !game.IsActive() {
			return model.ErrGameNotFound
		}

		if !game.Registered {
			unregisterGameStatus = model.UnregisterGameStatusNotRegistered
			return nil
		}

		game.Registered = false
		game.Payment = int32(commonpb.Payment_PAYMENT_INVALID)

		if err = f.gameStorage.Update(ctx, game); err != nil {
			return fmt.Errorf("update game error: %w", err)
		}

		if err := f.rabbitMQProducer.Send(ctx, ics.Event{
			GameID: game.ID,
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
