//go:generate mockery --case underscore --name Croupier --with-expecter
//go:generate mockery --case underscore --name GamesFacade --with-expecter
//go:generate mockery --case underscore --name GamePhotosFacade --with-expecter
//go:generate mockery --case underscore --name GameResultsFacade --with-expecter
//go:generate mockery --case underscore --name LeaguesFacade --with-expecter
//go:generate mockery --case underscore --name PlacesFacade --with-expecter
//go:generate mockery --case underscore --name UsersFacade --with-expecter

package registrator

import (
	"context"
	"fmt"
	"net"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/logger"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	certificatemanagerpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/certificate_manager"
	registratorpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/registrator"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Croupier ...
type Croupier interface {
	GetIsLotteryActive(ctx context.Context, game model.Game) bool
	RegisterForLottery(ctx context.Context, game model.Game, user model.User) (int32, error)
}

// GamesFacade ...
type GamesFacade interface {
	AddGame(ctx context.Context, game model.Game) (int32, error)
	AddGames(ctx context.Context, games []model.Game) error
	DeleteGame(ctx context.Context, gameID int32) error
	// GetGameByID guaranteed returns active game by game ID
	GetGameByID(ctx context.Context, id int32) (model.Game, error)
	GetGames(ctx context.Context) ([]model.Game, error)
	GetGamesByUserID(ctx context.Context, userID int32) ([]model.Game, error)
	GetPlayersByGameID(ctx context.Context, gameID int32) ([]model.GamePlayer, error)
	GetRegisteredGames(ctx context.Context) ([]model.Game, error)
	RegisterGame(ctx context.Context, gameID int32) (model.RegisterGameStatus, error)
	RegisterPlayer(ctx context.Context, fkGameID, fkUserID, registeredBy, degree int32) (model.RegisterPlayerStatus, error)
	UnregisterGame(ctx context.Context, gameID int32) (model.UnregisterGameStatus, error)
	UnregisterPlayer(ctx context.Context, gameID, userID, deletedBy int32) (model.UnregisterPlayerStatus, error)
	UpdatePayment(ctx context.Context, gameID int32, payment int32) error
}

// GamePhotosFacade ...
type GamePhotosFacade interface {
	AddGamePhotos(ctx context.Context, gameID int32, urls []string) error
	GetGamesWithPhotos(ctx context.Context, limit, offset uint32) ([]model.Game, error)
	GetNumberOfGamesWithPhotos(ctx context.Context) (uint32, error)
	GetPhotosByGameID(ctx context.Context, gameID int32) ([]string, error)
}

// GameResultsFacade ...
type GameResultsFacade interface {
	CreateGameResult(ctx context.Context, gameResult model.GameResult) (model.GameResult, error)
	ListGameResults(ctx context.Context) ([]model.GameResult, error)
	PatchGameResult(ctx context.Context, gameResult model.GameResult, paths []string) (model.GameResult, error)
}

// LeaguesFacade ...
type LeaguesFacade interface {
	GetLeagueByID(ctx context.Context, leagueID int32) (model.League, error)
}

// PlacesFacade ...
type PlacesFacade interface {
	GetPlaceByID(ctx context.Context, placeID int32) (model.Place, error)
	GetPlaceByNameAndAddress(ctx context.Context, name, address string) (model.Place, error)
}

// UsersFacade ...
type UsersFacade interface {
	CreateUser(ctx context.Context, user model.User) (int32, error)
	GetUser(ctx context.Context) (model.User, error)
	GetUserByID(ctx context.Context, userID int32) (model.User, error)
	GetUserByTelegramID(ctx context.Context, telegramID int64) (model.User, error)
	UpdateUserEmail(ctx context.Context, userID int32, email string) error
	UpdateUserName(ctx context.Context, userID int32, name string) error
	UpdateUserPhone(ctx context.Context, userID int32, phone string) error
	UpdateUserState(ctx context.Context, userID int32, state model.UserState) error
}

// Registrator ...
type Registrator struct {
	bindAddr   string
	grpcServer *grpc.Server

	croupier Croupier

	certificateManager certificatemanagerpb.ServiceServer
	gamesFacade        GamesFacade
	gamePhotosFacade   GamePhotosFacade
	gameResultsFacade  GameResultsFacade
	leaguesFacade      LeaguesFacade
	placesFacade       PlacesFacade
	usersFacade        UsersFacade

	registratorpb.UnimplementedCroupierServiceServer
	registratorpb.UnimplementedPhotographerServiceServer
	registratorpb.UnimplementedRegistratorServiceServer
}

// Config ...
type Config struct {
	BindAddr string

	Croupier Croupier

	CertificateManager certificatemanagerpb.ServiceServer
	GamesFacade        GamesFacade
	GamePhotosFacade   GamePhotosFacade
	GameResultsFacade  GameResultsFacade
	LeaguesFacade      LeaguesFacade
	PlacesFacade       PlacesFacade
	UsersFacade        UsersFacade
}

// New ...
func New(cfg Config) *Registrator {
	registrator := &Registrator{
		bindAddr: cfg.BindAddr,

		croupier: cfg.Croupier,

		certificateManager: cfg.CertificateManager,
		gamesFacade:        cfg.GamesFacade,
		gamePhotosFacade:   cfg.GamePhotosFacade,
		gameResultsFacade:  cfg.GameResultsFacade,
		leaguesFacade:      cfg.LeaguesFacade,
		placesFacade:       cfg.PlacesFacade,
		usersFacade:        cfg.UsersFacade,
	}

	var opts []grpc.ServerOption
	opts = append(opts, grpc.ChainUnaryInterceptor(
		grpc_recovery.UnaryServerInterceptor(),
		logInterceptor,
		grpc_auth.UnaryServerInterceptor(registrator.AuthFunc),
		userStateInterceptor,
	))
	s := grpc.NewServer(opts...)
	reflection.Register(s)

	certificatemanagerpb.RegisterServiceServer(s, registrator.certificateManager)
	registratorpb.RegisterCroupierServiceServer(s, registrator)
	registratorpb.RegisterPhotographerServiceServer(s, registrator)
	registratorpb.RegisterRegistratorServiceServer(s, registrator)

	registrator.grpcServer = s

	return registrator
}

// ListenAndServe ...
func (r *Registrator) ListenAndServe(ctx context.Context) error {
	var err error
	var lis net.Listener

	lis, err = net.Listen("tcp", r.bindAddr)
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}

	go func() {
		err = r.grpcServer.Serve(lis)
		return
	}()
	if err != nil {
		return err
	}

	<-ctx.Done()

	r.grpcServer.GracefulStop()

	logger.Info(ctx, "registrator gracefully stopped")
	return nil
}
