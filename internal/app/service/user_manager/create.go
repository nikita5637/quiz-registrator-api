package usermanager

import (
	"context"
	"errors"
	"fmt"

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

// CreateUser ...
func (i *Implementation) CreateUser(ctx context.Context, req *usermanagerpb.CreateUserRequest) (*usermanagerpb.User, error) {
	createdUser := convertProtoUserToModelUser(req.GetUser())

	logger.Debugf(ctx, "trying to create new user: %#v", createdUser)

	if err := validateCreatedUser(createdUser); err != nil {
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

	user, err := i.usersFacade.CreateUser(ctx, createdUser)
	if err != nil {
		st := status.New(codes.Internal, err.Error())
		if errors.Is(err, users.ErrUserTelegramIDAlreadyExists) {
			st = model.GetStatus(ctx, codes.AlreadyExists, err, reasonUserAlreadyExists, userAlreadyExistsLexeme)
		} else if errors.Is(err, users.ErrUserEmailAlreadyExists) {
			st = model.GetStatus(ctx, codes.AlreadyExists, err, reasonUserAlreadyExists, userAlreadyExistsLexeme)
		}

		return nil, st.Err()
	}

	logger.Debugf(ctx, "user created: %#v", user)

	return convertModelUserToProtoUser(user), nil
}

func validateCreatedUser(user model.User) error {
	return validation.ValidateStruct(&user,
		validation.Field(&user.Name, validation.Required, validation.Length(1, 100)),
		validation.Field(&user.TelegramID, validation.Required),
		validation.Field(&user.Email, validation.By(validateEmail)),
		validation.Field(&user.Phone, validation.By(validatePhone)),
		validation.Field(&user.State, validation.Required, validation.By(model.ValidateUserState)),
		validation.Field(&user.Birthdate, validation.By(validateBirthdate)),
		validation.Field(&user.Sex, validation.By(validateUserSex)),
	)
}
