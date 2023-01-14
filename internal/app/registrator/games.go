package registrator

import (
	"context"
	"errors"
	"fmt"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/i18n"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/logger"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	"github.com/nikita5637/quiz-registrator-api/pkg/pb/registrator"
	registratorpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/registrator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	invalidDateLexeme = i18n.Lexeme{
		Key:      "invalid_date",
		FallBack: "Invalid date",
	}
	invalidGameNumberLexeme = i18n.Lexeme{
		Key:      "invalid_game_number",
		FallBack: "Invalid game number",
	}
	invalidGameTypeLexeme = i18n.Lexeme{
		Key:      "invalid_game_type",
		FallBack: "Invalid game type",
	}
	invalidLeagueIDLexeme = i18n.Lexeme{
		Key:      "invalid_league_id",
		FallBack: "Invalid league ID",
	}
	invalidMaxPlayersLexeme = i18n.Lexeme{
		Key:      "invalid_max_players",
		FallBack: "Invalid max players",
	}
	invalidPlaceIDLexeme = i18n.Lexeme{
		Key:      "invalid_place_id",
		FallBack: "Invalid place ID",
	}
	invalidPriceLexeme = i18n.Lexeme{
		Key:      "invalid_price",
		FallBack: "Invalid price",
	}
)

// AddGame ...
func (r *Registrator) AddGame(ctx context.Context, req *registrator.AddGameRequest) (*registrator.AddGameResponse, error) {
	game := model.Game{
		ExternalID:  req.GetExternalId(),
		LeagueID:    req.GetLeagueId(),
		Type:        int32(req.GetGameType()),
		Number:      req.GetNumber(),
		Name:        req.GetName(),
		PlaceID:     req.GetPlaceId(),
		Price:       req.GetPrice(),
		PaymentType: req.GetPaymentType(),
		MaxPlayers:  req.GetMaxPlayers(),
	}

	if req.GetDate() != nil {
		game.Date = model.DateTime(req.GetDate().AsTime())
	}

	id, err := r.gamesFacade.AddGame(ctx, game)
	if err != nil {
		st := status.New(codes.Internal, err.Error())
		if errors.Is(err, model.ErrInvalidLeagueID) {
			reason := fmt.Sprintf("invalid league ID: %d", req.GetLeagueId())
			st = getStatus(ctx, codes.InvalidArgument, err, reason, invalidLeagueIDLexeme)
		} else if errors.Is(err, model.ErrInvalidGameType) {
			reason := fmt.Sprintf("invalid game type: %d", req.GetGameType())
			st = getStatus(ctx, codes.InvalidArgument, err, reason, invalidGameTypeLexeme)
		} else if errors.Is(err, model.ErrInvalidGameNumber) {
			reason := fmt.Sprintf("invalid game number: %s", req.GetNumber())
			st = getStatus(ctx, codes.InvalidArgument, err, reason, invalidGameNumberLexeme)
		} else if errors.Is(err, model.ErrInvalidPlaceID) {
			reason := fmt.Sprintf("invalid place ID: %d", req.GetPlaceId())
			st = getStatus(ctx, codes.InvalidArgument, err, reason, invalidPlaceIDLexeme)
		} else if errors.Is(err, model.ErrInvalidDate) {
			reason := fmt.Sprintf("invalid date: %s", req.GetDate())
			st = getStatus(ctx, codes.InvalidArgument, err, reason, invalidDateLexeme)
		} else if errors.Is(err, model.ErrInvalidPrice) {
			reason := fmt.Sprintf("invalid price: %d", req.GetPrice())
			st = getStatus(ctx, codes.InvalidArgument, err, reason, invalidPriceLexeme)
		} else if errors.Is(err, model.ErrInvalidMaxPlayers) {
			reason := fmt.Sprintf("invalid max players: %d", req.GetMaxPlayers())
			st = getStatus(ctx, codes.InvalidArgument, err, reason, invalidMaxPlayersLexeme)
		}

		return nil, st.Err()
	}

	return &registrator.AddGameResponse{
		Id: id,
	}, nil
}

// AddGames ...
func (r *Registrator) AddGames(ctx context.Context, req *registrator.AddGamesRequest) (*registrator.AddGamesResponse, error) {
	games := make([]model.Game, 0, len(req.GetGames()))
	for _, pbGame := range req.GetGames() {
		game := model.Game{
			ExternalID:  pbGame.GetExternalId(),
			LeagueID:    pbGame.GetLeagueId(),
			Type:        int32(pbGame.GetGameType()),
			Number:      pbGame.GetNumber(),
			Name:        pbGame.GetName(),
			PlaceID:     pbGame.GetPlaceId(),
			Price:       pbGame.GetPrice(),
			PaymentType: pbGame.GetPaymentType(),
			MaxPlayers:  pbGame.GetMaxPlayers(),
		}

		if pbGame.GetDate() != nil {
			game.Date = model.DateTime(pbGame.GetDate().AsTime())
		}

		if err := model.ValidateGame(game); err != nil {
			logger.WarnKV(ctx, "skipped game", "error", err.Error(), "game", game)
		} else {
			games = append(games, game)
		}
	}

	err := r.gamesFacade.AddGames(ctx, games)
	if err != nil {
		st := status.New(codes.Internal, err.Error())
		return nil, st.Err()
	}

	return &registrator.AddGamesResponse{}, nil
}

