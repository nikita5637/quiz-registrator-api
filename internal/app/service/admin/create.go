package admin

import (
	"context"
	"errors"
	"fmt"

	userroles "github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/userroles"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/i18n"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	adminpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/admin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	invalidRoleLexeme = i18n.Lexeme{
		Key:      "invalid_role",
		FallBack: "Invalid role",
	}
	userRoleAlreayExistsLexeme = i18n.Lexeme{
		Key:      "user_role_already_exists",
		FallBack: "User role already exists",
	}
)

// CreateUserRole ...
func (i *Implementation) CreateUserRole(ctx context.Context, req *adminpb.CreateUserRoleRequest) (*adminpb.UserRole, error) {
	if err := validateCreateUserRoleRequest(ctx, req); err != nil {
		st := status.New(codes.InvalidArgument, err.Error())
		if errors.Is(err, errInvalidRole) {
			reason := fmt.Sprintf("invalid role: \"%s\"", req.GetUserRole().GetRole())
			st = model.GetStatus(ctx, codes.InvalidArgument, err, reason, invalidRoleLexeme)
		}

		return nil, st.Err()
	}

	userRole, err := i.userRolesFacade.CreateUserRole(ctx, model.UserRole{
		UserID: req.GetUserRole().GetUserId(),
		Role:   model.Role(req.GetUserRole().GetRole()),
	})
	if err != nil {
		st := status.New(codes.Internal, err.Error())
		if errors.Is(err, userroles.ErrUserRoleAlreadyExists) {
			reason := fmt.Sprintf("role %s already exists for user %d", req.GetUserRole().GetRole(), req.GetUserRole().GetUserId())
			st = model.GetStatus(ctx, codes.AlreadyExists, err, reason, userRoleAlreayExistsLexeme)
		} else if errors.Is(err, userroles.ErrUserNotFound) {
			reason := fmt.Sprintf("user %d not found", req.GetUserRole().GetUserId())
			st = model.GetStatus(ctx, codes.InvalidArgument, err, reason, i18n.UserNotFoundLexeme)
		}

		return nil, st.Err()
	}

	return convertModelUserRoleToProtoUserRole(userRole), nil
}

func validateCreateUserRoleRequest(ctx context.Context, req *adminpb.CreateUserRoleRequest) error {
	return validateUserRole(ctx, req.GetUserRole())
}
