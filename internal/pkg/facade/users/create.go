package users

import (
	"context"
	"fmt"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
)

// CreateUser ...
func (f *Facade) CreateUser(ctx context.Context, user model.User) (model.User, error) {
	createdModelUser := model.User{}
	err := f.db.RunTX(ctx, "CreateUser", func(ctx context.Context) error {
		newDBUser := convertModelUserToDBUser(user)
		id, err := f.userStorage.Insert(ctx, newDBUser)
		if err != nil {
			if err, ok := err.(*mysql.MySQLError); ok {
				if err.Number == 1062 {
					if i := strings.Index(err.Message, "for key 'telegram_id'"); i != -1 {
						return fmt.Errorf("insert user error: %w", ErrUserTelegramIDAlreadyExists)
					} else if i := strings.Index(err.Message, "for key 'email'"); i != -1 {
						return fmt.Errorf("insert user error: %w", ErrUserEmailAlreadyExists)
					}
				}
			}

			return fmt.Errorf("insert user error: %w", err)
		}

		newDBUser.ID = id
		createdModelUser = convertDBUserToModelUser(newDBUser)

		return nil
	})
	if err != nil {
		return model.User{}, fmt.Errorf("CreateUser error: %w", err)
	}

	return createdModelUser, nil
}
