package quizlogger

import (
	"errors"
	"testing"

	"github.com/mono83/maybe"
	database "github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
	"github.com/stretchr/testify/assert"
)

func TestLogger_Write(t *testing.T) {
	t.Run("error", func(t *testing.T) {
		fx := tearUp(t)

		fx.logStorage.EXPECT().Create(fx.ctx, database.Log{
			ActionID: 1,
		}).Return(errors.New("some error"))

		err := fx.logger.Write(fx.ctx, Params{
			UserID:     maybe.Nothing[int32](),
			ActionID:   1,
			ObjectType: maybe.Nothing[string](),
			ObjectID:   maybe.Nothing[int32](),
		})
		assert.Error(t, err)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		fx.logStorage.EXPECT().Create(fx.ctx, database.Log{
			ActionID: 1,
		}).Return(nil)

		err := fx.logger.Write(fx.ctx, Params{
			UserID:     maybe.Nothing[int32](),
			ActionID:   1,
			ObjectType: maybe.Nothing[string](),
			ObjectID:   maybe.Nothing[int32](),
		})
		assert.NoError(t, err)
	})
}

func TestLogger_WriteBatch(t *testing.T) {
	t.Run("error", func(t *testing.T) {
		fx := tearUp(t)

		fx.logStorage.EXPECT().Create(fx.ctx, database.Log{
			ActionID: 1,
		}).Return(nil)

		fx.logStorage.EXPECT().Create(fx.ctx, database.Log{
			ActionID: 2,
		}).Return(errors.New("some error"))

		err := fx.logger.WriteBatch(fx.ctx, []Params{
			{
				UserID:     maybe.Nothing[int32](),
				ActionID:   1,
				ObjectType: maybe.Nothing[string](),
				ObjectID:   maybe.Nothing[int32](),
			},
			{

				UserID:     maybe.Nothing[int32](),
				ActionID:   2,
				ObjectType: maybe.Nothing[string](),
				ObjectID:   maybe.Nothing[int32](),
			},
		})
		assert.Error(t, err)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		fx.logStorage.EXPECT().Create(fx.ctx, database.Log{
			ActionID: 1,
		}).Return(nil)

		fx.logStorage.EXPECT().Create(fx.ctx, database.Log{
			ActionID: 2,
		}).Return(nil)

		err := fx.logger.WriteBatch(fx.ctx, []Params{
			{
				UserID:     maybe.Nothing[int32](),
				ActionID:   1,
				ObjectType: maybe.Nothing[string](),
				ObjectID:   maybe.Nothing[int32](),
			},
			{

				UserID:     maybe.Nothing[int32](),
				ActionID:   2,
				ObjectType: maybe.Nothing[string](),
				ObjectID:   maybe.Nothing[int32](),
			},
		})
		assert.NoError(t, err)
	})
}
