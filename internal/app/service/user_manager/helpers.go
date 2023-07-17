package usermanager

import (
	"errors"
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/i18n"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	usermanagerpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/user_manager"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type errorDetails struct {
	Reason string
	Lexeme i18n.Lexeme
}

const (
	reasonInvalidUserName       = "INVALID_USER_NAME"
	reasonInvalidUserTelegramID = "INVALID_USER_TELEGRAM_ID"
	reasonInvalidUserEmail      = "INVALID_USER_EMAIL"
	reasonInvalidUserPhone      = "INVALID_USER_PHONE"
	reasonInvalidUserState      = "INVALID_USER_STATE"
	reasonInvalidUserBirthdate  = "INVALID_USER_BIRTHDATE"
	reasonInvalidUserSex        = "INVALID_USER_SEX"
	reasonUserAlreadyExists     = "USER_ALREADY_EXISTS"
	reasonUserNotFound          = "USER_NOT_FOUND"
)

var (
	userAlreadyExistsLexeme = i18n.Lexeme{
		Key:      "user_already_exists",
		FallBack: "User already exists",
	}

	errorDetailsByField = map[string]errorDetails{
		"Name": {
			Reason: reasonInvalidUserName,
			Lexeme: i18n.Lexeme{
				Key:      "invalid_user_name",
				FallBack: "Invalid user name",
			},
		},
		"TelegramID": {
			Reason: reasonInvalidUserTelegramID,
			Lexeme: i18n.Lexeme{
				Key:      "invalid_telegram_id",
				FallBack: "Invalid telegram ID",
			},
		},
		"Email": {
			Reason: reasonInvalidUserEmail,
			Lexeme: i18n.Lexeme{
				Key:      "invalid_user_email",
				FallBack: "Invalid user email",
			},
		},
		"Phone": {
			Reason: reasonInvalidUserPhone,
			Lexeme: i18n.Lexeme{
				Key:      "invalid_user_phone",
				FallBack: "Invalid user phone",
			},
		},
		"State": {
			Reason: reasonInvalidUserState,
			Lexeme: i18n.Lexeme{
				Key:      "invalid_user_state",
				FallBack: "Invalid user state",
			},
		},
		"Birthdate": {
			Reason: reasonInvalidUserBirthdate,
			Lexeme: i18n.Lexeme{
				Key:      "invalid_user_birthdate",
				FallBack: "Invalid user birthdate",
			},
		},
		"Sex": {
			Reason: reasonInvalidUserSex,
			Lexeme: i18n.Lexeme{
				Key:      "invalid_user_sex",
				FallBack: "Invalid user sex",
			},
		},
	}
)

func convertModelUserToProtoUser(user model.User) *usermanagerpb.User {
	ret := &usermanagerpb.User{
		Id:         user.ID,
		Name:       user.Name,
		TelegramId: user.TelegramID,
		State:      usermanagerpb.UserState(user.State),
	}

	if user.Email.Valid {
		ret.Email = &wrapperspb.StringValue{
			Value: user.Email.Value,
		}
	}

	if user.Phone.Valid {
		ret.Phone = &wrapperspb.StringValue{
			Value: user.Phone.Value,
		}
	}

	if user.Birthdate.Valid {
		ret.Birthdate = &wrapperspb.StringValue{
			Value: user.Birthdate.Value,
		}
	}

	if user.Sex.Valid {
		sex := usermanagerpb.Sex(user.Sex.Value)
		ret.Sex = &sex
	}

	return ret
}

func validateBirthdate(value interface{}) error {
	v, ok := value.(model.MaybeString)
	if !ok {
		return errors.New("must be MaybeString")
	}

	return validation.Validate(v.Value, validation.When(v.Valid, validation.Required, validation.Date("2006-01-02")))
}

func validateEmail(value interface{}) error {
	v, ok := value.(model.MaybeString)
	if !ok {
		return errors.New("must be MaybeString")
	}

	return validation.Validate(v.Value, validation.When(v.Valid, validation.Required, is.Email))
}

func validatePhone(value interface{}) error {
	v, ok := value.(model.MaybeString)
	if !ok {
		return errors.New("must be MaybeString")
	}

	return validation.Validate(v.Value, validation.When(v.Valid, validation.Required, validation.Match(regexp.MustCompile(`^\+79[0-9]{9}$`))))
}

func validateUserSex(value interface{}) error {
	v, ok := value.(model.MaybeInt32)
	if !ok {
		return errors.New("must be MaybeInt32")
	}

	return validation.Validate(model.Sex(v.Value), validation.When(v.Valid, validation.By(model.ValidateSex)))
}
