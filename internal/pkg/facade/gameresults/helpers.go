package gameresults

import (
	"database/sql"

	"github.com/mono83/maybe"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	database "github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
)

func convertDBGameResultToModelGameResult(dbGameResult database.GameResult) model.GameResult {
	ret := model.GameResult{
		ID:          int32(dbGameResult.ID),
		FkGameID:    int32(dbGameResult.FkGameID),
		ResultPlace: uint32(dbGameResult.Place),
		RoundPoints: maybe.Nothing[string](),
	}

	if dbGameResult.Points.Valid {
		ret.RoundPoints = maybe.Just(dbGameResult.Points.String)
	}

	return ret
}

func convertModelGameResultToDBGameResult(modelGameResult model.GameResult) database.GameResult {
	return database.GameResult{
		ID:       int(modelGameResult.ID),
		FkGameID: int(modelGameResult.FkGameID),
		Place:    uint8(modelGameResult.ResultPlace),
		Points: sql.NullString{
			String: modelGameResult.RoundPoints.Value(),
			Valid:  modelGameResult.RoundPoints.IsPresent(),
		},
	}
}
