package userroles

import (
	"context"
	"fmt"

	"github.com/go-sql-driver/mysql"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
)

// CreateUserRole ...
func (f *Facade) CreateUserRole(ctx context.Context, newUserRole model.UserRole) (model.UserRole, error) {
	createdUserRole := model.UserRole{}
	err := f.db.RunTX(ctx, "CreateUserRole", func(ctx context.Context) error {
		existedUserRoles, err := f.userRoleStorage.GetUserRolesByUserID(ctx, int(newUserRole.UserID))
		if err != nil {
			return fmt.Errorf("get user roles error: %w", err)
		}

		for _, existedUserRole := range existedUserRoles {
			if existedUserRole.Role.String() == newUserRole.Role.String() {
				return model.ErrUserRoleAlreadyExists
			}
		}

		newDBUserRole := convertModelUserRoleToDBUserRole(newUserRole)
		id, err := f.userRoleStorage.Insert(ctx, newDBUserRole)
		if err != nil {
			if err, ok := err.(*mysql.MySQLError); ok {
				if err.Number == 1452 {
					return fmt.Errorf("create user role error: %w", model.ErrUserNotFound)
				}
			}

			return fmt.Errorf("create user role error: %w", err)
		}

		newDBUserRole.ID = id
		createdUserRole = convertDBUserRoleToModelUserRole(newDBUserRole)

		return nil
	})
	if err != nil {
		return model.UserRole{}, fmt.Errorf("create user role error: %w", err)
	}

	return createdUserRole, nil
}
