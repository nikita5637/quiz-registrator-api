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

const (
	reasonInvalidUserID       = "INVALID_USER_ID"
	reasonInvalidUserRole     = "INVALID_USER_ROLE"
	reasonRoleIsAlreadyAssign = "ROLE_IS_ALREADY_ASSIGN"
)

var (
	errorDetailsByField = map[string]errorDetails{
		"UserID": {
			Reason: reasonInvalidUserID,
			Lexeme: invalidUserIDLexeme,
		},
		"Role": {
			Reason: reasonInvalidUserRole,
			Lexeme: invalidUserRoleLexeme,
		},
	}

	invalidUserIDLexeme = i18n.Lexeme{
		Key:      "invalid_user_id",
		FallBack: "Invalid user ID",
	}
	invalidUserRoleLexeme = i18n.Lexeme{
		Key:      "invalid_user_role",
		FallBack: "Invalid user role",
	}

	roleIsAlreadyAssignedToUserLexeme = i18n.Lexeme{
		Key:      "role_is_already_assigned_to_user",
		FallBack: "Role is already assigned to user",
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

func getErrorDetails(keys []string) *errorDetails {
	if len(keys) == 0 {
		return nil
	}

	if v, ok := errorDetailsByField[keys[0]]; ok {
		return &v
	}

	return nil
}
