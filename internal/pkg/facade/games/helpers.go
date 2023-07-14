package games

import (
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	database "github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
)

func convertDBGamePlayerToModelGamePlayer(dbGamePlayer database.GamePlayer) model.GamePlayer {
	return model.GamePlayer{
		ID:       int32(dbGamePlayer.ID),
		FkGameID: int32(dbGamePlayer.FkGameID),
		FkUserID: model.MaybeInt32{
			Valid: dbGamePlayer.FkUserID.Valid,
			Value: int32(dbGamePlayer.FkUserID.Int64),
		},
		RegisteredBy: int32(dbGamePlayer.RegisteredBy),
		Degree:       int32(dbGamePlayer.Degree),
	}
}
