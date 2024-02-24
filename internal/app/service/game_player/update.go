package gameplayer

import (
	"context"
	"errors"
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/gameplayers"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/games"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/logger"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	gameplayerpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/game_player"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

// UpdatePlayerDegree ...
func (i *Implementation) UpdatePlayerDegree(ctx context.Context, req *gameplayerpb.UpdatePlayerDegreeRequest) (*emptypb.Empty, error) {
	originalGamePlayer, err := i.gamePlayersFacade.GetGamePlayer(ctx, req.GetId())
	if err != nil {
		st := status.New(codes.Internal, err.Error())
		if errors.Is(err, gameplayers.ErrGamePlayerNotFound) {
			st = model.GetStatus(ctx, codes.NotFound, gameplayers.ErrGamePlayerNotFound.Error(), gameplayers.ReasonGamePlayerNotFound, map[string]string{
				"error": err.Error(),
			}, gameplayers.GamePlayerNotFoundLexeme)
		}

		return nil, st.Err()
	}

	if err = validation.Validate(model.Degree(req.GetDegree()), validation.By(model.ValidateDegree)); err != nil {
		st := status.New(codes.InvalidArgument, err.Error())
		if _, ok := err.(validation.InternalError); !ok {
			st = model.GetStatus(ctx, codes.InvalidArgument, fmt.Sprintf("%s %s", "Degree", err.Error()), invalidDegreeReason, nil, invalidDegreeLexeme)
		}

		return nil, st.Err()
	}

	patchedGamePlayer := originalGamePlayer
	patchedGamePlayer.Degree = model.Degree(req.GetDegree())

	logger.DebugKV(ctx, "updating game player degree", zap.Reflect("game_player", patchedGamePlayer))

	game, err := i.gamesFacade.GetGameByID(ctx, patchedGamePlayer.GameID)
	if err != nil {
		st := status.New(codes.Internal, err.Error())
		if errors.Is(err, games.ErrGameNotFound) {
			st = model.GetStatus(ctx, codes.FailedPrecondition, games.ErrGameNotFound.Error(), games.ReasonGameNotFound, map[string]string{
				"error": err.Error(),
			}, games.GameNotFoundLexeme)
		}

		return nil, st.Err()
	}

	if game.HasPassed {
		st := model.GetStatus(ctx, codes.FailedPrecondition, games.ErrGameHasPassed.Error(), games.ReasonGameHasPassed, nil, games.GameHasPassedLexeme)
		return nil, st.Err()
	}

	if !game.Registered {
		st := model.GetStatus(ctx, codes.FailedPrecondition, thereAreNoRegistrationForTheGame, reasonThereAreNoRegistrationForTheGame, nil, thereAreNoRegistrationForTheGameLexeme)
		return nil, st.Err()
	}

	_, err = i.gamePlayersFacade.PatchGamePlayer(ctx, patchedGamePlayer)
	if err != nil {
		st := status.New(codes.Internal, err.Error())
		return nil, st.Err()

	}

	return &emptypb.Empty{}, nil
}
