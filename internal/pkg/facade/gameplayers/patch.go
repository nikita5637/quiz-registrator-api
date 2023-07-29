package gameplayers

import (
	"context"
	"fmt"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/go-xorm/builder"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/games"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/users"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
)

// PatchGamePlayer ...
func (f *Facade) PatchGamePlayer(ctx context.Context, gamePlayer model.GamePlayer) (model.GamePlayer, error) {
	err := f.db.RunTX(ctx, "PatchGamePlayer", func(ctx context.Context) error {
		dbGamePlayers, err := f.gamePlayerStorage.Find(ctx, builder.NewCond().And(
			builder.Neq{
				"id": gamePlayer.ID,
			},
			builder.Eq{
				"fk_game_id": gamePlayer.GameID,
				"fk_user_id": gamePlayer.UserID.Value(),
			},
			builder.IsNull{
				"deleted_at",
			},
		))
		if err != nil {
			return fmt.Errorf("find error: %w", err)
		}

		if len(dbGamePlayers) > 0 {
			return ErrGamePlayerAlreadyRegistered
		}

		patchedDBGamePlayer := convertModelGamePlayerToDBGamePlayer(gamePlayer)
		if err := f.gamePlayerStorage.PatchGamePlayer(ctx, patchedDBGamePlayer); err != nil {
			if err, ok := err.(*mysql.MySQLError); ok {
				if err.Number == 1452 {
					if i := strings.Index(err.Message, gameIDFK1ConstraintName); i != -1 {
						return fmt.Errorf("patch game player error: %w", users.ErrUserNotFound)
					} else if i := strings.Index(err.Message, gameIDFK2ConstraintName); i != -1 {
						return fmt.Errorf("patch game player error: %w", games.ErrGameNotFound)
					}
				}
			}

			return fmt.Errorf("patch game player error: %w", err)
		}

		return nil

	})
	if err != nil {
		return model.GamePlayer{}, fmt.Errorf("PatchGamePlayer error: %w", err)
	}

	return gamePlayer, nil
}
