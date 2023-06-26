package users

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
)

const (
	fieldNameName       = "name"
	fieldNameTelegramID = "telegram_id"
	fieldNameEmail      = "email"
	fieldNamePhone      = "phone"
	fieldNameState      = "state"
)

// PatchUser ...
func (f *Facade) PatchUser(ctx context.Context, user model.User, paths []string) (model.User, error) {
	patchedUser := model.User{}
	err := f.db.RunTX(ctx, "PatchUser", func(ctx context.Context) error {
		originalDBUser, err := f.userStorage.GetUserByID(ctx, int(user.ID))
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return fmt.Errorf("get user by ID error: %w", ErrUserNotFound)
			}

			return fmt.Errorf("get user by ID error: %w", err)
		}

		patchedDBUser := *originalDBUser
		for _, path := range paths {
			switch path {
			case fieldNameName:
				patchedDBUser.Name = user.Name
			case fieldNameTelegramID:
				patchedDBUser.TelegramID = user.TelegramID
			case fieldNameEmail:
				patchedDBUser.Email = user.Email.ToSQL()
			case fieldNamePhone:
				patchedDBUser.Phone = user.Phone.ToSQL()
			case fieldNameState:
				patchedDBUser.State = user.State.ToSQL()
			}
		}

		err = f.userStorage.Update(ctx, patchedDBUser)
		if err != nil {
			if err, ok := err.(*mysql.MySQLError); ok {
				if err.Number == 1062 {
					if i := strings.Index(err.Message, "for key 'telegram_id'"); i != -1 {
						return fmt.Errorf("update error: %w", ErrUserTelegramIDAlreadyExists)
					} else if i := strings.Index(err.Message, "for key 'email'"); i != -1 {
						return fmt.Errorf("update error: %w", ErrUserEmailAlreadyExists)
					}
				}
			}

			return fmt.Errorf("update error: %w", err)
		}

		patchedUser = convertDBUserToModelUser(patchedDBUser)

		return nil
	})
	if err != nil {
		return model.User{}, fmt.Errorf("PatchUser error: %w", err)
	}

	return patchedUser, nil
}
