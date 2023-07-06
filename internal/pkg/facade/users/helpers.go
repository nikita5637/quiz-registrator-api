package users

import (
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	database "github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
)

func convertDBUserToModelUser(user database.User) model.User {
	return model.User{
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
	}
}

func convertModelUserToDBUser(user model.User) database.User {
	return database.User{
		ID:         int(user.ID),
		Name:       user.Name,
		TelegramID: user.TelegramID,
		Email:      user.Email.ToSQL(),
		Phone:      user.Phone.ToSQL(),
		State:      user.State.ToSQL(),
	}
}
