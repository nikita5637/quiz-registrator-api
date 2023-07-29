package croupier

import (
	"context"
	"fmt"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/games"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func getGameNotFoundStatus(ctx context.Context, err error, gameID int32) *status.Status {
	reason := fmt.Sprintf("game with id %d not found", gameID)
	return model.GetStatus(ctx, codes.NotFound, err, reason, games.GameNotFoundLexeme)
}
