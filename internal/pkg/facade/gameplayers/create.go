package gameplayers

import (
	"context"
	"fmt"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/games"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/users"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
)

// CreateGamePlayer ...
func (f *Facade) CreateGamePlayer(ctx context.Context, gamePlayer model.GamePlayer) (model.GamePlayer, error) {
	createdGamePlayer := model.GamePlayer{}

	err := f.db.RunTX(ctx, "CreateGamePlayer", func(ctx context.Context) error {
		if v, ok := gamePlayer.UserID.Get(); ok {
			existedGamePlayers, err := f.GetGamePlayersByGameID(ctx, gamePlayer.GameID)
			if err != nil {
				return fmt.Errorf("get gamme players by game ID error: %w", err)
			}

			for _, existedGamePlayer := range existedGamePlayers {
				if existedGamePlayer.UserID.Value() == v {
					return ErrGamePlayerAlreadyExists
				}
			}
		}

		newDBGamePlayer := convertModelGamePlayerToDBGamePlayer(gamePlayer)
		id, err := f.gamePlayerStorage.CreateGamePlayer(ctx, newDBGamePlayer)
		if err != nil {
			if err, ok := err.(*mysql.MySQLError); ok {
				if err.Number == 1452 {
					if i := strings.Index(err.Message, gamePlayerIBFK1ConstraintName); i != -1 {
						return fmt.Errorf("create game player error: %w", users.ErrUserNotFound)
					} else if i := strings.Index(err.Message, gamePlayerIBFK3ConstraintName); i != -1 {
						return fmt.Errorf("create game player error: %w", users.ErrUserNotFound)
					} else if i := strings.Index(err.Message, gamePlayerIBFK2ConstraintName); i != -1 {
						return fmt.Errorf("create game player error: %w", games.ErrGameNotFound)
					}
				}
			}

			return fmt.Errorf("create game player error: %w", err)
		}

		newDBGamePlayer.ID = id
		createdGamePlayer = convertDBGamePlayerToModelGamePlayer(newDBGamePlayer)

		return nil
	})
	if err != nil {
		return model.GamePlayer{}, fmt.Errorf("CreateGamePlayer error: %w", err)
	}

	return createdGamePlayer, nil
}
