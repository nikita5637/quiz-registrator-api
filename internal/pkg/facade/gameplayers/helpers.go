package gameplayers

import (
	"database/sql"

	"github.com/mono83/maybe"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	database "github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
)

const (
	// fk_user_id field
	gamePlayerIBFK1ConstraintName = "game_player_ibfk_1"
	// fk_game_id field
	gamePlayerIBFK2ConstraintName = "game_player_ibfk_2"
	// registered_by field
	gamePlayerIBFK3ConstraintName = "game_player_ibfk_3"
)

func convertDBGamePlayerToModelGamePlayer(gamePlayer database.GamePlayer) model.GamePlayer {
	ret := model.GamePlayer{
		ID:           int32(gamePlayer.ID),
		GameID:       int32(gamePlayer.FkGameID),
		UserID:       maybe.Nothing[int32](),
		RegisteredBy: int32(gamePlayer.RegisteredBy),
		Degree:       model.Degree(gamePlayer.Degree),
	}

	if gamePlayer.FkUserID.Valid {
		ret.UserID = maybe.Just(int32(gamePlayer.FkUserID.Int64))
	}

	return ret
}

func convertModelGamePlayerToDBGamePlayer(gamePlayer model.GamePlayer) database.GamePlayer {
	return database.GamePlayer{
		ID:       int(gamePlayer.ID),
		FkGameID: int(gamePlayer.GameID),
		FkUserID: sql.NullInt64{
			Int64: int64(gamePlayer.UserID.Value()),
			Valid: gamePlayer.UserID.IsPresent(),
		},
		RegisteredBy: int(gamePlayer.RegisteredBy),
		Degree:       uint8(gamePlayer.Degree),
	}
}
