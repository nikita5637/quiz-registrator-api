package usermanager

import (
	"errors"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/i18n"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	usermanagerpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/user_manager"
)

var (
	errInvalidEmailFormat = errors.New("invalid email format")
	errInvalidPhoneFormat = errors.New("invalid phone format")
	errInvalidUserState   = errors.New("invalid user state")
	errUserNameLength     = errors.New("name length must be between 1 and 100 characters")

	errNameLengthValidateLexeme = i18n.Lexeme{
		Key:      "err_name_length_validation",
		FallBack: "Name length must be between 1 and 100 characters",
	}
	invalidEmailLexeme = i18n.Lexeme{
		Key:      "invalid_email",
		FallBack: "Invalid email",
	}
	invalidPhoneLexeme = i18n.Lexeme{
		Key:      "invalid_phone",
		FallBack: "Invalid phone",
	}
	invalidStateLexeme = i18n.Lexeme{
		Key:      "invalid_state",
		FallBack: "Invalid state",
	}
	userAlreadyExistsLexeme = i18n.Lexeme{
		Key:      "user_already_exists",
		FallBack: "User already exists",
	}
)

func convertModelUserToProtoUser(user model.User) *usermanagerpb.User {
	return &usermanagerpb.User{
		Id:         user.ID,
		Name:       user.Name,
		TelegramId: user.TelegramID,
		Email:      user.Email.Value,
		Phone:      user.Phone.Value,
		State:      usermanagerpb.UserState(user.State),
	}
}
