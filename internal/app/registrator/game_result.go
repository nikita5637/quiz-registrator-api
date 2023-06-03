package registrator

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/i18n"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	"github.com/nikita5637/quiz-registrator-api/pkg/pb/registrator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

var (
	errInvalidJSONRoundPointsValue = errors.New("invalid JSON round points value")
	errInvalidResultPlace          = errors.New("invalid result place")

	gameResultAlreadyExistsLexeme = i18n.Lexeme{
		Key:      "game_result_already_exists_lexeme",
		FallBack: "Game result already exists",
	}
	gameResultNotFoundLexeme = i18n.Lexeme{
		Key:      "game_result_not_found_lexeme",
		FallBack: "Game result not found",
	}
	invalidGameResultRoundPointsJSONValueLexeme = i18n.Lexeme{
		Key:      "invalid_game_result_round_points_json_value",
		FallBack: "Invalid game result round points JSON value",
	}
	invalidGameResultResultPlaceValueLexeme = i18n.Lexeme{
		Key:      "invalid_game_result_result_place_value",
		FallBack: "Invalid game result result place value",
	}
)

// CreateGameResult ...
func (r *Registrator) CreateGameResult(ctx context.Context, req *registrator.CreateGameResultRequest) (*registrator.GameResult, error) {
	if err := validateCreateGameResultRequest(ctx, req); err != nil {
		st := status.New(codes.InvalidArgument, err.Error())
		if errors.Is(err, errInvalidJSONRoundPointsValue) {
			reason := fmt.Sprintf("invalid game result round points JSON value: \"%s\"", req.GetGameResult().GetRoundPoints())
			st = getStatus(ctx, codes.InvalidArgument, err, reason, invalidGameResultRoundPointsJSONValueLexeme)
		} else if errors.Is(err, errInvalidResultPlace) {
			reason := fmt.Sprintf("invalid game result result place: \"%d\"", req.GetGameResult().GetResultPlace())
			st = getStatus(ctx, codes.InvalidArgument, err, reason, invalidGameResultResultPlaceValueLexeme)
		}

		return nil, st.Err()
	}

	gameResult, err := r.gameResultsFacade.CreateGameResult(ctx, model.GameResult{
		FkGameID:    req.GetGameResult().GetGameId(),
		ResultPlace: req.GetGameResult().GetResultPlace(),
		RoundPoints: model.NewMaybeString(req.GetGameResult().GetRoundPoints()),
	})
	if err != nil {
		st := status.New(codes.Internal, err.Error())
		if errors.Is(err, model.ErrGameNotFound) {
			reason := fmt.Sprintf("game with id %d not found", req.GetGameResult().GetGameId())
			st = getStatus(ctx, codes.InvalidArgument, err, reason, gameNotFoundLexeme)
		} else if errors.Is(err, model.ErrGameResultAlreadyExists) {
			reason := fmt.Sprintf("game result for game id %d already exists", req.GetGameResult().GetGameId())
			st = getStatus(ctx, codes.AlreadyExists, err, reason, gameResultAlreadyExistsLexeme)
		}

		return nil, st.Err()
	}

	return convertModelGameResultToProtoGameResult(gameResult), nil
}

// ListGameResults ...
func (r *Registrator) ListGameResults(ctx context.Context, _ *emptypb.Empty) (*registrator.ListGameResultsResponse, error) {
	gameResults, err := r.gameResultsFacade.ListGameResults(ctx)
	if err != nil {
		st := status.New(codes.Internal, err.Error())
		return nil, st.Err()
	}

	respGameResults := make([]*registrator.GameResult, 0, len(gameResults))
	for _, gameResult := range gameResults {
		respGameResults = append(respGameResults, &registrator.GameResult{
			Id:          gameResult.ID,
			GameId:      gameResult.FkGameID,
			ResultPlace: gameResult.ResultPlace,
			RoundPoints: gameResult.RoundPoints.Value,
		})
	}

	return &registrator.ListGameResultsResponse{
		GameResults: respGameResults,
	}, nil
}

// PatchGameResult ...
func (r *Registrator) PatchGameResult(ctx context.Context, req *registrator.PatchGameResultRequest) (*registrator.GameResult, error) {
	if err := validatePatchGameResultRequest(ctx, req); err != nil {
		st := status.New(codes.InvalidArgument, err.Error())
		if errors.Is(err, errInvalidJSONRoundPointsValue) {
			reason := fmt.Sprintf("invalid game result round points JSON value: \"%s\"", req.GetGameResult().GetRoundPoints())
			st = getStatus(ctx, codes.InvalidArgument, err, reason, invalidGameResultRoundPointsJSONValueLexeme)
		} else if errors.Is(err, errInvalidResultPlace) {
			reason := fmt.Sprintf("invalid game result result place: \"%d\"", req.GetGameResult().GetResultPlace())
			st = getStatus(ctx, codes.InvalidArgument, err, reason, invalidGameResultResultPlaceValueLexeme)
		}

		return nil, st.Err()
	}

	gameResult, err := r.gameResultsFacade.PatchGameResult(ctx, model.GameResult{
		ID:          req.GetGameResult().GetId(),
		FkGameID:    req.GetGameResult().GetGameId(),
		ResultPlace: req.GetGameResult().GetResultPlace(),
		RoundPoints: model.NewMaybeString(req.GetGameResult().GetRoundPoints()),
	}, req.GetUpdateMask().GetPaths())
	if err != nil {
		st := status.New(codes.Internal, err.Error())
		if errors.Is(err, model.ErrGameResultNotFound) {
			reason := fmt.Sprintf("game result with ID %d not found", req.GetGameResult().GetId())
			st = getStatus(ctx, codes.NotFound, err, reason, gameResultNotFoundLexeme)
		} else if errors.Is(err, model.ErrGameNotFound) {
			reason := fmt.Sprintf("game with id %d not found", req.GetGameResult().GetGameId())
			st = getStatus(ctx, codes.InvalidArgument, err, reason, gameNotFoundLexeme)
		} else if errors.Is(err, model.ErrGameResultAlreadyExists) {
			reason := fmt.Sprintf("game result for game id %d already exists", req.GetGameResult().GetGameId())
			st = getStatus(ctx, codes.AlreadyExists, err, reason, gameResultAlreadyExistsLexeme)
		}

		return nil, st.Err()
	}

	return convertModelGameResultToProtoGameResult(gameResult), nil
}

func convertModelGameResultToProtoGameResult(gameResult model.GameResult) *registrator.GameResult {
	return &registrator.GameResult{
		Id:          gameResult.ID,
		GameId:      gameResult.FkGameID,
		ResultPlace: gameResult.ResultPlace,
		RoundPoints: gameResult.RoundPoints.Value,
	}
}

func validateCreateGameResultRequest(ctx context.Context, req *registrator.CreateGameResultRequest) error {
	return validateGameResult(ctx, req.GetGameResult())
}

func validatePatchGameResultRequest(ctx context.Context, req *registrator.PatchGameResultRequest) error {
	return validateGameResult(ctx, req.GetGameResult())
}

func validateGameResult(ctx context.Context, gameResult *registrator.GameResult) error {
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
