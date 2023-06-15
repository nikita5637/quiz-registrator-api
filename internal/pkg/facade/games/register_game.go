package games

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	"github.com/nikita5637/quiz-registrator-api/pkg/ics"
	"github.com/nikita5637/quiz-registrator-api/pkg/pb/registrator"
)

// RegisterGame ...
func (f *Facade) RegisterGame(ctx context.Context, gameID int32) (model.RegisterGameStatus, error) {
	registerGameStatus := model.RegisterGameStatusOK
	err := f.db.RunTX(ctx, "RegisterGame", func(ctx context.Context) error {
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

		if game.Registered {
			registerGameStatus = model.RegisterGameStatusAlreadyRegistered
			return nil
		}

		game.Registered = true
		game.Payment = int32(registrator.Payment_PAYMENT_CASH)

		err = f.gameStorage.Update(ctx, game)
		if err != nil {
			return fmt.Errorf("update game error: %w", err)
		}

		if err := f.rabbitMQProducer.Send(ctx, ics.Event{
			GameID: game.ID,
			Event:  ics.EventRegistered,
		}); err != nil {
			return fmt.Errorf("send message to rabbitMQ error: %w", err)
		}

		return nil
	})
	if err != nil {
		return model.RegisterGameStatusInvalid, fmt.Errorf("register game error: %w", err)
	}

	return registerGameStatus, nil
}
