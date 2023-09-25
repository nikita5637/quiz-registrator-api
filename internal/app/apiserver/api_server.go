package apiserver

import (
	"context"
	"net"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/nikita5637/quiz-registrator-api/internal/app/middleware/authentication"
	"github.com/nikita5637/quiz-registrator-api/internal/app/middleware/authorization"
	errorwrap "github.com/nikita5637/quiz-registrator-api/internal/app/middleware/error_wrap"
	"github.com/nikita5637/quiz-registrator-api/internal/app/middleware/log"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/logger"
	adminpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/admin"
	certificatemanagerpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/certificate_manager"
	croupierpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/croupier"
	gamepb "github.com/nikita5637/quiz-registrator-api/pkg/pb/game"
	gameplayerpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/game_player"
	gameresultmanagerpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/game_result_manager"
	leaguepb "github.com/nikita5637/quiz-registrator-api/pkg/pb/league"
	photomanagerpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/photo_manager"
	placepb "github.com/nikita5637/quiz-registrator-api/pkg/pb/place"
	usermanagerpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/user_manager"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// APIServer ...
type APIServer struct {
	grpcServer *grpc.Server
}

// Config ...
type Config struct {
	// middlewares
	AuthenticationMiddleware *authentication.Middleware
	AuthorizationMiddleware  *authorization.Middleware
	ErrorWrapMiddleware      *errorwrap.Middleware
	LogMiddleware            *log.Middleware

	// services
	AdminService                 adminpb.ServiceServer
	CertificateManagerService    certificatemanagerpb.ServiceServer
	CroupierService              croupierpb.ServiceServer
	GameService                  gamepb.ServiceServer
	GameRegistratorService       gamepb.RegistratorServiceServer
	GamePlayerService            gameplayerpb.ServiceServer
	GamePlayerRegistratorService gameplayerpb.RegistratorServiceServer
	GameResultManagerService     gameresultmanagerpb.ServiceServer
	LeagueService                leaguepb.ServiceServer
	PhotoManagerService          photomanagerpb.ServiceServer
	PlaceService                 placepb.ServiceServer
	UserManagerService           usermanagerpb.ServiceServer
}

// New ...
func New(cfg Config) *APIServer {
	var opts []grpc.ServerOption
	opts = append(opts, grpc.ChainUnaryInterceptor(
		grpc_recovery.UnaryServerInterceptor(),
		cfg.LogMiddleware.Log(),
		cfg.ErrorWrapMiddleware.ErrorWrap(),
		grpc_auth.UnaryServerInterceptor(cfg.AuthenticationMiddleware.Authentication()),
		cfg.AuthorizationMiddleware.Authorization(),
	))

	s := grpc.NewServer(opts...)
	reflection.Register(s)

	adminpb.RegisterServiceServer(s, cfg.AdminService)
	certificatemanagerpb.RegisterServiceServer(s, cfg.CertificateManagerService)
	croupierpb.RegisterServiceServer(s, cfg.CroupierService)
	gamepb.RegisterServiceServer(s, cfg.GameService)
	gamepb.RegisterRegistratorServiceServer(s, cfg.GameRegistratorService)
	gameplayerpb.RegisterServiceServer(s, cfg.GamePlayerService)
	gameplayerpb.RegisterRegistratorServiceServer(s, cfg.GamePlayerRegistratorService)
	gameresultmanagerpb.RegisterServiceServer(s, cfg.GameResultManagerService)
	leaguepb.RegisterServiceServer(s, cfg.LeagueService)
	photomanagerpb.RegisterServiceServer(s, cfg.PhotoManagerService)
	placepb.RegisterServiceServer(s, cfg.PlaceService)
	usermanagerpb.RegisterServiceServer(s, cfg.UserManagerService)

	return &APIServer{
		grpcServer: s,
	}
}

// ListenAndServe ...
func (s *APIServer) ListenAndServe(ctx context.Context, lis net.Listener) error {
	var err error
	go func() {
		err = s.grpcServer.Serve(lis)
		return
	}()
	if err != nil {
		return err
	}

	<-ctx.Done()

	s.grpcServer.GracefulStop()

	logger.Info(ctx, "registrator gracefully stopped")
	return nil
}
