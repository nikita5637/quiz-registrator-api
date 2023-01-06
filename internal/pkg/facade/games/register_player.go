package games

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-xorm/builder"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
)

// RegisterPlayer ...
func (f *Facade) RegisterPlayer(ctx context.Context, fkGameID, fkUserID, registeredBy, degree int32) (model.RegisterPlayerStatus, error) {
	if err := validation.Validate(fkGameID, validation.Required); err != nil {
		return model.RegisterPlayerStatusInvalid, fmt.Errorf("register player error: %w", model.ErrInvalidGameID)
	}

	if err := validation.Validate(degree, validation.Required); err != nil {
		return model.RegisterPlayerStatusInvalid, fmt.Errorf("register player error: %w", model.ErrInvalidPlayerDegree)
	}

	// TODO use facade method with additional game info
	game, err := f.gameStorage.GetGameByID(ctx, fkGameID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.RegisterPlayerStatusInvalid, fmt.Errorf("register player error: %w", model.ErrGameNotFound)
		}

		return model.RegisterPlayerStatusInvalid, fmt.Errorf("register player error: %w", err)
	}

	if (game.NumberOfPlayers + game.NumberOfLegioners) == game.MaxPlayers {
		return model.RegisterPlayerStatusInvalid, fmt.Errorf("register player error: %w", model.ErrGameNoFreeSlots)
	}

	if fkUserID != 0 {
		var records []model.GamePlayer
		records, err = f.gamePlayerStorage.Find(ctx, builder.NewCond().And(
			builder.Eq{
				"fk_game_id": fkGameID,
				"fk_user_id": fkUserID,
			},
			builder.IsNull{
				"deleted_at",
			},
		))
		if err != nil {
			return model.RegisterPlayerStatusInvalid, fmt.Errorf("register player error: %w", err)
		}

		if len(records) > 0 {
			return model.RegisterPlayerStatusAlreadyRegistered, nil
		}
	}

	_, err = f.gamePlayerStorage.Insert(ctx, model.GamePlayer{
		FkGameID:     fkGameID,
		FkUserID:     fkUserID,
		RegisteredBy: registeredBy,
		Degree:       degree,
	})
	if err != nil {
		return model.RegisterPlayerStatusInvalid, fmt.Errorf("register player error: %w", err)
	}

	return model.RegisterPlayerStatusOK, nil
}
