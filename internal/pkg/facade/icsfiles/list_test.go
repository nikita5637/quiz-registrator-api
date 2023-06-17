package icsfiles

import (
	"errors"
	"testing"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	database "github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFacade_ListICSFiles(t *testing.T) {
	t.Run("error while getting ICS files", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.icsFileStorage.EXPECT().GetICSFiles(mock.Anything).Return(nil, errors.New("some error"))

		got, err := fx.facade.ListICSFiles(fx.ctx)

		assert.Nil(t, got)
		assert.Error(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectCommit()

		fx.icsFileStorage.EXPECT().GetICSFiles(mock.Anything).Return([]database.IcsFile{
			{
				ID: 1,
			},
			{
				ID: 2,
			},
		}, nil)

		got, err := fx.facade.ListICSFiles(fx.ctx)

		assert.ElementsMatch(t, []model.ICSFile{
			{
				ID: 1,
			},
			{
				ID: 2,
			},
		}, got)
		assert.NoError(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})
}
