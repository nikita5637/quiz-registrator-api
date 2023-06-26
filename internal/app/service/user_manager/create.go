package usermanager

import (
	"context"
	"errors"
	"fmt"
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	users "github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/users"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/i18n"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	usermanagerpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/user_manager"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	errUserTelegramIDIsRequired = errors.New("user Telegram ID is required")
	errUserNameIsRequired       = errors.New("user name is required")

	errNameRequiredhValidateLexeme = i18n.Lexeme{
		Key:      "err_name_is_required",
		FallBack: "User name is required",
	}
	errTelegramIDRequiredValidateLexeme = i18n.Lexeme{
		Key:      "err_telegram_id_is_required",
		FallBack: "Telegram ID is required",
	}
)

// CreateUser ...
func (i *Implementation) CreateUser(ctx context.Context, req *usermanagerpb.CreateUserRequest) (*usermanagerpb.User, error) {
	if err := validateCreateUserRequest(req); err != nil {
		st := status.New(codes.InvalidArgument, err.Error())
		if errors.Is(err, errUserNameIsRequired) ||
			errors.Is(err, errUserNameLength) {
			lexeme := i18n.Lexeme{}
			if errors.Is(err, errUserNameIsRequired) {
				lexeme = errNameRequiredhValidateLexeme
			} else if errors.Is(err, errUserNameLength) {
				lexeme = errNameLengthValidateLexeme
			}

			reason := fmt.Sprintf("invalid user name")
			st = model.GetStatus(ctx, codes.InvalidArgument, err, reason, lexeme)
		} else if errors.Is(err, errUserTelegramIDIsRequired) {
			reason := fmt.Sprintf("invalid Telegram ID")
			st = model.GetStatus(ctx, codes.InvalidArgument, err, reason, errTelegramIDRequiredValidateLexeme)
		} else if errors.Is(err, errInvalidEmailFormat) {
			reason := fmt.Sprintf("invalid email")
			st = model.GetStatus(ctx, codes.InvalidArgument, err, reason, invalidEmailLexeme)
		} else if errors.Is(err, errInvalidPhoneFormat) {
			reason := fmt.Sprintf("invalid phone")
			st = model.GetStatus(ctx, codes.InvalidArgument, err, reason, invalidPhoneLexeme)
		} else if errors.Is(err, errInvalidUserState) {
			reason := fmt.Sprintf("invalid state")
			st = model.GetStatus(ctx, codes.InvalidArgument, err, reason, invalidStateLexeme)
		}

		return nil, st.Err()
	}

	user, err := i.usersFacade.CreateUser(ctx, model.User{
		Name:       req.GetUser().GetName(),
		TelegramID: req.GetUser().GetTelegramId(),
		Email:      model.NewMaybeString(req.GetUser().GetEmail()),
		Phone:      model.NewMaybeString(req.GetUser().GetPhone()),
		State:      model.UserState(req.GetUser().GetState()),
	})
	if err != nil {
		st := status.New(codes.Internal, err.Error())
		if errors.Is(err, users.ErrUserTelegramIDAlreadyExists) {
			reason := fmt.Sprintf("user with specified Telegram ID already exists")
			st = model.GetStatus(ctx, codes.AlreadyExists, err, reason, userAlreadyExistsLexeme)
		} else if errors.Is(err, users.ErrUserEmailAlreadyExists) {
			reason := fmt.Sprintf("user with specified email already exists")
			st = model.GetStatus(ctx, codes.AlreadyExists, err, reason, userAlreadyExistsLexeme)
		}

		return nil, st.Err()
	}

	return convertModelUserToProtoUser(user), nil
}

func validateCreateUserRequest(req *usermanagerpb.CreateUserRequest) error {
	if err := validation.Validate(req.GetUser().GetName(), validation.Required); err != nil {
		return errUserNameIsRequired
	}

	if err := validation.Validate(req.GetUser().GetName(), validation.Length(1, 100)); err != nil {
		return errUserNameLength
	}

	if err := validation.Validate(req.GetUser().GetTelegramId(), validation.Required); err != nil {
		return errUserTelegramIDIsRequired
	}

	if err := validation.Validate(req.GetUser().GetEmail(), validation.When(len(req.GetUser().GetEmail()) > 0, is.Email)); err != nil {
		return errInvalidEmailFormat
	}

	if err := validation.Validate(req.GetUser().GetPhone(), validation.When(len(req.GetUser().GetPhone()) > 0, validation.Match(regexp.MustCompile(`^\+79[0-9]{9}$`)))); err != nil {
		return errInvalidPhoneFormat
	}

	if err := validation.Validate(model.UserState(req.GetUser().GetState()), validation.Required, validation.By(model.ValidateUserState)); err != nil {
		return errInvalidUserState
	}

	return nil
}
