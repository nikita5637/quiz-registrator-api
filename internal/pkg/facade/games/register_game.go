package games

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	"github.com/nikita5637/quiz-registrator-api/pkg/ics"
)

// RegisterGame ...
func (f *Facade) RegisterGame(ctx context.Context, id int32) (model.RegisterGameStatus, error) {
	registerGameStatus := model.RegisterGameStatusOK
	err := f.db.RunTX(ctx, "RegisterGame", func(ctx context.Context) error {
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

		if dbGame.Registered {
			registerGameStatus = model.RegisterGameStatusAlreadyRegistered
			return nil
		}

		dbGame.Registered = true
		dbGame.Payment = sql.NullInt64{
			Int64: int64(model.PaymentCash),
			Valid: true,
		}

		err = f.gameStorage.Update(ctx, *dbGame)
		if err != nil {
			return fmt.Errorf("update game error: %w", err)
		}

		if err := f.rabbitMQProducer.Send(ctx, ics.Event{
			GameID: int32(dbGame.ID),
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
