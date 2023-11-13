package usermanager

import (
	"context"
	"errors"
	"fmt"
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/mono83/maybe"
	users "github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/users"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/logger"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	usermanagerpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/user_manager"
	"go.uber.org/zap"
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
			st = model.GetStatus(ctx, codes.NotFound, err.Error(), reasonUserNotFound, nil, users.UserNotFoundLexeme)
		}

		return nil, st.Err()
	}

	logger.DebugKV(ctx, "trying to patch user", zap.Reflect("original_user", originalUser))

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
			if req.GetUser().GetEmail() != nil {
				patchedUser.Email = maybe.Just(req.GetUser().GetEmail().GetValue())
			}
		case "phone":
			if req.GetUser().GetPhone() != nil {
				patchedUser.Phone = maybe.Just(req.GetUser().GetPhone().GetValue())
			}
		case "state":
			patchedUser.State = model.UserState(req.GetUser().GetState())
		case "birthdate":
			if req.GetUser().GetBirthdate() != nil {
				patchedUser.Birthdate = maybe.Just(req.GetUser().GetBirthdate().GetValue())
			}
		case "sex":
			if req.GetUser().Sex != nil {
				patchedUser.Sex = maybe.Just(model.Sex(req.GetUser().GetSex()))
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

			if errorDetails := getErrorDetails(keys); errorDetails != nil {
				st = model.GetStatus(ctx,
					codes.InvalidArgument,
					fmt.Sprintf("%s %s", keys[0], validationErrors[keys[0]].Error()),
					errorDetails.Reason,
					map[string]string{
						"error":   err.Error(),
						"request": req.String(),
					},
					errorDetails.Lexeme,
				)
			}
		}

		return nil, st.Err()
	}

	user, err := i.usersFacade.PatchUser(ctx, patchedUser)
	if err != nil {
		st := status.New(codes.Internal, err.Error())
		if errors.Is(err, users.ErrUserTelegramIDAlreadyExists) {
			st = model.GetStatus(ctx, codes.AlreadyExists, err.Error(), reasonUserAlreadyExists, nil, userAlreadyExistsLexeme)
		} else if errors.Is(err, users.ErrUserEmailAlreadyExists) {
			st = model.GetStatus(ctx, codes.AlreadyExists, err.Error(), reasonUserAlreadyExists, nil, userAlreadyExistsLexeme)
		}

		return nil, st.Err()
	}

	logger.DebugKV(ctx, "user has been patched", zap.Reflect("user", user))

	return convertModelUserToProtoUser(user), nil
}

func validatePatchedUser(user model.User, onlyRussianAlphabet bool) error {
	return validation.ValidateStruct(&user,
		validation.Field(&user.ID, validation.Required, validation.Min(minID)),
		validation.Field(&user.Name, validation.Required, validation.Length(1, 100), validation.When(onlyRussianAlphabet, validation.Match(regexp.MustCompile("^[а-яА-Я ]+$")))),
		validation.Field(&user.TelegramID, validation.Required),
		validation.Field(&user.Email, validation.By(validateEmail)),
		validation.Field(&user.Phone, validation.By(validatePhone)),
		validation.Field(&user.State, validation.Required, validation.By(model.ValidateUserState)),
		validation.Field(&user.Birthdate, validation.By(validateBirthdate)),
		validation.Field(&user.Sex, validation.By(validateUserSex)),
	)
}
