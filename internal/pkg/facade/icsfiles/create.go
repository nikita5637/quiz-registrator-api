package icsfiles

import (
	"context"
	"fmt"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
)

// CreateICSFile ...
func (f *Facade) CreateICSFile(ctx context.Context, icsFile model.ICSFile) (model.ICSFile, error) {
	createdModelICSFile := model.ICSFile{}
	err := f.db.RunTX(ctx, "CreateICSFile", func(ctx context.Context) error {
		newDBICSFile := convertModelICSFileToDBICSFile(icsFile)
		id, err := f.icsFileStorage.CreateICSFile(ctx, newDBICSFile)
		if err != nil {
			if err, ok := err.(*mysql.MySQLError); ok {
				if err.Number == 1452 {
					return fmt.Errorf("create ICS file error: %w", model.ErrGameNotFound)
				} else if err.Number == 1062 {
					if i := strings.Index(err.Message, "for key 'name'"); i != -1 {
						return fmt.Errorf("create ICS file error: %w", model.ErrICSFileNameAlreadyExists)
					}
					return fmt.Errorf("create ICS file error: %w", model.ErrICSFileForGameAlreadyExists)
				}
			}

			return fmt.Errorf("create ICS file error: %w", err)
		}

		newDBICSFile.ID = id
		createdModelICSFile = convertDBICSFileToModelICSFile(newDBICSFile)

		return nil
	})
	if err != nil {
		return model.ICSFile{}, fmt.Errorf("create ICS file error: %w", err)
	}

	return createdModelICSFile, nil
}
