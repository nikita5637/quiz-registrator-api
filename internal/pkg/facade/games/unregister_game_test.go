package games

import (
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	"github.com/nikita5637/quiz-registrator-api/pkg/ics"
	commonpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/common"
	time_utils "github.com/nikita5637/quiz-registrator-api/utils/time"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFacade_UnregisterGame(t *testing.T) {
	timeNow := time_utils.TimeNow()

	t.Run("internal error while get game by ID", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.gameStorage.EXPECT().GetGameByID(mock.Anything, int32(1)).Return(model.Game{}, errors.New("some error"))

		got, err := fx.facade.UnregisterGame(fx.ctx, 1)
		assert.Equal(t, model.UnregisterGameStatusInvalid, got)
		assert.Error(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("error game not found while get game by ID", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.gameStorage.EXPECT().GetGameByID(mock.Anything, int32(1)).Return(model.Game{}, sql.ErrNoRows)

		got, err := fx.facade.UnregisterGame(fx.ctx, 1)
		assert.Equal(t, model.UnregisterGameStatusInvalid, got)
		assert.Error(t, err)
		assert.ErrorIs(t, err, model.ErrGameNotFound)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("found not active game", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.gameStorage.EXPECT().GetGameByID(mock.Anything, int32(1)).Return(model.Game{}, nil)

		got, err := fx.facade.UnregisterGame(fx.ctx, 1)
		assert.Equal(t, model.UnregisterGameStatusInvalid, got)
		assert.Error(t, err)
		assert.ErrorIs(t, err, model.ErrGameNotFound)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("internal error while update game", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.gameStorage.EXPECT().GetGameByID(mock.Anything, int32(1)).Return(model.Game{
			Date:       model.DateTime(timeNow.Add(1 * time.Second)),
			Registered: true,
		}, nil)

		fx.gameStorage.EXPECT().Update(mock.Anything, model.Game{
			Date:       model.DateTime(timeNow.Add(1 * time.Second)),
			Payment:    int32(commonpb.Payment_PAYMENT_INVALID),
			Registered: false,
		}).Return(errors.New("some error"))

		got, err := fx.facade.UnregisterGame(fx.ctx, 1)
		assert.Equal(t, model.UnregisterGameStatusInvalid, got)
		assert.Error(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("game not registered yet", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectCommit()

		fx.gameStorage.EXPECT().GetGameByID(mock.Anything, int32(1)).Return(model.Game{
			Date:       model.DateTime(timeNow.Add(1 * time.Second)),
			Registered: false,
		}, nil)

		got, err := fx.facade.UnregisterGame(fx.ctx, 1)
		assert.Equal(t, model.UnregisterGameStatusNotRegistered, got)
		assert.NoError(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("send message errro", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.gameStorage.EXPECT().GetGameByID(mock.Anything, int32(1)).Return(model.Game{
			ID:         1,
			Date:       model.DateTime(timeNow.Add(1 * time.Second)),
			Registered: true,
		}, nil)

		fx.gameStorage.EXPECT().Update(mock.Anything, model.Game{
			ID:         1,
			Date:       model.DateTime(timeNow.Add(1 * time.Second)),
			Payment:    int32(commonpb.Payment_PAYMENT_INVALID),
			Registered: false,
		}).Return(nil)

		fx.rabbitMQProducer.EXPECT().Send(mock.Anything, ics.Event{
			GameID: 1,
			Event:  ics.EventUnregistered,
		}).Return(errors.New("some error"))

		got, err := fx.facade.UnregisterGame(fx.ctx, 1)
		assert.Equal(t, model.UnregisterGameStatusInvalid, got)
		assert.Error(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectCommit()

		fx.gameStorage.EXPECT().GetGameByID(mock.Anything, int32(1)).Return(model.Game{
			ID:         1,
			Date:       model.DateTime(timeNow.Add(1 * time.Second)),
			Registered: true,
		}, nil)

		fx.gameStorage.EXPECT().Update(mock.Anything, model.Game{
			ID:         1,
			Date:       model.DateTime(timeNow.Add(1 * time.Second)),
			Payment:    int32(commonpb.Payment_PAYMENT_INVALID),
			Registered: false,
		}).Return(nil)

		fx.rabbitMQProducer.EXPECT().Send(mock.Anything, ics.Event{
			GameID: 1,
			Event:  ics.EventUnregistered,
		}).Return(nil)

		got, err := fx.facade.UnregisterGame(fx.ctx, 1)
		assert.Equal(t, model.UnregisterGameStatusOK, got)
		assert.NoError(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})
}
