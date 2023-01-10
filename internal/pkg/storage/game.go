//go:generate mockery --case underscore --name GameStorage --with-expecter

package storage

import (
	"context"

	"github.com/go-xorm/builder"
	"github.com/nikita5637/quiz-registrator-api/internal/config"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/tx"
)

// GameStorage ...
type GameStorage interface {
	Delete(ctx context.Context, gameID int32) error
	Find(ctx context.Context, q builder.Cond, sort string) ([]model.Game, error)
	GetGameByID(ctx context.Context, id int32) (model.Game, error)
	Insert(ctx context.Context, game model.Game) (int32, error)
	Update(ctx context.Context, game model.Game) error
}

// NewGameStorage ...
func NewGameStorage(txManager *tx.Manager) GameStorage {
	switch config.GetValue("Driver").String() {
	case config.DriverMySQL:
		return mysql.NewGameStorageAdapter(txManager)
	}

	return nil
}
