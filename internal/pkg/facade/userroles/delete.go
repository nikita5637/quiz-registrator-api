package userroles

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

// DeleteUserRole ...
func (f *Facade) DeleteUserRole(ctx context.Context, id int32) error {
	err := f.db.RunTX(ctx, "DeleteUserRole", func(ctx context.Context) error {
		dbCertificate, err := f.userRoleStorage.GetUserRoleByID(ctx, int(id))
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return fmt.Errorf("get user role by ID error: %w", ErrUserRoleNotFound)
			}

			return fmt.Errorf("get user role by ID error: %w", err)
		}

		if dbCertificate.DeletedAt.Valid {
			return fmt.Errorf("get user role by ID error: %w", ErrUserRoleNotFound)
		}

		return f.userRoleStorage.DeleteUserRole(ctx, int(id))
	})
	if err != nil {
		return fmt.Errorf("delete user role error: %w", err)
	}

	return nil
}
