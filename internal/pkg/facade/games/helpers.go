package games

import (
	"database/sql"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	database "github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
)

func convertDBGameToModelGame(game database.Game) model.Game {
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

func convertModelGameToDBGame(game model.Game) database.Game {
	ret := database.Game{
		ID:         int(game.ID),
		LeagueID:   int(game.LeagueID),
		Type:       uint8(game.Type),
		Number:     game.Number,
		PlaceID:    int(game.PlaceID),
		Date:       game.Date.AsTime(),
		Price:      uint(game.Price),
		MaxPlayers: uint8(game.MaxPlayers),
		Registered: game.Registered,
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

	if game.PaymentType != "" {
		ret.PaymentType = []byte(game.PaymentType)
	}

	if game.Payment > 0 {
		ret.Payment = sql.NullInt64{
			Int64: int64(game.Payment),
			Valid: true,
		}
	}

	return ret
}
