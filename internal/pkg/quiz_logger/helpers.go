package quizlogger

import (
	"database/sql"
	"encoding/json"

	database "github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
)

func convertParamsToDBLog(params Params) database.Log {
	ret := database.Log{
		ActionID:  params.ActionID,
		MessageID: params.MessageID,
	}

	if userID, isPresent := params.UserID.Get(); isPresent {
		ret.UserID = sql.NullInt64{
			Int64: int64(userID),
			Valid: true,
		}
	}

	if objectType, isPresent := params.ObjectType.Get(); isPresent {
		ret.ObjectType = sql.NullString{
			String: objectType,
			Valid:  true,
		}
	}

	if objectID, isPresent := params.ObjectID.Get(); isPresent {
		ret.ObjectID = sql.NullInt64{
			Int64: int64(objectID),
			Valid: true,
		}
	}

	if params.Metadata != nil {
		b, err := json.Marshal(params.Metadata)
		if err == nil {
			ret.Metadata = sql.NullString{
				String: string(b),
				Valid:  true,
			}
		}
	}

	return ret
}
