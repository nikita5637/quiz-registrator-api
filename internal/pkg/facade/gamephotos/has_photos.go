package gamephotos

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/mono83/maybe"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/games"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/logger"
	quizlogger "github.com/nikita5637/quiz-registrator-api/internal/pkg/quiz_logger"
	usersutils "github.com/nikita5637/quiz-registrator-api/utils/users"
)

// IsGameHasPhotos ...
func (f *Facade) IsGameHasPhotos(ctx context.Context, gameID int32) (bool, error) {
	hasPhotos := false
	err := f.db.RunTX(ctx, "IsGameHasPhotos", func(ctx context.Context) error {
		if _, err := f.gameStorage.GetGameByID(ctx, int(gameID)); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return fmt.Errorf("get game by ID error: %w", games.ErrGameNotFound)
			}

			return fmt.Errorf("get game by ID error: %w", err)
		}

		urls, err := f.gamePhotoStorage.GetGamePhotosByGameID(ctx, int(gameID))
		if err != nil {
			return fmt.Errorf("get game photos by game ID error: %w", err)
		}

		logger.Debugf(ctx, "there are %d photos for a game %d", len(urls), gameID)

		hasPhotos = len(urls) > 0

		userID := maybe.Nothing[int32]()
		userFromContext := usersutils.UserFromContext(ctx)
		if userFromContext != nil {
			userID = maybe.Just(userFromContext.ID)
		}

		if err := f.quizLogger.Write(ctx, quizlogger.Params{
			UserID:     userID,
			ActionID:   quizlogger.ReadingActionID,
			MessageID:  quizlogger.GotIndicationIfGameHasPhotos,
			ObjectType: maybe.Just(quizlogger.ObjectTypeGame),
			ObjectID:   maybe.Just(gameID),
			Metadata:   nil,
		}); err != nil {
			return fmt.Errorf("write log error: %w", err)
		}

		return nil
	})
	if err != nil {
		return false, fmt.Errorf("IsGameHasPhotos error: %w", err)
	}

	return hasPhotos, nil
}
