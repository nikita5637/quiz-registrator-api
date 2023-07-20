package users

import (
	"database/sql"
	"time"

	"github.com/mono83/maybe"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	database "github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
)

func convertDBUserToModelUser(user database.User) model.User {
	ret := model.User{
		ID:         int32(user.ID),
		Name:       user.Name,
		TelegramID: user.TelegramID,
		Email:      maybe.Nothing[string](),
		Phone:      maybe.Nothing[string](),
		State:      model.UserState(user.State),
		Birthdate:  maybe.Nothing[string](),
		Sex:        maybe.Nothing[model.Sex](),
	}

	if user.Email.Valid {
		ret.Email = maybe.Just(user.Email.String)
	}

	if user.Phone.Valid {
		ret.Phone = maybe.Just(user.Phone.String)
	}

	if user.Birthdate.Valid {
		ret.Birthdate = maybe.Just(user.Birthdate.Time.Format("2006-01-02"))
	}

	if user.Sex.Valid {
		ret.Sex = maybe.Just(model.Sex(user.Sex.Int64))
	}

	return ret
}

func convertModelUserToDBUser(user model.User) database.User {
	ret := database.User{
		ID:         int(user.ID),
		Name:       user.Name,
		TelegramID: user.TelegramID,
		Email: sql.NullString{
			String: user.Email.Value(),
			Valid:  user.Email.IsPresent(),
		},
		Phone: sql.NullString{
			String: user.Phone.Value(),
			Valid:  user.Phone.IsPresent(),
		},
		State: user.State.ToSQL(),
		Sex: sql.NullInt64{
			Int64: int64(user.Sex.Value()),
			Valid: user.Sex.IsPresent(),
		},
	}

	if v, ok := user.Birthdate.Get(); ok {
		date, err := time.Parse("2006-01-02", v)
		valid := err == nil && !date.IsZero()
		ret.Birthdate = sql.NullTime{
			Time:  date,
			Valid: valid,
		}
	}

	return ret
}
