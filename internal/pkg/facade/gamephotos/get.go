package gamephotos

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/mono83/maybe"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/games"
	quizlogger "github.com/nikita5637/quiz-registrator-api/internal/pkg/quiz_logger"
	usersutils "github.com/nikita5637/quiz-registrator-api/utils/users"
)

// GetPhotosByGameID ...
func (f *Facade) GetPhotosByGameID(ctx context.Context, id int32) ([]string, error) {
	_, err := f.gameStorage.GetGameByID(ctx, int(id))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, games.ErrGameNotFound
		}

		return nil, fmt.Errorf("get photos by game id error: %w", err)
	}

	gamePhotos, err := f.gamePhotoStorage.GetGamePhotosByGameID(ctx, int(id))
	if err != nil {
		return nil, fmt.Errorf("get photos by game id error: %w", err)
	}

	urls := make([]string, 0, len(gamePhotos))
	for _, gamePhoto := range gamePhotos {
		urls = append(urls, gamePhoto.URL)
	}

	userID := maybe.Nothing[int32]()
	userFromContext := usersutils.UserFromContext(ctx)
	if userFromContext != nil {
		userID = maybe.Just(userFromContext.ID)
	}

	if err := f.quizLogger.Write(ctx, quizlogger.Params{
		UserID:     userID,
		ActionID:   quizlogger.ReadingActionID,
		MessageID:  quizlogger.GotGamePhotos,
		ObjectType: maybe.Just(quizlogger.ObjectTypeGame),
		ObjectID:   maybe.Just(id),
		Metadata:   nil,
	}); err != nil {
		return nil, fmt.Errorf("write log error: %w", err)
	}

	return urls, nil
}
