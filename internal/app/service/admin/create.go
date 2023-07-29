package admin

import (
	"context"
	"errors"
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	userroles "github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/userroles"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/users"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	adminpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/admin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CreateUserRole ...
func (i *Implementation) CreateUserRole(ctx context.Context, req *adminpb.CreateUserRoleRequest) (*adminpb.UserRole, error) {
	createdUserRole := convertProtoUserRoleToModelUserRole(req.GetUserRole())
	if err := validateCreatedUserRole(createdUserRole); err != nil {
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
						"error": err.Error(),
					},
					errorDetails.Lexeme,
				)
			}
		}

		return nil, st.Err()
	}

	userRole, err := i.userRolesFacade.CreateUserRole(ctx, createdUserRole)
	if err != nil {
		st := status.New(codes.Internal, err.Error())
		if errors.Is(err, userroles.ErrRoleIsAlreadyAssigned) {
			st = model.GetStatus(ctx, codes.AlreadyExists, err.Error(), reasonRoleIsAlreadyAssign, nil, roleIsAlreadyAssignedToUserLexeme)
		} else if errors.Is(err, users.ErrUserNotFound) {
			st = model.GetStatus(ctx, codes.InvalidArgument, err.Error(), users.ReasonUserNotFound, nil, users.UserNotFoundLexeme)
		}

		return nil, st.Err()
	}

	return convertModelUserRoleToProtoUserRole(userRole), nil
}

func validateCreatedUserRole(userRole model.UserRole) error {
	return validation.ValidateStruct(&userRole,
		validation.Field(&userRole.UserID, validation.Required),
		validation.Field(&userRole.Role, validation.Required, validation.By(model.ValidateRole)))
}
