package gameplayer

import (
	"errors"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/mono83/maybe"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/i18n"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	gameplayerpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/game_player"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type errorDetails struct {
	Reason string
	Lexeme i18n.Lexeme
}

const (
	invalidGameIDReason       = "INVALID_GAME_ID"
	invalidUserIDReason       = "INVALID_USER_ID"
	invalidRegisteredByReason = "INVALID_REGISTERED_BY"
	invalidDegreeReason       = "INVALID_DEGREE"
)

var (
	invalidGameIDLexeme = i18n.Lexeme{
		Key:      "invalid_game_id",
		FallBack: "Invalid game ID",
	}
	invalidUserIDLexeme = i18n.Lexeme{
		Key:      "invalid_user_id",
		FallBack: "Invalid user ID",
	}
	invalidRegisteredByLexeme = i18n.Lexeme{
		Key:      "invalid_registered_by",
		FallBack: "Invalid registered by",
	}
	invalidDegreeLexeme = i18n.Lexeme{
		Key:      "invalid_degree",
		FallBack: "Invalid degree",
	}

	errorDetailsByField = map[string]errorDetails{
		"GameID": {
			Reason: invalidGameIDReason,
			Lexeme: invalidGameIDLexeme,
		},
		"UserID": {
			Reason: invalidUserIDReason,
			Lexeme: invalidUserIDLexeme,
		},
		"RegisteredBy": {
			Reason: invalidRegisteredByReason,
			Lexeme: invalidRegisteredByLexeme,
		},
		"Degree": {
			Reason: invalidDegreeReason,
			Lexeme: invalidDegreeLexeme,
		},
	}
)

func convertModelGamePlayerToProtoGamePlayer(gamePlayer model.GamePlayer) *gameplayerpb.GamePlayer {
	ret := &gameplayerpb.GamePlayer{
		Id:           gamePlayer.ID,
		GameId:       gamePlayer.GameID,
		RegisteredBy: gamePlayer.RegisteredBy,
		Degree:       gameplayerpb.Degree(gamePlayer.Degree),
	}
	if v, ok := gamePlayer.UserID.Get(); ok {
		ret.UserId = &wrapperspb.Int32Value{
			Value: v,
		}
	}

	return ret
}

func convertProtoGamePlayerToModelGamePlayer(gamePlayer *gameplayerpb.GamePlayer) model.GamePlayer {
	ret := model.GamePlayer{
		ID:           gamePlayer.GetId(),
		GameID:       gamePlayer.GetGameId(),
		UserID:       maybe.Nothing[int32](),
		RegisteredBy: gamePlayer.GetRegisteredBy(),
		Degree:       model.Degree(gamePlayer.GetDegree()),
	}
	if gamePlayer.GetUserId() != nil {
		ret.UserID = maybe.Just(gamePlayer.GetUserId().GetValue())
	}

	return ret
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

func validateUserID(value interface{}) error {
	v, ok := value.(maybe.Maybe[int32])
	if !ok {
		return errors.New("must be Maybe[int32]")
	}

	return validation.Validate(v.Value(), validation.When(v.IsPresent(), validation.Required))
}
