package usermanager

import (
	"context"
	"errors"
	"fmt"
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	users "github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/users"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/i18n"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/logger"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	usermanagerpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/user_manager"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	minID = int32(1)
)

// PatchUser ...
func (i *Implementation) PatchUser(ctx context.Context, req *usermanagerpb.PatchUserRequest) (*usermanagerpb.User, error) {
	originalUser, err := i.usersFacade.GetUser(ctx, req.GetUser().GetId())
	if err != nil {
		st := status.New(codes.Internal, err.Error())
		if errors.Is(err, users.ErrUserNotFound) {
			st = status.New(codes.NotFound, err.Error())
			errorInfo := &errdetails.ErrorInfo{
				Reason: reasonUserNotFound,
			}
			localizedMessage := &errdetails.LocalizedMessage{
				Locale:  i18n.GetLangFromContext(ctx),
				Message: i18n.GetTranslator(i18n.UserNotFoundLexeme)(ctx),
			}
			st, _ = st.WithDetails(errorInfo, localizedMessage)
		}
		return nil, st.Err()
	}

	logger.Debugf(ctx, "trying to patch user: %#v", originalUser)

	onlyRussianAlphabet := false
	patchedUser := originalUser
	for _, path := range req.GetUpdateMask().GetPaths() {
		switch path {
		case "name":
			onlyRussianAlphabet = true
			patchedUser.Name = req.GetUser().GetName()
		case "telegram_id":
			patchedUser.TelegramID = req.GetUser().GetTelegramId()
		case "email":
			patchedUser.Email = model.MaybeString{
				Valid: req.GetUser().GetEmail() != nil,
				Value: req.GetUser().GetEmail().GetValue(),
			}
		case "phone":
			patchedUser.Phone = model.MaybeString{
				Valid: req.GetUser().GetPhone() != nil,
				Value: req.GetUser().GetPhone().GetValue(),
			}
		case "state":
			patchedUser.State = model.UserState(req.GetUser().GetState())
		case "birthdate":
			patchedUser.Birthdate = model.MaybeString{
				Valid: req.GetUser().GetBirthdate() != nil,
				Value: req.GetUser().GetBirthdate().GetValue(),
			}
		case "sex":
			patchedUser.Sex = model.MaybeInt32{
				Valid: req.GetUser().Sex != nil,
				Value: int32(req.GetUser().GetSex()),
			}
		}
	}

	err = validatePatchedUser(patchedUser, onlyRussianAlphabet)
	if err != nil {
		st := status.New(codes.InvalidArgument, err.Error())
		if validationErrors, ok := err.(validation.Errors); ok && len(validationErrors) > 0 {
			keys := make([]string, 0, len(validationErrors))
			for k := range validationErrors {
				keys = append(keys, k)
			}

			if ed, ok := errorDetailsByField[keys[0]]; ok {
				st = status.New(codes.InvalidArgument, fmt.Sprintf("%s %s", keys[0], validationErrors[keys[0]].Error()))
				errorInfo := &errdetails.ErrorInfo{
					Reason: ed.Reason,
					Metadata: map[string]string{
						"error":   err.Error(),
						"request": req.String(),
					},
				}
				localizedMessage := &errdetails.LocalizedMessage{
					Locale:  i18n.GetLangFromContext(ctx),
					Message: i18n.GetTranslator(ed.Lexeme)(ctx),
				}
				st, _ = st.WithDetails(errorInfo, localizedMessage)
			}
		}

		return nil, st.Err()
	}

	user, err := i.usersFacade.PatchUser(ctx, patchedUser)
	if err != nil {
		st := status.New(codes.Internal, err.Error())
		if errors.Is(err, users.ErrUserTelegramIDAlreadyExists) {
			st = model.GetStatus(ctx, codes.AlreadyExists, err, reasonUserAlreadyExists, userAlreadyExistsLexeme)
		} else if errors.Is(err, users.ErrUserEmailAlreadyExists) {
			st = model.GetStatus(ctx, codes.AlreadyExists, err, reasonUserAlreadyExists, userAlreadyExistsLexeme)
		}

		return nil, st.Err()
	}

	logger.Debugf(ctx, "user patched: %#v", user)

	return convertModelUserToProtoUser(user), nil
}

func validatePatchedUser(user model.User, onlyRussianAlphabet bool) error {
	return validation.ValidateStruct(&user,
		validation.Field(&user.ID, validation.Required, validation.Min(minID)),
		validation.Field(&user.Name, validation.Required, validation.Length(1, 100), validation.When(onlyRussianAlphabet, validation.Match(regexp.MustCompile("^[а-яА-Я ]+$")))),
		validation.Field(&user.TelegramID, validation.Required),
		validation.Field(&user.Email, validation.When(user.Email.Valid, validation.By(validateEmail))),
		validation.Field(&user.Phone, validation.When(user.Phone.Valid, validation.By(validatePhone))),
		validation.Field(&user.State, validation.Required, validation.By(model.ValidateUserState)),
		validation.Field(&user.Birthdate, validation.When(user.Birthdate.Valid, validation.By(validateBirthdate))),
		validation.Field(&user.Sex, validation.When(user.Sex.Valid, validation.By(validateUserSex))),
	)
}
