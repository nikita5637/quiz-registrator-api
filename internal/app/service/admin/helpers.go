package admin

import (
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/i18n"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	adminpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/admin"
)

type errorDetails struct {
	Reason string
	Lexeme i18n.Lexeme
}

var (
	roleIsAlreadyAssignedToUserLexeme = i18n.Lexeme{
		Key:      "role_is_already_assigned_to_user",
		FallBack: "Role is already assigned to user",
	}
)

const (
	invalidUserIDReason       = "INVALID_USER_ID"
	invalidUserRoleReason     = "INVALID_USER_ROLE"
	roleIsAlreadyAssignReason = "ROLE_IS_ALREADY_ASSIGN"
)

var (
	errorDetailsByField = map[string]errorDetails{
		"UserID": {
			Reason: invalidUserIDReason,
			Lexeme: i18n.Lexeme{
				Key:      "invalid_user_id",
				FallBack: "Invalid user ID",
			},
		},
		"Role": {
			Reason: invalidUserRoleReason,
			Lexeme: i18n.Lexeme{
				Key:      "invalid_user_role",
				FallBack: "Invalid user role",
			},
		},
	}
)

func convertModelUserRoleToProtoUserRole(userRole model.UserRole) *adminpb.UserRole {
	return &adminpb.UserRole{
		Id:     userRole.ID,
		UserId: userRole.UserID,
		Role:   adminpb.Role(userRole.Role),
	}
}

func convertProtoUserRoleToModelUserRole(userRole *adminpb.UserRole) model.UserRole {
	return model.UserRole{
		ID:     userRole.GetId(),
		UserID: userRole.GetUserId(),
		Role:   model.Role(userRole.GetRole()),
	}
}
