package model

import (
	"context"
	"testing"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/i18n"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
)

func TestGetStatus(t *testing.T) {
	t.Run("panic", func(t *testing.T) {
		assert.Panics(t, func() {
			GetStatus(context.Background(),
				codes.OK,
				"some error",
				"some reason",
				nil,
				i18n.Lexeme{
					Key:      "key",
					FallBack: "fallback",
				})
		})
	})

	t.Run("ok", func(t *testing.T) {
		got := GetStatus(context.Background(),
			codes.InvalidArgument,
			"some error",
			"some reason",
			map[string]string{},
			i18n.Lexeme{
				Key:      "key",
				FallBack: "fallback",
			})
		assert.NotNil(t, got)
	})
}
