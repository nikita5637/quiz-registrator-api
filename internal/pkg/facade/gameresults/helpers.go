package gameresults

import (
	"database/sql"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	database "github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
)

func convertDBGameResultToModelGameResult(dbGameResult database.GameResult) model.GameResult {
	return model.GameResult{
		ID:          int32(dbGameResult.ID),
		FkGameID:    int32(dbGameResult.FkGameID),
		ResultPlace: uint32(dbGameResult.Place),
		RoundPoints: model.MaybeString{
			Valid: dbGameResult.Points.Valid,
			Value: dbGameResult.Points.String,
		},
	}
}

func convertModelGameResultToDBGameResult(modelGameResult model.GameResult) database.GameResult {
	return database.GameResult{
		ID:       int(modelGameResult.ID),
		FkGameID: int(modelGameResult.FkGameID),
		Place:    uint8(modelGameResult.ResultPlace),
		Points: sql.NullString{
			Valid:  modelGameResult.RoundPoints.Valid,
			String: modelGameResult.RoundPoints.Value,
		},
	}
}
