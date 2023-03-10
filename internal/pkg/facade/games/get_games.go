package games

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	users_utils "github.com/nikita5637/quiz-registrator-api/utils/users"

	"github.com/go-xorm/builder"
)

var (
	availibilityGameTypes = []int32{
		model.GameTypeClassic,
		model.GameTypeThematic,
	}
)

// GetGameByID guaranteed returns active game by game ID
func (f *Facade) GetGameByID(ctx context.Context, gameID int32) (model.Game, error) {
	game, err := f.gameStorage.GetGameByID(ctx, gameID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.Game{}, fmt.Errorf("get game by id error: %w", model.ErrGameNotFound)
		}

		return model.Game{}, fmt.Errorf("get game by id error: %w", err)
	}

	if !game.IsActive() {
		return model.Game{}, fmt.Errorf("get game by id error: %w", model.ErrGameNotFound)
	}

	user := users_utils.UserFromContext(ctx)

	players, err := f.gamePlayerStorage.Find(ctx, builder.NewCond().And(
		builder.Eq{
			"fk_game_id": gameID,
		},
		builder.IsNull{
			"deleted_at",
		},
	))
	if err != nil {
		return model.Game{}, fmt.Errorf("get game by id error: %w", err)
	}

	var numberLegioners uint32
	var numberPlayers uint32
	var my bool
	var myLegioners uint32

	for _, player := range players {
		if player.FkUserID == 0 {
			numberLegioners++
			if player.RegisteredBy == user.ID {
				myLegioners++
			}
		} else {
			numberPlayers++
			if player.FkUserID == user.ID {
				my = true
			}
		}
	}

	game.My = my
	game.NumberOfMyLegioners = myLegioners
	game.NumberOfLegioners = numberLegioners
	game.NumberOfPlayers = numberPlayers

	return game, nil
}

// GetGames ...
func (f *Facade) GetGames(ctx context.Context) ([]model.Game, error) {
	user := users_utils.UserFromContext(ctx)

	var err error
	var games []model.Game

	games, err = f.gameStorage.Find(ctx, builder.NewCond().And(
		builder.In(
			"type", availibilityGameTypes,
		),
	))
	if err != nil {
		return nil, fmt.Errorf("get games error: %w", err)
	}

	for i, game := range games {
		var players []model.GamePlayer
		players, err = f.gamePlayerStorage.Find(ctx, builder.And(
			builder.Eq{
				"fk_game_id": game.ID,
			},
			builder.IsNull{
				"deleted_at",
			},
		))
		if err != nil {
			return nil, fmt.Errorf("get games error: %w", err)
		}

		for _, player := range players {
			if player.FkUserID > 0 {
				if player.FkUserID == user.ID {
					games[i].My = true
				}
				games[i].NumberOfPlayers++
			} else {
				games[i].NumberOfLegioners++
				if player.RegisteredBy == user.ID {
					games[i].NumberOfMyLegioners++
				}
			}
		}
	}

	return games, nil
}

// GetGamesByUserID ...
func (f *Facade) GetGamesByUserID(ctx context.Context, userID int32) ([]model.Game, error) {
	// TODO validate userID
	playerGames, err := f.gamePlayerStorage.Find(ctx, builder.NewCond().And(
		builder.Eq{
			"fk_user_id": userID,
		},
		builder.IsNull{
			"deleted_at",
		},
	))
	if err != nil {
		return nil, fmt.Errorf("get games by user error: %w", err)
	}

	playerGameIDs := make([]int32, 0, len(playerGames))
	for _, playerGame := range playerGames {
		playerGameIDs = append(playerGameIDs, playerGame.FkGameID)
	}

	games, err := f.gameStorage.Find(ctx, builder.In("id", playerGameIDs))
	if err != nil {
		return nil, fmt.Errorf("get games by user error: %w", err)
	}

	return games, nil
}

// GetPlayersByGameID ...
func (f *Facade) GetPlayersByGameID(ctx context.Context, gameID int32) ([]model.GamePlayer, error) {
	game, err := f.gameStorage.GetGameByID(ctx, gameID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("get players by game by id error: %w", model.ErrGameNotFound)
		}

		return nil, fmt.Errorf("get players by game by id error: %w", err)
	}

	if !game.IsActive() {
		return nil, fmt.Errorf("get players by game by id error: %w", model.ErrGameNotFound)
	}

	players, err := f.gamePlayerStorage.Find(ctx, builder.NewCond().And(
		builder.Eq{
			"fk_game_id": gameID,
		},
		builder.IsNull{
			"deleted_at",
		},
	))
	if err != nil {
		return nil, fmt.Errorf("get players by game id error: %w", err)
	}

	return players, nil
}

// GetRegisteredGames ...
func (f *Facade) GetRegisteredGames(ctx context.Context) ([]model.Game, error) {
	var err error
	var games []model.Game

	user := users_utils.UserFromContext(ctx)
	games, err = f.gameStorage.Find(ctx, builder.NewCond().And(
		builder.Eq{
			"registered": true,
		},
		builder.In(
			"type", availibilityGameTypes,
		),
	))
	if err != nil {
		return nil, err
	}

	for i, game := range games {
		var players []model.GamePlayer
		players, err = f.gamePlayerStorage.Find(ctx, builder.And(
			builder.Eq{
				"fk_game_id": game.ID,
			},
			builder.IsNull{
				"deleted_at",
			},
		))
		if err != nil {
			return nil, err
		}

		for _, player := range players {
			if player.FkUserID > 0 {
				if player.FkUserID == user.ID {
					games[i].My = true
				}
				games[i].NumberOfPlayers++
			} else {
				games[i].NumberOfLegioners++
				if player.RegisteredBy == user.ID {
					games[i].NumberOfMyLegioners++
				}
			}
		}
	}

	return games, err
}
