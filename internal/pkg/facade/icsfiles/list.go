package icsfiles

import (
	"context"
	"fmt"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
)

// ListICSFiles ...
func (f *Facade) ListICSFiles(ctx context.Context) ([]model.ICSFile, error) {
	var modelICSFiles []model.ICSFile
	err := f.db.RunTX(ctx, "ListICSFiles", func(ctx context.Context) error {
		dbICSFiles, err := f.icsFileStorage.GetICSFiles(ctx)
		if err != nil {
			return fmt.Errorf("get ICS files error: %w", err)
		}

		modelICSFiles = make([]model.ICSFile, 0, len(dbICSFiles))
		for _, dbICSFile := range dbICSFiles {
			modelICSFiles = append(modelICSFiles, convertDBICSFileToModelICSFile(dbICSFile))
		}

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("list ICS files error: %w", err)
	}

	return modelICSFiles, nil
}
