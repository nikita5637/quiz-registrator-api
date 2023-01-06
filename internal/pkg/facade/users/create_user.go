package users

import (
	"context"
	"fmt"

	"github.com/go-sql-driver/mysql"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
)

// CreateUser ...
func (f *Facade) CreateUser(ctx context.Context, user model.User) (int32, error) {
	id, err := f.userStorage.Insert(ctx, user)
	if err != nil {
		if err, ok := err.(*mysql.MySQLError); ok {
			if err.Number == 1062 {
				return 0, fmt.Errorf("create user error: %w", model.ErrUserAlreadyExists)
			}
		}

		return 0, fmt.Errorf("crete user error: %w", err)
	}

	return id, nil
}
