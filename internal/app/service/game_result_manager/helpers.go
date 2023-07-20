package gameresultmanager

import (
	"context"
	"encoding/json"
	"errors"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/i18n"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	gameresultmanagerpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/game_result_manager"
)

var (
	errInvalidJSONRoundPointsValue = errors.New("invalid JSON round points value")
	errInvalidResultPlace          = errors.New("invalid result place")

	gameResultAlreadyExistsLexeme = i18n.Lexeme{
		Key:      "game_result_already_exists_lexeme",
		FallBack: "Game result already exists",
	}
	invalidGameResultResultPlaceValueLexeme = i18n.Lexeme{
		Key:      "invalid_game_result_result_place_value",
		FallBack: "Invalid game result result place value",
	}
	invalidGameResultRoundPointsJSONValueLexeme = i18n.Lexeme{
		Key:      "invalid_game_result_round_points_json_value",
		FallBack: "Invalid game result round points JSON value",
	}
)

func convertModelGameResultToProtoGameResult(gameResult model.GameResult) *gameresultmanagerpb.GameResult {
	return &gameresultmanagerpb.GameResult{
		Id:          gameResult.ID,
		GameId:      gameResult.FkGameID,
		ResultPlace: gameResult.ResultPlace,
		RoundPoints: gameResult.RoundPoints.Value(),
	}
}

func validateGameResult(ctx context.Context, gameResult *gameresultmanagerpb.GameResult) error {
	if valid := json.Valid([]byte(gameResult.GetRoundPoints())); !valid {
		return errInvalidJSONRoundPointsValue
	}

	err := validation.Validate(gameResult.GetResultPlace(), validation.Required)
	if err != nil {
		return errInvalidResultPlace
	}

	if len(gameResult.GetRoundPoints()) > 256 {
		return errInvalidJSONRoundPointsValue
	}

	return nil
}
