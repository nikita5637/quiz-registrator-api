package userroles

import (
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	database "github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
)

func convertModelUserRoleToDBUserRole(modelUserRole model.UserRole) database.UserRole {
	return database.UserRole{
		ID:       int(modelUserRole.ID),
		FkUserID: int(modelUserRole.UserID),
		Role:     database.Role(modelUserRole.Role),
	}
}

func convertDBUserRoleToModelUserRole(dbUserRole database.UserRole) model.UserRole {
	return model.UserRole{
		ID:     int32(dbUserRole.ID),
		UserID: int32(dbUserRole.FkUserID),
		Role:   model.Role(dbUserRole.Role),
	}
}