// DeleteGame ...
func (r *Registrator) DeleteGame(ctx context.Context, req *registrator.DeleteGameRequest) (*registrator.DeleteGameResponse, error) {
	err := r.gamesFacade.DeleteGame(ctx, req.GetId())
	if err != nil {
		st := status.New(codes.Internal, err.Error())

		if errors.Is(err, model.ErrGameNotFound) {
			st = getGameNotFoundStatus(ctx, err, req.GetId())
		}

		return nil, st.Err()
	}

	return &registrator.DeleteGameResponse{}, nil
}

// GetGameByID returns game or Not Found
func (r *Registrator) GetGameByID(ctx context.Context, req *registrator.GetGameByIDRequest) (*registrator.GetGameByIDResponse, error) {
	game, err := r.gamesFacade.GetGameByID(ctx, req.GetGameId())
	if err != nil {
		st := status.New(codes.Internal, err.Error())

		if errors.Is(err, model.ErrGameNotFound) {
			st = getGameNotFoundStatus(ctx, err, req.GetGameId())
		}

		return nil, st.Err()
	}

	return &registrator.GetGameByIDResponse{
		Game: convertModelGameToPBGame(game),
	}, nil
}

// GetGames ...
func (r *Registrator) GetGames(ctx context.Context, req *registrator.GetGamesRequest) (*registrator.GetGamesResponse, error) {
	allGames, err := r.gamesFacade.GetGames(ctx)
	if err != nil {
		st := status.New(codes.Internal, err.Error())
		return nil, st.Err()
	}

	games := allGames
	if req.GetActive() {
		games = make([]model.Game, 0)
		for _, game := range allGames {
			if game.IsActive() {
				games = append(games, game)
			}
		}
	}

	pbGames := make([]*registrator.Game, 0, len(games))
	for _, game := range games {
		pbGames = append(pbGames, convertModelGameToPBGame(game))
	}

	return &registrator.GetGamesResponse{
		Games: pbGames,
	}, nil
}

// GetPlayersByGameID ...
func (r *Registrator) GetPlayersByGameID(ctx context.Context, req *registrator.GetPlayersByGameIDRequest) (*registrator.GetPlayersByGameIDResponse, error) {
	players, err := r.gamesFacade.GetPlayersByGameID(ctx, req.GetGameId())
	if err != nil {
		st := status.New(codes.Internal, err.Error())

		if errors.Is(err, model.ErrGameNotFound) {
			st = getGameNotFoundStatus(ctx, err, req.GetGameId())
		}

		return nil, st.Err()
	}

	pbPlayers := make([]*registrator.Player, 0, len(players))
	for _, player := range players {
		pbPlayers = append(pbPlayers, &registrator.Player{
			UserId:       player.FkUserID,
			RegisteredBy: player.RegisteredBy,
			Degree:       registrator.Degree(player.Degree),
		})
	}

	return &registrator.GetPlayersByGameIDResponse{
		Players: pbPlayers,
	}, nil
}

// GetRegisteredGames ...
func (r *Registrator) GetRegisteredGames(ctx context.Context, req *registrator.GetRegisteredGamesRequest) (*registrator.GetRegisteredGamesResponse, error) {
	allGames, err := r.gamesFacade.GetRegisteredGames(ctx)
	if err != nil {
		st := status.New(codes.Internal, err.Error())
		return nil, st.Err()
	}

	games := allGames
	if req.GetActive() {
		games = make([]model.Game, 0)
		for _, game := range allGames {
			if game.IsActive() {
				games = append(games, game)
			}
		}

	}

	pbGames := make([]*registrator.Game, 0, len(games))
	for _, game := range games {
		pbGames = append(pbGames, convertModelGameToPBGame(game))
	}

	return &registrator.GetRegisteredGamesResponse{
		Games: pbGames,
	}, nil
}

// GetUserGames ...
func (r *Registrator) GetUserGames(ctx context.Context, req *registrator.GetUserGamesRequest) (*registrator.GetUserGamesResponse, error) {
	allGames, err := r.gamesFacade.GetGamesByUserID(ctx, req.GetUserId())
	if err != nil {
		st := status.New(codes.Internal, err.Error())
		return nil, st.Err()
	}

	games := allGames
	if req.GetActive() {
		games = make([]model.Game, 0)
		for _, game := range allGames {
			if game.IsActive() {
				games = append(games, game)
			}
		}

	}

	pbGames := make([]*registrator.Game, 0, len(games))
	for _, game := range games {
		pbGames = append(pbGames, convertModelGameToPBGame(game))
	}

	return &registrator.GetUserGamesResponse{
		Games: pbGames,
	}, nil
}

func convertModelGameToPBGame(game model.Game) *registratorpb.Game {
	pbGame := &registratorpb.Game{
		Id:          game.ID,
		ExternalId:  game.ExternalID,
		LeagueId:    game.LeagueID,
		Type:        registratorpb.GameType(game.Type),
		Number:      game.Number,
		Name:        game.Name,
		PlaceId:     game.PlaceID,
		Date:        timestamppb.New(game.DateTime().AsTime()),
		Price:       game.Price,
		PaymentType: game.PaymentType,
		MaxPlayers:  game.MaxPlayers,
		Payment:     registratorpb.Payment(game.Payment),
		Registered:  game.Registered,
	}

	pbGame.My = game.My
	pbGame.NumberOfMyLegioners = game.NumberOfMyLegioners
	pbGame.NumberOfLegioners = game.NumberOfLegioners
	pbGame.NumberOfPlayers = game.NumberOfPlayers
	pbGame.ResultPlace = game.ResultPlace

	return pbGame
}
