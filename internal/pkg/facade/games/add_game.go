package games

import (
	"context"
	"fmt"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
)

// AddGame ...
func (f *Facade) AddGame(ctx context.Context, game model.Game) (int32, error) {
	if err := model.ValidateGame(game); err != nil {
		return 0, fmt.Errorf("add game error: %w", err)
	}

	return f.gameStorage.Insert(ctx, game)
}
