package userroles

import (
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	database "github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
)

func convertDBUserRoleToModelUserRole(userRole database.UserRole) model.UserRole {
	return model.UserRole{
		ID:     int32(userRole.ID),
		UserID: int32(userRole.FkUserID),
		Role:   model.Role(userRole.Role),
	}
}

func convertModelUserRoleToDBUserRole(userRole model.UserRole) database.UserRole {
	return database.UserRole{
		ID:       int(userRole.ID),
		FkUserID: int(userRole.UserID),
		Role:     database.Role(userRole.Role),
	}
}
