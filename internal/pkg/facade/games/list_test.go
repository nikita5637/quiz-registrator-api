package games

import (
	"errors"
	"testing"
	"time"

	"github.com/go-xorm/builder"
	"github.com/mono83/maybe"
	"github.com/nikita5637/quiz-registrator-api/internal/config"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	database "github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
	timeutils "github.com/nikita5637/quiz-registrator-api/utils/time"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFacade_ListGames(t *testing.T) {
	timeutils.TimeNow = func() time.Time {
		return timeutils.ConvertTime("2006-01-02 15:04")
	}

	globalConfig := config.GlobalConfig{}
	globalConfig.ActiveGameLag = 3600
	config.UpdateGlobalConfig(globalConfig)

	t.Run("error: find error", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.gameStorage.EXPECT().Find(mock.Anything, builder.And(
			builder.IsNull{
				"deleted_at",
			},
		), "date").Return(nil, errors.New("some error"))

		got, err := fx.facade.ListGames(fx.ctx)
		assert.Nil(t, got)
		assert.Error(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("ok: empty list", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectCommit()

		fx.gameStorage.EXPECT().Find(mock.Anything, builder.And(
			builder.IsNull{
				"deleted_at",
			},
		), "date").Return([]database.Game{}, nil)

		got, err := fx.facade.ListGames(fx.ctx)
		assert.Equal(t, []model.Game{}, got)
		assert.NoError(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectCommit()

		fx.gameStorage.EXPECT().Find(mock.Anything, builder.And(
			builder.IsNull{
				"deleted_at",
			},
		), "date").Return([]database.Game{
			{
				ID:   1,
				Date: timeutils.TimeNow().Add(time.Hour),
			},
			{
				ID:   3,
				Date: timeutils.TimeNow().Add(-1 * time.Hour),
			},
		}, nil)

		got, err := fx.facade.ListGames(fx.ctx)
		assert.Equal(t, []model.Game{
			{
				ID:          1,
				ExternalID:  maybe.Nothing[int32](),
				Name:        maybe.Nothing[string](),
				Date:        model.DateTime(timeutils.TimeNow().Add(time.Hour)),
				PaymentType: maybe.Nothing[string](),
				Payment:     maybe.Nothing[model.Payment](),
				HasPassed:   false,
			},
			{
				ID:          3,
				ExternalID:  maybe.Nothing[int32](),
				Name:        maybe.Nothing[string](),
				Date:        model.DateTime(timeutils.TimeNow().Add(-1 * time.Hour)),
				PaymentType: maybe.Nothing[string](),
				Payment:     maybe.Nothing[model.Payment](),
				HasPassed:   true,
			},
		}, got)
		assert.NoError(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})
}
