//go:generate mockery --case underscore --name GamePlayerStorage --with-expecter

package storage

import (
	"context"
	"database/sql"

	"github.com/go-xorm/builder"
	"github.com/nikita5637/quiz-registrator-api/internal/config"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
)

// GamePlayerStorage ...
type GamePlayerStorage interface {
	Delete(ctx context.Context, id int32) error
	Find(ctx context.Context, q builder.Cond) ([]model.GamePlayer, error)
	Insert(ctx context.Context, gamePlayer model.GamePlayer) (int32, error)
}

// NewGamePlayerStorage ...
func NewGamePlayerStorage(db *sql.DB) GamePlayerStorage {
	switch config.GetValue("Driver").String() {
	case config.DriverMySQL:
		return mysql.NewGamePlayerStorageAdapter(db)
	}

	return nil
}
