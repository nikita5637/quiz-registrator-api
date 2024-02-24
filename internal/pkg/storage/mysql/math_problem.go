package mysql

import (
	"context"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/tx"
)

// MathProblemStorageAdapter ...
type MathProblemStorageAdapter struct {
	mathProblemStorage *MathProblemStorage
}

// NewMathProblemStorageAdapter ...
func NewMathProblemStorageAdapter(txManager *tx.Manager) *MathProblemStorageAdapter {
	return &MathProblemStorageAdapter{
		mathProblemStorage: NewMathProblemStorage(txManager),
	}
}

// CreateMathProblem ...
func (a *MathProblemStorageAdapter) CreateMathProblem(ctx context.Context, mathProblem MathProblem) (int, error) {
	return a.mathProblemStorage.Insert(ctx, mathProblem)
}

// GetMathProblemByGameID ...
func (a *MathProblemStorageAdapter) GetMathProblemByGameID(ctx context.Context, gameID int) ([]*MathProblem, error) {
	return a.mathProblemStorage.GetMathProblemByFkGameID(ctx, gameID)
}
