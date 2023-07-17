package users

import (
	"database/sql"
	"time"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	database "github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
)

func convertDBUserToModelUser(user database.User) model.User {
	ret := model.User{
		ID:         int32(user.ID),
		Name:       user.Name,
		TelegramID: user.TelegramID,
		Email: model.MaybeString{
			Valid: user.Email.Valid,
			Value: user.Email.String,
		},
		Phone: model.MaybeString{
			Valid: user.Phone.Valid,
			Value: user.Phone.String,
		},
		State: model.UserState(user.State),
		Sex: model.MaybeInt32{
			Valid: user.Sex.Valid,
			Value: int32(user.Sex.Int64),
		},
	}

	if user.Birthdate.Valid {
		ret.Birthdate = model.MaybeString{
			Valid: user.Birthdate.Valid,
			Value: user.Birthdate.Time.Format("2006-01-02"),
		}
	}

	return ret
}

func convertModelUserToDBUser(user model.User) database.User {
	ret := database.User{
		ID:         int(user.ID),
		Name:       user.Name,
		TelegramID: user.TelegramID,
		Email:      user.Email.ToSQL(),
		Phone:      user.Phone.ToSQL(),
		State:      user.State.ToSQL(),
		Sex:        user.Sex.ToSQLNullInt64(),
	}

	if user.Birthdate.Valid {
		date, err := time.Parse("2006-01-02", user.Birthdate.Value)
		valid := err == nil && !date.IsZero()
		ret.Birthdate = sql.NullTime{
			Time:  date,
			Valid: valid,
		}
	}

	return ret
}
