package game

import (
	"errors"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/mono83/maybe"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/i18n"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	gamepb "github.com/nikita5637/quiz-registrator-api/pkg/pb/game"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type errorDetails struct {
	Reason string
	Lexeme i18n.Lexeme
}

const (
	reasonInvalidExternalID  = "INVALID_EXTERNAL_ID"
	reasonInvalidGameDate    = "INVALID_GAME_DATE"
	reasonInvalidGameName    = "INVALID_GAME_NAME"
	reasonInvalidGameNumber  = "INVALID_GAME_NUMBER"
	reasonInvalidGameType    = "INVALID_GAME_TYPE"
	reasonInvalidLeagueID    = "INVALID_LEAGUE_ID"
	reasonInvalidMaxPlayers  = "INVALID_MAX_PLAYERS"
	reasonInvalidPayment     = "INVALID_PAYMENT"
	reasonInvalidPaymentType = "INVALID_PAYMENT_TYPE"
	reasonInvalidPlaceID     = "INVALID_PLACE_ID"
)

var (
	errorDetailsByField = map[string]errorDetails{
		"Date": {
			Reason: reasonInvalidGameDate,
			Lexeme: invalidGameDateLexeme,
		},
		"ExternalID": {
			Reason: reasonInvalidExternalID,
			Lexeme: invalidExternalIDLexeme,
		},
		"LeagueID": {
			Reason: reasonInvalidLeagueID,
			Lexeme: invalidLeagueIDLexeme,
		},
		"MaxPlayers": {
			Reason: reasonInvalidMaxPlayers,
			Lexeme: invalidMaxPlayersLexeme,
		},
		"Name": {
			Reason: reasonInvalidGameName,
			Lexeme: invalidGameNameLexeme,
		},
		"Number": {
			Reason: reasonInvalidGameNumber,
			Lexeme: invalidGameNumberLexeme,
		},
		"Payment": {
			Reason: reasonInvalidPayment,
			Lexeme: invalidPaymentLexeme,
		},
		"PaymentType": {
			Reason: reasonInvalidPaymentType,
			Lexeme: invalidPaymentTypeLexeme,
		},
		"PlaceID": {
			Reason: reasonInvalidPlaceID,
			Lexeme: invalidPlaceIDLexeme,
		},
		"Type": {
			Reason: reasonInvalidGameType,
			Lexeme: invalidGameTypeLexeme,
		},
	}

	invalidGameDateLexeme = i18n.Lexeme{
		Key:      "invalid_game_date",
		FallBack: "Invalid game date",
	}
	invalidGameNameLexeme = i18n.Lexeme{
		Key:      "invalid_game_name",
		FallBack: "Invalid game name",
	}
	invalidGameNumberLexeme = i18n.Lexeme{
		Key:      "invalid_game_number",
		FallBack: "Invalid game number",
	}
	invalidGameTypeLexeme = i18n.Lexeme{
		Key:      "invalid_game_type",
		FallBack: "Invalid game type",
	}
	invalidExternalIDLexeme = i18n.Lexeme{
		Key:      "invalid_external_id",
		FallBack: "Invalid external ID",
	}
	invalidLeagueIDLexeme = i18n.Lexeme{
		Key:      "invalid_league_id",
		FallBack: "Invalid league ID",
	}
	invalidMaxPlayersLexeme = i18n.Lexeme{
		Key:      "invalid_max_players",
		FallBack: "Invalid max players",
	}
	invalidPaymentLexeme = i18n.Lexeme{
		Key:      "invalid_payment",
		FallBack: "Invalid payment",
	}
	invalidPaymentTypeLexeme = i18n.Lexeme{
		Key:      "invalid_payment_type",
		FallBack: "Invalid payment type",
	}
	invalidPlaceIDLexeme = i18n.Lexeme{
		Key:      "invalid_place_id",
		FallBack: "Invalid place ID",
	}
)

func convertModelGameToProtoGame(game model.Game) *gamepb.Game {
	pbGame := &gamepb.Game{
		Id:         game.ID,
		LeagueId:   game.LeagueID,
		Type:       gamepb.GameType(game.Type),
		Number:     game.Number,
		PlaceId:    game.PlaceID,
		Date:       timestamppb.New(game.DateTime().AsTime()),
		Price:      game.Price,
		MaxPlayers: game.MaxPlayers,
		Registered: game.Registered,
		IsInMaster: game.IsInMaster,
		// additional
		HasPassed: game.HasPassed,
	}

	if externalID, isPresent := game.ExternalID.Get(); isPresent {
		pbGame.ExternalId = wrapperspb.Int32(externalID)
	}

	if name, isPresent := game.Name.Get(); isPresent {
		pbGame.Name = wrapperspb.String(name)
	}

	if paymentType, isPresent := game.PaymentType.Get(); isPresent {
		pbGame.PaymentType = wrapperspb.String(paymentType)
	}

	if payment, isPresent := game.Payment.Get(); isPresent {
		p := gamepb.Payment(payment)
		pbGame.Payment = &p
	}

	if gameLink, isPresent := game.GameLink.Get(); isPresent {
		pbGame.GameLink = wrapperspb.String(gameLink)
	}

	return pbGame
}

func convertProtoGameToModelGame(game *gamepb.Game) model.Game {
	modelGame := model.NewGame()

	modelGame.ID = game.GetId()

	if externalID := game.GetExternalId(); externalID != nil {
		modelGame.ExternalID = maybe.Just(externalID.GetValue())
	}

	modelGame.LeagueID = game.GetLeagueId()
	modelGame.Type = model.GameType(game.GetType())
	modelGame.Number = game.GetNumber()

	if name := game.GetName(); name != nil {
		modelGame.Name = maybe.Just(name.GetValue())
	}

	modelGame.PlaceID = game.GetPlaceId()
	modelGame.Date = model.DateTime(game.GetDate().AsTime())
	modelGame.Price = game.GetPrice()

	if paymentType := game.GetPaymentType(); paymentType != nil {
		modelGame.PaymentType = maybe.Just(paymentType.GetValue())
	}

	modelGame.MaxPlayers = game.GetMaxPlayers()

	if game != nil && game.Payment != nil {
		modelGame.Payment = maybe.Just(model.Payment(game.GetPayment()))
	}

	modelGame.Registered = game.GetRegistered()
	modelGame.IsInMaster = game.GetIsInMaster()
	// additional
	modelGame.HasPassed = game.GetHasPassed()

	if gameLink := game.GetGameLink(); gameLink != nil {
		modelGame.GameLink = maybe.Just(gameLink.GetValue())
	}

	return modelGame
}

func getErrorDetails(keys []string) *errorDetails {
	if len(keys) == 0 {
		return nil
	}

	if v, ok := errorDetailsByField[keys[0]]; ok {
		return &v
	}

	return nil
}

func validateExternalID(value interface{}) error {
	v, ok := value.(maybe.Maybe[int32])
	if !ok {
		return errors.New("must be Maybe[int32]")
	}

	return validation.Validate(v.Value(), validation.When(v.IsPresent(), validation.Required, validation.Min(1)))
}

func validateName(value interface{}) error {
	v, ok := value.(maybe.Maybe[string])
	if !ok {
		return errors.New("must be Maybe[string]")
	}

	return validation.Validate(v.Value(), validation.When(v.IsPresent(), validation.Required, validation.Length(1, 128)))
}

func validatePayment(value interface{}) error {
	v, ok := value.(maybe.Maybe[model.Payment])
	if !ok {
		return errors.New("must be Maybe[model.Payment]")
	}

	return validation.Validate(v.Value(), validation.When(v.IsPresent(), validation.Required, validation.By(model.ValidatePayment)))
}

func validatePaymentType(value interface{}) error {
	v, ok := value.(maybe.Maybe[string])
	if !ok {
		return errors.New("must be Maybe[string]")
	}

	return validation.Validate(v.Value(), validation.When(v.IsPresent(), validation.Required, validation.By(validatePaymentType2)))
}

func validatePaymentType2(value interface{}) error {
	v, ok := value.(string)
	if !ok {
		return errors.New("must be string")
	}

	return validation.Validate(v, validation.In("cash,card", "cash", "card"))
}
