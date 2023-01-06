package games

import (
	"context"
	"errors"
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
)

// AddGame ...
func (f *Facade) AddGame(ctx context.Context, game model.Game) (int32, error) {
	err := validation.Validate(game.LeagueID, validation.Required, validation.Min(1), validation.Max(model.NumberOfLeagues-1))
	if err != nil {
		return 0, fmt.Errorf("add game error: %w", model.ErrInvalidLeagueID)
	}

	err = validation.Validate(game.Type, validation.Required, validation.By(validateGameType))
	if err != nil {
		return 0, fmt.Errorf("add game error: %w", model.ErrInvalidGameType)
	}

	err = validation.Validate(game.Number, validation.Required)
	if err != nil {
		return 0, fmt.Errorf("add game error: %w", model.ErrInvalidGameNumber)
	}

	err = validation.Validate(game.PlaceID, validation.Required)
	if err != nil {
		return 0, fmt.Errorf("add game error: %w", model.ErrInvalidPlaceID)
	}

	err = validation.Validate(game.Date, validation.Required, validation.By(validateGameDate))
	if err != nil {
		return 0, fmt.Errorf("add game error: %w", model.ErrInvalidDate)
	}

	err = validation.Validate(game.Price, validation.Required)
	if err != nil {
		return 0, fmt.Errorf("add game error: %w", model.ErrInvalidPrice)
	}

	err = validation.Validate(game.MaxPlayers, validation.Required)
	if err != nil {
		return 0, fmt.Errorf("add game error: %w", model.ErrInvalidMaxPlayers)
	}

	return f.gameStorage.Insert(ctx, game)
}

func validateGameType(value interface{}) error {
	gameType, ok := value.(int32)
	if !ok {
		return errors.New("game type is not int32")
	}

	if gameType == model.GameTypeClassic ||
		gameType == model.GameTypeThematic ||
		gameType == model.GameTypeMoviesAndMusic ||
		gameType == model.GameTypeClosed {
		return nil
	}

	return errors.New("invalid game type value")
}

func validateGameDate(value interface{}) error {
	gameDate, ok := value.(model.DateTime)
	if !ok {
		return errors.New("game date is not model.DateTime")
	}

	if gameDate.AsTime().IsZero() {
		return errors.New("invalid game date")
	}

	return nil
}
