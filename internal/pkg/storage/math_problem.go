//go:generate mockery --case underscore --name MathProblemStorage --with-expecter

package storage

import (
	"context"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
	database "github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/tx"
)

// MathProblemStorage ...
type MathProblemStorage interface {
	CreateMathProblem(ctx context.Context, mathProblem database.MathProblem) (int, error)
	GetMathProblemByGameID(ctx context.Context, gameID int) ([]*database.MathProblem, error)
}

// NewMathProblemStorage ...
func NewMathProblemStorage(driver string, txManager *tx.Manager) MathProblemStorage {
	switch driver {
	case mysql.DriverName:
		return mysql.NewMathProblemStorageAdapter(txManager)
	}

	return nil
}
