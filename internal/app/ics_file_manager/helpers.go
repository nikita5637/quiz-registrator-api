package icsfilemanager

import (
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	icsfilemanagerpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/ics_file_manager"
)

func convertModelICSFileToProtoICSFile(modelICSFile model.ICSFile) *icsfilemanagerpb.ICSFile {
	return &icsfilemanagerpb.ICSFile{
		Id:     modelICSFile.ID,
		GameId: modelICSFile.GameID,
		Name:   modelICSFile.Name,
	}
}
