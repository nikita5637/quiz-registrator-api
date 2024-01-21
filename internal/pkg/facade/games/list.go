package games

import (
	"context"
	"fmt"

	"github.com/go-xorm/builder"
	"github.com/mono83/maybe"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	quizlogger "github.com/nikita5637/quiz-registrator-api/internal/pkg/quiz_logger"
	usersutils "github.com/nikita5637/quiz-registrator-api/utils/users"
)

// ListGames ...
func (f *Facade) ListGames(ctx context.Context) ([]model.Game, error) {
	var modelGames []model.Game
	err := f.db.RunTX(ctx, "ListGames", func(ctx context.Context) error {
		dbGames, err := f.gameStorage.Find(ctx, builder.And(
			builder.IsNull{
				"deleted_at",
			},
		), "date")
		if err != nil {
			return fmt.Errorf("find error: %w", err)
		}

		modelGames = make([]model.Game, 0, len(dbGames))
		for _, dbGame := range dbGames {
			modelGame := convertDBGameToModelGame(dbGame)
			modelGame.HasPassed = gameHasPassed(modelGame)

			if gameLink := getGameLink(modelGame); gameLink != "" {
				modelGame.GameLink = maybe.Just(gameLink)
			}

			modelGames = append(modelGames, modelGame)
		}

		userID := maybe.Nothing[int32]()
		userFromContext := usersutils.UserFromContext(ctx)
		if userFromContext != nil {
			userID = maybe.Just(userFromContext.ID)
		}

		if err := f.quizLogger.Write(ctx, quizlogger.Params{
			UserID:     userID,
			ActionID:   quizlogger.ReadingActionID,
			MessageID:  quizlogger.GotCompleteListOfGames,
			ObjectType: maybe.Nothing[string](),
			ObjectID:   maybe.Nothing[int32](),
			Metadata:   nil,
		}); err != nil {
			return fmt.Errorf("write log error: %w", err)
		}

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("ListGames error: %w", err)
	}

	return modelGames, nil
}
