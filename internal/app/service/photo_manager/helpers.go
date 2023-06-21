package photomanager

import (
	"context"
	"fmt"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/i18n"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	commonpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/common"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertModelGameToPBGame(game model.Game) *commonpb.Game {
	pbGame := &commonpb.Game{
		Id:          game.ID,
		ExternalId:  game.ExternalID,
		LeagueId:    game.LeagueID,
		Type:        commonpb.GameType(game.Type),
		Number:      game.Number,
		Name:        game.Name,
		PlaceId:     game.PlaceID,
		Date:        timestamppb.New(game.DateTime().AsTime()),
		Price:       game.Price,
		PaymentType: game.PaymentType,
		MaxPlayers:  game.MaxPlayers,
		Payment:     commonpb.Payment(game.Payment),
		Registered:  game.Registered,
	}

	pbGame.My = game.My
	pbGame.NumberOfMyLegioners = game.NumberOfMyLegioners
	pbGame.NumberOfLegioners = game.NumberOfLegioners
	pbGame.NumberOfPlayers = game.NumberOfPlayers
	pbGame.ResultPlace = game.ResultPlace

	return pbGame
}

func getGameNotFoundStatus(ctx context.Context, err error, gameID int32) *status.Status {
	reason := fmt.Sprintf("game with id %d not found", gameID)
	return model.GetStatus(ctx, codes.NotFound, err, reason, i18n.GameNotFoundLexeme)
}
