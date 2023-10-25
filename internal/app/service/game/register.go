package game

import (
	"context"
	"errors"

	"github.com/mono83/maybe"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/games"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/logger"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	"github.com/nikita5637/quiz-registrator-api/pkg/ics"
	gamepb "github.com/nikita5637/quiz-registrator-api/pkg/pb/game"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

// RegisterGame ...
func (i *Implementation) RegisterGame(ctx context.Context, req *gamepb.RegisterGameRequest) (*emptypb.Empty, error) {
	game, err := i.gamesFacade.GetGameByID(ctx, req.GetId())
	if err != nil {
		st := status.New(codes.Internal, err.Error())
		if errors.Is(err, games.ErrGameNotFound) {
			st = model.GetStatus(ctx, codes.NotFound, games.ErrGameNotFound.Error(), games.ReasonGameNotFound, map[string]string{
				"error": err.Error(),
			}, games.GameNotFoundLexeme)
		}

		return nil, st.Err()
	}

	if game.HasPassed {
		st := model.GetStatus(ctx, codes.FailedPrecondition, games.ErrGameHasPassed.Error(), games.ReasonGameHasPassed, nil, games.GameHasPassedLexeme)
		return nil, st.Err()
	}

	game.Registered = true
	game.Payment = maybe.Just(model.PaymentCash)

	if _, err = i.gamesFacade.PatchGame(ctx, game); err != nil {
		st := status.New(codes.Internal, err.Error())
		return nil, st.Err()
	}

	icsEvent := ics.Event{
		GameID: game.ID,
		Event:  ics.EventRegistered,
	}
	if err := i.rabbitMQProducer.Send(ctx, icsEvent); err != nil {
		logger.ErrorKV(ctx, "sending ICS event error", zap.Error(err), "event", icsEvent)
	}

	return &emptypb.Empty{}, nil
}
