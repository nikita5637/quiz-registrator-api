package games

import (
	"context"
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-xorm/builder"
	"github.com/mono83/maybe"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	database "github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
)

// RegisterPlayer ...
func (f *Facade) RegisterPlayer(ctx context.Context, fkGameID, fkUserID, registeredBy, degree int32) (model.RegisterPlayerStatus, error) {
	if err := validation.Validate(fkGameID, validation.Required); err != nil {
		return model.RegisterPlayerStatusInvalid, fmt.Errorf("register player error: %w", model.ErrInvalidGameID)
	}

	if err := validation.Validate(degree, validation.Required); err != nil {
		return model.RegisterPlayerStatusInvalid, fmt.Errorf("register player error: %w", model.ErrInvalidPlayerDegree)
	}

	game, err := f.GetGameByID(ctx, fkGameID)
	if err != nil {
		return model.RegisterPlayerStatusInvalid, fmt.Errorf("register player error: %w", err)
	}

	if fkUserID != 0 {
		var dbGamePlayers []database.GamePlayer
		dbGamePlayers, err = f.gamePlayerStorage.Find(ctx, builder.NewCond().And(
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

		if len(dbGamePlayers) > 0 {
			return model.RegisterPlayerStatusAlreadyRegistered, nil
		}
	}

	if (game.NumberOfPlayers + game.NumberOfLegioners) == game.MaxPlayers {
		return model.RegisterPlayerStatusInvalid, fmt.Errorf("register player error: %w", model.ErrGameNoFreeSlots)
	}

	gamePlayer := model.GamePlayer{
		FkGameID:     fkGameID,
		FkUserID:     maybe.Nothing[int32](),
		RegisteredBy: registeredBy,
		Degree:       degree,
	}

	if fkUserID != 0 {
		gamePlayer.FkUserID = maybe.Just(fkUserID)
	}

	_, err = f.gamePlayerStorage.Insert(ctx, gamePlayer)
	if err != nil {
		return model.RegisterPlayerStatusInvalid, fmt.Errorf("register player error: %w", err)
	}

	return model.RegisterPlayerStatusOK, nil
}
