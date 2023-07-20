package games

import (
	"github.com/mono83/maybe"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	database "github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
)

func convertDBGamePlayerToModelGamePlayer(dbGamePlayer database.GamePlayer) model.GamePlayer {
	ret := model.GamePlayer{
		ID:           int32(dbGamePlayer.ID),
		FkGameID:     int32(dbGamePlayer.FkGameID),
		FkUserID:     maybe.Nothing[int32](),
		RegisteredBy: int32(dbGamePlayer.RegisteredBy),
		Degree:       int32(dbGamePlayer.Degree),
	}

	if dbGamePlayer.FkUserID.Valid {
		ret.FkUserID = maybe.Just(int32(dbGamePlayer.FkUserID.Int64))
	}

	return ret
}
