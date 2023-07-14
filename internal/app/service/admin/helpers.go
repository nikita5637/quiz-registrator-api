package admin

import (
	"context"
	"errors"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	adminpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/admin"
)

var (
	errInvalidRole = errors.New("invalid role")
)

func convertModelUserRoleToProtoUserRole(userRole model.UserRole) *adminpb.UserRole {
	return &adminpb.UserRole{
		Id:     userRole.ID,
		UserId: userRole.UserID,
		Role:   adminpb.Role(userRole.Role),
	}
}

func validateUserRole(ctx context.Context, userRole *adminpb.UserRole) error {
	err := validation.Validate(userRole.GetRole(), validation.Required, validation.Min(int32(1)))
	if err != nil {
		return errInvalidRole
	}

	return nil
}
