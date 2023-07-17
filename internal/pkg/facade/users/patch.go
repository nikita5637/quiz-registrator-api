package users

import (
	"context"
	"fmt"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
)

// PatchUser ...
func (f *Facade) PatchUser(ctx context.Context, user model.User) (model.User, error) {
	err := f.db.RunTX(ctx, "PatchUser", func(ctx context.Context) error {
		patchedDBUser := convertModelUserToDBUser(user)
		if err := f.userStorage.PatchUser(ctx, patchedDBUser); err != nil {
			if err, ok := err.(*mysql.MySQLError); ok {
				if err.Number == 1062 {
					if i := strings.Index(err.Message, "for key 'telegram_id'"); i != -1 {
						return fmt.Errorf("patch user error: %w", ErrUserTelegramIDAlreadyExists)
					} else if i := strings.Index(err.Message, "for key 'email'"); i != -1 {
						return fmt.Errorf("patch user error: %w", ErrUserEmailAlreadyExists)
					}
				}
			}

			return fmt.Errorf("patch user error: %w", err)
		}

		return nil
	})
	if err != nil {
		return model.User{}, fmt.Errorf("PatchUser error: %w", err)
	}

	return user, nil
}
