package users

import (
	"context"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
)

type userKeyType struct{}

var (
	userKey = userKeyType{}
)

// NewContextWithUser ...
func NewContextWithUser(ctx context.Context, user *model.User) context.Context {
	return context.WithValue(ctx, userKey, user)
}

// UserFromContext ...
func UserFromContext(ctx context.Context) *model.User {
	val, ok := ctx.Value(userKey).(*model.User)
	if !ok {
		return nil
	}

	return val
}
