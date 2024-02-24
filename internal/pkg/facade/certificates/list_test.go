package certificates

import (
	"errors"
	"testing"

	"github.com/mono83/maybe"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	quizlogger "github.com/nikita5637/quiz-registrator-api/internal/pkg/quiz_logger"
	database "github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
	usersutils "github.com/nikita5637/quiz-registrator-api/utils/users"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFacade_ListCertificates(t *testing.T) {
	t.Run("error while getting certificates", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.certificateStorage.EXPECT().GetCertificates(mock.Anything).Return(nil, errors.New("some error"))

		got, err := fx.facade.ListCertificates(fx.ctx)

		assert.Nil(t, got)
		assert.Error(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("errors: write logs error", func(t *testing.T) {
		fx := tearUp(t)

		ctx := usersutils.NewContextWithUser(fx.ctx, &model.User{
			ID: 1,
		})

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.certificateStorage.EXPECT().GetCertificates(mock.Anything).Return([]database.Certificate{
			{
				ID: 1,
			},
			{
				ID: 2,
			},
		}, nil)

		fx.quizLogger.EXPECT().Write(mock.Anything, quizlogger.Params{
			UserID:     maybe.Just(int32(1)),
			ActionID:   quizlogger.ReadingActionID,
			MessageID:  quizlogger.GotCertificatesList,
			ObjectType: maybe.Nothing[string](),
			ObjectID:   maybe.Nothing[int32](),
			Metadata:   nil,
		}).Return(errors.New("some error"))

		got, err := fx.facade.ListCertificates(ctx)

		assert.Nil(t, got)
		assert.Error(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		ctx := usersutils.NewContextWithUser(fx.ctx, &model.User{
			ID: 1,
		})

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectCommit()

		fx.certificateStorage.EXPECT().GetCertificates(mock.Anything).Return([]database.Certificate{
			{
				ID: 1,
			},
			{
				ID: 2,
			},
		}, nil)

		fx.quizLogger.EXPECT().Write(mock.Anything, quizlogger.Params{
			UserID:     maybe.Just(int32(1)),
			ActionID:   quizlogger.ReadingActionID,
			MessageID:  quizlogger.GotCertificatesList,
			ObjectType: maybe.Nothing[string](),
			ObjectID:   maybe.Nothing[int32](),
			Metadata:   nil,
		}).Return(nil)

		got, err := fx.facade.ListCertificates(ctx)

		assert.ElementsMatch(t, []model.Certificate{
			{
				ID:      1,
				SpentOn: maybe.Nothing[int32](),
				Info:    maybe.Nothing[string](),
			},
			{
				ID:      2,
				SpentOn: maybe.Nothing[int32](),
				Info:    maybe.Nothing[string](),
			},
		}, got)
		assert.NoError(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})
}
