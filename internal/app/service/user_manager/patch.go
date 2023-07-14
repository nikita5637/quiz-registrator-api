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
	errUserNameAlphabet = errors.New("only Russian character set are allowed")

	errNameAlphabetValidateLexeme = i18n.Lexeme{
		Key:      "err_name_alphabet_validation",
		FallBack: "Only Russian character set are allowed",
	}
)

// PatchUser ...
func (i *Implementation) PatchUser(ctx context.Context, req *usermanagerpb.PatchUserRequest) (*usermanagerpb.User, error) {
	if err := validatePatchUserRequest(req); err != nil {
		st := status.New(codes.InvalidArgument, err.Error())
		if errors.Is(err, errUserNameAlphabet) ||
			errors.Is(err, errUserNameLength) {
			lexeme := i18n.Lexeme{}
			if errors.Is(err, errUserNameAlphabet) {
				lexeme = errNameAlphabetValidateLexeme
			} else if errors.Is(err, errUserNameLength) {
				lexeme = errNameLengthValidateLexeme
			}

			reason := fmt.Sprintf("invalid user name")
			st = model.GetStatus(ctx, codes.InvalidArgument, err, reason, lexeme)
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

	user, err := i.usersFacade.PatchUser(ctx, model.User{
		ID:         req.GetUser().GetId(),
		Name:       req.GetUser().GetName(),
		TelegramID: req.GetUser().GetTelegramId(),
		Email:      model.NewMaybeString(req.GetUser().GetEmail()),
		Phone:      model.NewMaybeString(req.GetUser().GetPhone()),
		State:      model.UserState(req.GetUser().GetState()),
	}, req.GetUpdateMask().GetPaths())
	if err != nil {
		st := status.New(codes.Internal, err.Error())
		if errors.Is(err, users.ErrUserNotFound) {
			reason := fmt.Sprintf("user not found")
			st = model.GetStatus(ctx, codes.NotFound, err, reason, i18n.UserNotFoundLexeme)
		} else if errors.Is(err, users.ErrUserTelegramIDAlreadyExists) {
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

func validatePatchUserRequest(req *usermanagerpb.PatchUserRequest) error {
	if err := validation.Validate(req.GetUser().GetName(), validation.When(len(req.GetUser().GetName()) > 0, validation.Match(regexp.MustCompile("^[а-яА-Я ]+$")))); err != nil {
		return errUserNameAlphabet
	}

	if err := validation.Validate(req.GetUser().GetName(), validation.When(len(req.GetUser().GetName()) > 0, validation.Length(1, 100))); err != nil {
		return errUserNameLength
	}

	if err := validation.Validate(req.GetUser().GetEmail(), validation.When(len(req.GetUser().GetEmail()) > 0, is.Email)); err != nil {
		return errInvalidEmailFormat
	}

	if err := validation.Validate(req.GetUser().GetPhone(), validation.When(len(req.GetUser().GetPhone()) > 0, validation.Match(regexp.MustCompile(`^\+79[0-9]{9}$`)))); err != nil {
		return errInvalidPhoneFormat
	}

	if err := validation.Validate(model.UserState(req.GetUser().GetState()), validation.When(req.GetUser().GetState() != usermanagerpb.UserState_USER_STATE_INVALID, validation.By(model.ValidateUserState))); err != nil {
		return errInvalidUserState
	}

	return nil
}
