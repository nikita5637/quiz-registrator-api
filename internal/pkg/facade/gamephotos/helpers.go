package gamephotos

import (
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	database "github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
)

func convertDBGamePhotoToModelGamePhoto(game database.GamePhoto) model.GamePhoto {
	return model.GamePhoto{
		ID:       int32(game.ID),
		FkGameID: int32(game.FkGameID),
		URL:      game.URL,
	}
}

func convertModelGamePhotoToDBGamePhoto(game model.GamePhoto) database.GamePhoto {
	return database.GamePhoto{
		ID:       int(game.ID),
		FkGameID: int(game.FkGameID),
		URL:      game.URL,
	}
}
