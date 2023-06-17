package icsfiles

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
)

// GetICSFile ...
func (f *Facade) GetICSFile(ctx context.Context, id int32) (model.ICSFile, error) {
	var modelICSFile model.ICSFile
	err := f.db.RunTX(ctx, "GetICSFile", func(ctx context.Context) error {
		dbICSFile, err := f.icsFileStorage.GetICSFileByID(ctx, int(id))
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return fmt.Errorf("get ICS file by ID error: %w", model.ErrICSFileNotFound)
			}

			return fmt.Errorf("get ICS file by ID error: %w", err)
		}

		modelICSFile = convertDBICSFileToModelICSFile(*dbICSFile)

		return nil
	})
	if err != nil {
		return model.ICSFile{}, fmt.Errorf("get ICS file error: %w", err)
	}

	return modelICSFile, nil
}

// GetICSFileByGameID ...
func (f *Facade) GetICSFileByGameID(ctx context.Context, gameID int32) (model.ICSFile, error) {
	var modelICSFile model.ICSFile
	err := f.db.RunTX(ctx, "GetICSFile", func(ctx context.Context) error {
		dbICSFile, err := f.icsFileStorage.GetICSFileByGameID(ctx, int(gameID))
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return fmt.Errorf("get ICS file by game ID error: %w", model.ErrICSFileNotFound)
			}

			return fmt.Errorf("get ICS file by game ID error: %w", err)
		}

		modelICSFile = convertDBICSFileToModelICSFile(*dbICSFile)

		return nil
	})
	if err != nil {
		return model.ICSFile{}, fmt.Errorf("get ICS file by game ID error: %w", err)
	}

	return modelICSFile, nil
}
