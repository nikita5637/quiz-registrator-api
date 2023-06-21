//go:generate mockery --case underscore --name GamePlayerStorage --with-expecter

package storage

import (
	"context"

	"github.com/go-xorm/builder"
	"github.com/nikita5637/quiz-registrator-api/internal/config"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
	database "github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/tx"
)

// GamePlayerStorage ...
type GamePlayerStorage interface {
	Delete(ctx context.Context, id int32) error
	Find(ctx context.Context, q builder.Cond) ([]database.GamePlayer, error)
	Insert(ctx context.Context, gamePlayer model.GamePlayer) (int32, error)
}

// NewGamePlayerStorage ...
func NewGamePlayerStorage(txManager *tx.Manager) GamePlayerStorage {
	switch config.GetValue("Driver").String() {
	case config.DriverMySQL:
		return mysql.NewGamePlayerStorageAdapter(txManager)
	}

	return nil
}
