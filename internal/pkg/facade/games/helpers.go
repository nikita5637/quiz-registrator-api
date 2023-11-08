package games

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/mono83/maybe"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	database "github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
	timeutils "github.com/nikita5637/quiz-registrator-api/utils/time"
	"github.com/spf13/viper"
)

const (
	// league_id field
	leagueIBFK1ConstraintName = "league_ibfk_1"
)

func convertDBGameToModelGame(game database.Game) model.Game {
	modelGame := model.NewGame()

	modelGame.ID = int32(game.ID)

	if game.ExternalID.Valid {
		modelGame.ExternalID = maybe.Just(int32(game.ExternalID.Int64))
	}

	modelGame.LeagueID = int32(game.LeagueID)
	modelGame.Type = model.GameType(game.Type)
	modelGame.Number = game.Number

	if game.Name.Valid {
		modelGame.Name = maybe.Just(game.Name.String)
	}

	modelGame.PlaceID = int32(game.PlaceID)
	modelGame.Date = model.DateTime(game.Date)
	modelGame.Price = uint32(game.Price)

	if game.PaymentType != nil {
		modelGame.PaymentType = maybe.Just(string(game.PaymentType))
	}

	modelGame.MaxPlayers = uint32(game.MaxPlayers)

	if game.Payment.Valid {
		modelGame.Payment = maybe.Just(model.Payment(game.Payment.Int64))
	}

	modelGame.Registered = game.Registered
	modelGame.IsInMaster = game.IsInMaster

	return modelGame
}

func convertModelGameToDBGame(game model.Game) database.Game {
	databaseGame := database.Game{
		ID: int(game.ID),
		ExternalID: sql.NullInt64{
			Int64: int64(game.ExternalID.Value()),
			Valid: game.ExternalID.IsPresent(),
		},
		LeagueID: int(game.LeagueID),
		Type:     uint8(game.Type),
		Number:   game.Number,
		Name: sql.NullString{
			String: game.Name.Value(),
			Valid:  game.Name.IsPresent(),
		},
		PlaceID:    int(game.PlaceID),
		Date:       game.Date.AsTime(),
		Price:      uint(game.Price),
		MaxPlayers: uint8(game.MaxPlayers),
		Payment: sql.NullInt64{
			Int64: int64(game.Payment.Value()),
			Valid: game.Payment.IsPresent(),
		},
		Registered: game.Registered,
		IsInMaster: game.IsInMaster,
	}

	if paymentType, isPresent := game.PaymentType.Get(); isPresent {
		databaseGame.PaymentType = []byte(paymentType)
	}

	return databaseGame
}

func gameHasPassed(modelGame model.Game) bool {
	hasPassedGameLag := viper.GetDuration("service.game.has_passed_game_lag")
	return timeutils.TimeNow().UTC().After(modelGame.DateTime().AsTime().Add(hasPassedGameLag * time.Second))
}

func getGameLink(modelGame model.Game) string {
	externalID, isPresent := modelGame.ExternalID.Get()
	if !isPresent {
		return ""
	}

	switch modelGame.LeagueID {
	case model.LeagueQuizPlease:
		return fmt.Sprintf("https://spb.quizplease.ru/game-page?id=%d", externalID)
	case model.LeagueSixtySeconds:
		return fmt.Sprintf("https://60sec.online/game/%d/", externalID)
	}

	return ""
}
