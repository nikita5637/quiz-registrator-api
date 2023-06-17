package icsfiles

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
)

// DeleteICSFile ...
func (f *Facade) DeleteICSFile(ctx context.Context, id int32) error {
	err := f.db.RunTX(ctx, "DeleteICSFile", func(ctx context.Context) error {
		_, err := f.icsFileStorage.GetICSFileByID(ctx, int(id))
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return fmt.Errorf("get ICS file by ID error: %w", model.ErrICSFileNotFound)
			}

			return fmt.Errorf("get ICS file by ID error: %w", err)
		}

		return f.icsFileStorage.DeleteICSFile(ctx, int(id))
	})
	if err != nil {
		return fmt.Errorf("delete ICS file error: %w", err)
	}

	return nil
}
