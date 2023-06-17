package icsfiles

import (
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	database "github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
)

func convertDBICSFileToModelICSFile(dbICSFile database.IcsFile) model.ICSFile {
	return model.ICSFile{
		ID:     int32(dbICSFile.ID),
		GameID: int32(dbICSFile.FkGameID),
		Name:   dbICSFile.Name,
	}
}

func convertModelICSFileToDBICSFile(modelICSFile model.ICSFile) database.IcsFile {
	return database.IcsFile{
		ID:       int(modelICSFile.ID),
		FkGameID: int(modelICSFile.GameID),
		Name:     modelICSFile.Name,
	}
}
