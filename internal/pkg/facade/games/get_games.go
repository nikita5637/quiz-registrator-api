package games

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	database "github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
	time_utils "github.com/nikita5637/quiz-registrator-api/utils/time"
	users_utils "github.com/nikita5637/quiz-registrator-api/utils/users"

	"github.com/go-xorm/builder"
)

var (
	availibilityGameTypes = []model.GameType{
		model.GameTypeClassic,
		model.GameTypeThematic,
	}
)

// GetGameByID guaranteed returns active game by game ID
func (f *Facade) GetGameByID(ctx context.Context, id int32) (model.Game, error) {
	dbGame, err := f.gameStorage.GetGameByID(ctx, int(id))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.Game{}, fmt.Errorf("get game by id error: %w", ErrGameNotFound)
		}

		return model.Game{}, fmt.Errorf("get game by id error: %w", err)
	}

	modelGame := convertDBGameToModelGame(*dbGame)

	if !modelGame.DeletedAt.AsTime().IsZero() {
		return model.Game{}, ErrGameNotFound
	}

	if !modelGame.IsActive() {
		return model.Game{}, ErrGameHasPassed
	}

	user := users_utils.UserFromContext(ctx)

	players, err := f.gamePlayerStorage.Find(ctx, builder.NewCond().And(
		builder.Eq{
			"fk_game_id": id,
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
		if player.FkUserID.Int64 == 0 {
			numberLegioners++
			if int32(player.RegisteredBy) == user.ID {
				myLegioners++
			}
		} else {
			numberPlayers++
			if int32(player.FkUserID.Int64) == user.ID {
				my = true
			}
		}
	}

	modelGame.My = my
	modelGame.NumberOfMyLegioners = myLegioners
	modelGame.NumberOfLegioners = numberLegioners
	modelGame.NumberOfPlayers = numberPlayers

	return modelGame, nil
}

// GetGames ...
func (f *Facade) GetGames(ctx context.Context) ([]model.Game, error) {
	user := users_utils.UserFromContext(ctx)

	dbGames, err := f.gameStorage.Find(ctx, builder.NewCond().And(
		builder.In(
			"type", availibilityGameTypes,
		),
	), "date")
	if err != nil {
		return nil, fmt.Errorf("get games error: %w", err)
	}

	modelGames := make([]model.Game, 0, len(dbGames))
	for _, dbGame := range dbGames {
		modelGame := convertDBGameToModelGame(dbGame)
		var players []database.GamePlayer
		players, err = f.gamePlayerStorage.Find(ctx, builder.And(
			builder.Eq{
				"fk_game_id": modelGame.ID,
			},
			builder.IsNull{
				"deleted_at",
			},
		))
		if err != nil {
			return nil, fmt.Errorf("get games error: %w", err)
		}

		for _, player := range players {
			if player.FkUserID.Int64 > 0 {
				if int32(player.FkUserID.Int64) == user.ID {
					modelGame.My = true
				}
				modelGame.NumberOfPlayers++
			} else {
				modelGame.NumberOfLegioners++
				if int32(player.RegisteredBy) == user.ID {
					modelGame.NumberOfMyLegioners++
				}
			}
		}

		modelGames = append(modelGames, modelGame)
	}

	return modelGames, nil
}

// GetGamesByUserID ...
func (f *Facade) GetGamesByUserID(ctx context.Context, id int32) ([]model.Game, error) {
	// TODO validate userID
	playerGames, err := f.gamePlayerStorage.Find(ctx, builder.NewCond().And(
		builder.Eq{
			"fk_user_id": id,
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
		playerGameIDs = append(playerGameIDs, int32(playerGame.FkGameID))
	}

	dbGames, err := f.gameStorage.Find(ctx, builder.In("id", playerGameIDs), "date")
	if err != nil {
		return nil, fmt.Errorf("get games by user error: %w", err)
	}

	modelGames := make([]model.Game, 0, len(dbGames))
	for _, dbGame := range dbGames {
		modelGames = append(modelGames, convertDBGameToModelGame(dbGame))
	}

	return modelGames, nil
}

// GetRegisteredGames ...
func (f *Facade) GetRegisteredGames(ctx context.Context) ([]model.Game, error) {
	user := users_utils.UserFromContext(ctx)
	dbGames, err := f.gameStorage.Find(ctx, builder.NewCond().And(
		builder.Eq{
			"registered": true,
		},
		builder.In(
			"type", availibilityGameTypes,
		),
	), "date")
	if err != nil {
		return nil, err
	}

	modelGames := make([]model.Game, 0, len(dbGames))
	for _, dbGame := range dbGames {
		modelGame := convertDBGameToModelGame(dbGame)
		var players []database.GamePlayer
		players, err = f.gamePlayerStorage.Find(ctx, builder.And(
			builder.Eq{
				"fk_game_id": modelGame.ID,
			},
			builder.IsNull{
				"deleted_at",
			},
		))
		if err != nil {
			return nil, err
		}

		for _, player := range players {
			if player.FkUserID.Int64 > 0 {
				if int32(player.FkUserID.Int64) == user.ID {
					modelGame.My = true
				}
				modelGame.NumberOfPlayers++
			} else {
				modelGame.NumberOfLegioners++
				if int32(player.RegisteredBy) == user.ID {
					modelGame.NumberOfMyLegioners++
				}
			}
		}

		modelGames = append(modelGames, modelGame)
	}

	return modelGames, err
}

// GetTodaysGames ...
func (f *Facade) GetTodaysGames(ctx context.Context) ([]model.Game, error) {
	timeNow := time_utils.TimeNow()

	dateExpr := fmt.Sprintf("date LIKE \"%04d-%02d-%02d%%\"", timeNow.Year(), timeNow.Month(), timeNow.Day())

	dbGames, err := f.gameStorage.Find(ctx, builder.NewCond().And(
		builder.Eq{
			"registered": true,
		},
		builder.Expr(dateExpr),
	), "")
	if err != nil {
		return nil, fmt.Errorf("get todays games error: %w", err)
	}

	modelGames := make([]model.Game, 0, len(dbGames))
	for _, dbGame := range dbGames {
		modelGames = append(modelGames, convertDBGameToModelGame(dbGame))
	}

	return modelGames, nil
}
