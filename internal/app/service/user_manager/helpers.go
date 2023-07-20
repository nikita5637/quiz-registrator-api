package usermanager

import (
	"errors"
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/mono83/maybe"
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
	if v, ok := user.Email.Get(); ok {
		ret.Email = &wrapperspb.StringValue{
			Value: v,
		}
	}
	if v, ok := user.Phone.Get(); ok {
		ret.Phone = &wrapperspb.StringValue{
			Value: v,
		}
	}
	if v, ok := user.Birthdate.Get(); ok {
		ret.Birthdate = &wrapperspb.StringValue{
			Value: v,
		}
	}
	if v, ok := user.Sex.Get(); ok {
		sex := usermanagerpb.Sex(v)
		ret.Sex = &sex
	}

	return ret
}

func convertProtoUserToModelUser(user *usermanagerpb.User) model.User {
	ret := model.User{
		ID:         user.GetId(),
		Name:       user.GetName(),
		TelegramID: user.GetTelegramId(),
		Email:      maybe.Nothing[string](),
		Phone:      maybe.Nothing[string](),
		State:      model.UserState(user.GetState()),
		Birthdate:  maybe.Nothing[string](),
		Sex:        maybe.Nothing[model.Sex](),
	}

	if user.GetEmail() != nil {
		ret.Email = maybe.Just(user.GetEmail().GetValue())
	}

	if user.GetPhone() != nil {
		ret.Phone = maybe.Just(user.GetPhone().GetValue())
	}

	if user.GetBirthdate() != nil {
		ret.Birthdate = maybe.Just(user.GetBirthdate().GetValue())
	}

	if user.Sex != nil {
		ret.Sex = maybe.Just(model.Sex(user.GetSex()))
	}

	return ret
}

func validateBirthdate(value interface{}) error {
	v, ok := value.(maybe.Maybe[string])
	if !ok {
		return errors.New("must be Maybe[string]")
	}

	return validation.Validate(v.Value(), validation.When(v.IsPresent(), validation.Required, validation.Date("2006-01-02")))
}

func validateEmail(value interface{}) error {
	v, ok := value.(maybe.Maybe[string])
	if !ok {
		return errors.New("must be Maybe[string]")
	}

	return validation.Validate(v.Value(), validation.When(v.IsPresent(), validation.Required, is.Email))
}

func validatePhone(value interface{}) error {
	v, ok := value.(maybe.Maybe[string])
	if !ok {
		return errors.New("must be Maybe[string]")
	}

	return validation.Validate(v.Value(), validation.When(v.IsPresent(), validation.Required, validation.Match(regexp.MustCompile(`^\+79[0-9]{9}$`))))
}

func validateUserSex(value interface{}) error {
	v, ok := value.(maybe.Maybe[model.Sex])
	if !ok {
		return errors.New("must be Maybe[model.Sex]")
	}

	return validation.Validate(v.Value(), validation.When(v.IsPresent(), validation.By(model.ValidateSex)))
}
