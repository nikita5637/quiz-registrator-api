package games

import (
	"testing"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	"github.com/nikita5637/quiz-registrator-api/pkg/pb/registrator"
	"github.com/stretchr/testify/assert"
)

func TestFacade_RegisterPlayer(t *testing.T) {
	// TODO add tests after add method f.GetGameByID
	t.Run("empty fk game ID", func(t *testing.T) {
		fx := tearUp(t)

		got, err := fx.facade.RegisterPlayer(fx.ctx, 0, 1, 1, int32(registrator.Degree_DEGREE_LIKELY))
		assert.Equal(t, model.RegisterPlayerStatusInvalid, got)
		assert.Error(t, err)
	})
}
