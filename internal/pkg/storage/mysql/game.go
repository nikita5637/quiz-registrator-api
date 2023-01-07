package mysql

import (
	"context"
	"database/sql"

	"github.com/go-xorm/builder"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
)

// GameStorageAdapter ...
type GameStorageAdapter struct {
	gameStorage *GameStorage
}

// NewGameStorageAdapter ...
func NewGameStorageAdapter(db *sql.DB) *GameStorageAdapter {
	return &GameStorageAdapter{
		gameStorage: NewGameStorage(db),
	}
}

// Delete ...
func (a *GameStorageAdapter) Delete(ctx context.Context, gameID int32) error {
	return a.gameStorage.Delete(ctx, int(gameID))
}

// Find ...
func (a *GameStorageAdapter) Find(ctx context.Context, q builder.Cond, sort string) ([]model.Game, error) {
	dbGames, err := a.gameStorage.Find(ctx, q, sort)
	if err != nil {
		return nil, err
	}

	modelGames := make([]model.Game, 0, len(dbGames))
	for _, dbGame := range dbGames {
		modelGames = append(modelGames, convertDBGameToModelGame(dbGame))
	}

	return modelGames, nil
}

// GetGameByID ...
func (a *GameStorageAdapter) GetGameByID(ctx context.Context, id int32) (model.Game, error) {
	dbGame, err := a.gameStorage.GetGameByID(ctx, int(id))
	if err != nil {
		return model.Game{}, err
	}

	return convertDBGameToModelGame(*dbGame), nil
}

// Insert ...
func (a *GameStorageAdapter) Insert(ctx context.Context, game model.Game) (int32, error) {
	id, err := a.gameStorage.Insert(ctx, convertModelGameToDBGame(game))
	if err != nil {
		return 0, err
	}

	return int32(id), nil
}

// Update ....
func (a *GameStorageAdapter) Update(ctx context.Context, game model.Game) error {
	return a.gameStorage.Update(ctx, convertModelGameToDBGame(game))
}

func convertDBGameToModelGame(game Game) model.Game {
	return model.Game{
		ID:          int32(game.ID),
		ExternalID:  int32(game.ExternalID.Int64),
		LeagueID:    int32(game.LeagueID),
		Type:        int32(game.Type),
		Number:      game.Number,
		Name:        game.Name.String,
		PlaceID:     int32(game.PlaceID),
		Date:        model.DateTime(game.Date),
		Price:       uint32(game.Price),
		PaymentType: string(game.PaymentType),
		MaxPlayers:  uint32(game.MaxPlayers),
		Payment:     int32(game.Payment.Int64),
		Registered:  game.Registered,
		CreatedAt:   model.DateTime(game.CreatedAt.Time),
		UpdatedAt:   model.DateTime(game.UpdatedAt.Time),
		DeletedAt:   model.DateTime(game.DeletedAt.Time),
	}
}

func convertModelGameToDBGame(game model.Game) Game {
	ret := Game{
		ID:          int(game.ID),
		LeagueID:    int(game.LeagueID),
		Type:        uint8(game.Type),
		Number:      game.Number,
		PlaceID:     int(game.PlaceID),
		Date:        game.Date.AsTime(),
		Price:       uint(game.Price),
		PaymentType: []byte(game.PaymentType),
		MaxPlayers:  uint8(game.MaxPlayers),
		Registered:  game.Registered,
	}

	if game.ExternalID > 0 {
		ret.ExternalID = sql.NullInt64{
			Int64: int64(game.ExternalID),
			Valid: true,
		}
	}

	if len(game.Name) > 0 {
		ret.Name = sql.NullString{
			String: game.Name,
			Valid:  true,
		}
	}

	if game.Payment > 0 {
		ret.Payment = sql.NullInt64{
			Int64: int64(game.Payment),
			Valid: true,
		}
	}

	return ret
}
