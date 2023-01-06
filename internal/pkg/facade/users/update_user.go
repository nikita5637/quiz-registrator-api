package users

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"regexp"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

// UpdateUserEmail ...
func (f *Facade) UpdateUserEmail(ctx context.Context, userID int32, email string) error {
	user, err := f.userStorage.GetUserByID(ctx, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("update user email error: %w", model.ErrUserNotFound)
		}

		return fmt.Errorf("update user email error: %w", err)
	}

	if err = validation.Validate(email, validation.When(len(email) > 0), is.Email); err != nil {
		return fmt.Errorf("update user email error: %w", model.ErrUserEmailValidate)
	}

	user.Email = email
	user.State = model.UserStateRegistered

	err = f.userStorage.Update(ctx, user)
	if err != nil {
		return fmt.Errorf("update user email error: %w", err)
	}

	return nil
}

// UpdateUserName ...
func (f *Facade) UpdateUserName(ctx context.Context, userID int32, name string) error {
	user, err := f.userStorage.GetUserByID(ctx, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("update user name error: %w", model.ErrUserNotFound)
		}

		return fmt.Errorf("update user name error: %w", err)
	}

	if err = validation.Validate(name, validation.Required); err != nil {
		return fmt.Errorf("update user name error: %w", model.ErrUserNameValidateRequired)
	}

	if err = validation.Validate(name, validation.Match(regexp.MustCompile("^[а-яА-Я ]+$"))); err != nil {
		return fmt.Errorf("update user name error: %w", model.ErrUserNameValidateAlphabet)
	}

	if err = validation.Validate(name, validation.Length(1, 100)); err != nil {
		return fmt.Errorf("update user name error: %w", model.ErrUserNameValidateLength)
	}

	user.Name = name
	user.State = model.UserStateRegistered

	err = f.userStorage.Update(ctx, user)
	if err != nil {
		return fmt.Errorf("update user name error: %w", err)
	}

	return nil
}

// UpdateUserPhone ...
func (f *Facade) UpdateUserPhone(ctx context.Context, userID int32, phone string) error {
	user, err := f.userStorage.GetUserByID(ctx, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("update user phone error: %w", model.ErrUserNotFound)
		}

		return fmt.Errorf("update user phone error: %w", err)
	}

	if err = validation.Validate(phone, validation.When(len(phone) > 0), validation.Match(regexp.MustCompile(`^\+79[0-9]{9}$`))); err != nil {
		return fmt.Errorf("update user phone error: %w", model.ErrUserPhoneValidate)
	}

	user.Phone = phone
	user.State = model.UserStateRegistered

	err = f.userStorage.Update(ctx, user)
	if err != nil {
		return fmt.Errorf("update user phone error: %w", err)
	}

	return nil
}

// UpdateUserState ...
func (f *Facade) UpdateUserState(ctx context.Context, userID int32, state model.UserState) error {
	user, err := f.userStorage.GetUserByID(ctx, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("update user state error: %w", model.ErrUserNotFound)
		}

		return fmt.Errorf("update user state error: %w", err)
	}

	if err = validation.Validate(state, validation.Required, validation.Min(model.UserStateWelcome), validation.Max(model.UserStatesNumber-1)); err != nil {
		return fmt.Errorf("update user state error: %w", model.ErrUserStateValidate)
	}

	user.State = state

	err = f.userStorage.Update(ctx, user)
	if err != nil {
		return fmt.Errorf("update user state error: %w", err)
	}

	return nil
}
