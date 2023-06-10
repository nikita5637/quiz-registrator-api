package userroles

import (
	"context"
	"fmt"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
)

// ListUserRoles ...
func (f *Facade) ListUserRoles(ctx context.Context) ([]model.UserRole, error) {
	var modelUserRoles []model.UserRole
	err := f.db.RunTX(ctx, "ListUserRoles", func(ctx context.Context) error {
		dbUserRoles, err := f.userRoleStorage.GetUserRoles(ctx)
		if err != nil {
			return fmt.Errorf("list user roles error: %w", err)
		}

		modelUserRoles = make([]model.UserRole, 0, len(dbUserRoles))
		for _, dbUserRole := range dbUserRoles {
			modelUserRoles = append(modelUserRoles, convertDBUserRoleToModelUserRole(dbUserRole))
		}

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("list user roles error: %w", err)
	}

	return modelUserRoles, nil
}
