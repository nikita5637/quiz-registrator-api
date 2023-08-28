package games

import (
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/nikita5637/quiz-registrator-api/internal/config"
	database "github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
	time_utils "github.com/nikita5637/quiz-registrator-api/utils/time"
	"github.com/stretchr/testify/assert"
)

func TestFacade_DeleteGame(t *testing.T) {
	t.Run("error sql.ErrNoRows while get game by ID", func(t *testing.T) {
		fx := tearUp(t)

		fx.gameStorage.EXPECT().GetGameByID(fx.ctx, 1).Return(nil, sql.ErrNoRows)

		err := fx.facade.DeleteGame(fx.ctx, 1)
		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrGameNotFound)
	})

	t.Run("internal error  while get game by ID", func(t *testing.T) {
		fx := tearUp(t)

		fx.gameStorage.EXPECT().GetGameByID(fx.ctx, 1).Return(nil, errors.New("some error"))

		err := fx.facade.DeleteGame(fx.ctx, 1)
		assert.Error(t, err)
	})

	t.Run("game is not active", func(t *testing.T) {
		fx := tearUp(t)

		fx.gameStorage.EXPECT().GetGameByID(fx.ctx, 1).Return(&database.Game{
			Date: time_utils.TimeNow().Add(-1 * time.Hour),
		}, nil)

		err := fx.facade.DeleteGame(fx.ctx, 1)
		assert.Error(t, err)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		globalConfig := config.GlobalConfig{}
		globalConfig.ActiveGameLag = 3600
		config.UpdateGlobalConfig(globalConfig)

		fx.gameStorage.EXPECT().GetGameByID(fx.ctx, 1).Return(&database.Game{
			Date: time_utils.TimeNow(),
		}, nil)
		fx.gameStorage.EXPECT().Delete(fx.ctx, 1).Return(nil)

		err := fx.facade.DeleteGame(fx.ctx, 1)
		assert.NoError(t, err)
	})
}
