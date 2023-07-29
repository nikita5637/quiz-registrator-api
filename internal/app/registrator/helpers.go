package registrator

import (
	"context"
	"fmt"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/i18n"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	gameNotFoundLexeme = i18n.Lexeme{
		Key:      "game_not_found",
		FallBack: "Game not found",
	}
)

func getGameNotFoundStatus(ctx context.Context, err error, gameID int32) *status.Status {
	reason := fmt.Sprintf("game with id %d not found", gameID)
	return model.GetStatus(ctx, codes.NotFound, err, reason, gameNotFoundLexeme)
}
