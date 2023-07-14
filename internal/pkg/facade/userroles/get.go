package userroles

import (
	"context"
	"fmt"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
)

// GetUserRolesByUserID ...
func (f *Facade) GetUserRolesByUserID(ctx context.Context, userID int32) ([]model.UserRole, error) {
	var modelUserRoles []model.UserRole
	err := f.db.RunTX(ctx, "GetUserRolesByUserID", func(ctx context.Context) error {
		dbUserRoles, err := f.userRoleStorage.GetUserRolesByUserID(ctx, int(userID))
		if err != nil {
			return fmt.Errorf("get user roles by user ID error: %w", err)
		}

		modelUserRoles = make([]model.UserRole, 0, len(dbUserRoles))
		for _, dbUserRole := range dbUserRoles {
			modelUserRoles = append(modelUserRoles, convertDBUserRoleToModelUserRole(dbUserRole))
		}

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("get user roles by user ID error: %w", err)
	}

	return modelUserRoles, nil
}
