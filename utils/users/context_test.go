package users

import (
	"context"
	"testing"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	"github.com/stretchr/testify/assert"
)

func TestNewContextWithUser(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		ctx := NewContextWithUser(context.Background(), &model.User{
			ID:   1,
			Name: "name",
		})

		user := UserFromContext(ctx)
		assert.Equal(t, &model.User{
			ID:   1,
			Name: "name",
		}, user)
	})
}

func TestUserFromContext(t *testing.T) {
	t.Run("not ok", func(t *testing.T) {
		user := UserFromContext(context.Background())
		assert.Equal(t, model.User{}, user)
	})

	t.Run("ok", func(t *testing.T) {
		ctx := NewContextWithUser(context.Background(), &model.User{
			ID:   1,
			Name: "name",
		})

		user := UserFromContext(ctx)
		assert.Equal(t, &model.User{
			ID:   1,
			Name: "name",
		}, user)
	})
}
