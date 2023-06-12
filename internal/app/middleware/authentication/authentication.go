package authentication

import (
	"context"
	"strconv"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	usersutils "github.com/nikita5637/quiz-registrator-api/utils/users"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const (
	unauthenticated = "unauthenticated"
)

// Authentication ...
func (m *Middleware) Authentication() grpc_auth.AuthFunc {
	return func(ctx context.Context) (context.Context, error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return ctx, status.New(codes.Unauthenticated, unauthenticated).Err()
		}

		switch getAuthenticationType(ctx) {
		case authenticationTypeTelegramID:
			telegramClientID := ""

			telegramClientIDs := md.Get(telegramClientIDHeader)
			telegramClientID = telegramClientIDs[0]

			id, err := strconv.ParseInt(telegramClientID, 10, 64)
			if err != nil {
				return ctx, status.New(codes.Unauthenticated, unauthenticated).Err()
			}

			user, err := m.usersFacade.GetUserByTelegramID(ctx, id)
			if err != nil {
				return ctx, status.New(codes.Unauthenticated, unauthenticated).Err()
			}

			return usersutils.NewContextWithUser(ctx, user), nil
		case authenticationTypeServiceName:
			serviceName := ""

			serviceNames := md.Get(serviceNameHeader)
			if len(serviceNames) > 0 {
				serviceName = serviceNames[0]
			}

			if serviceName == "" {
				moduleNames := md.Get(moduleNameHeader)
				if len(moduleNames) > 0 {
					serviceName = moduleNames[0]
				}
			}

			if serviceName != "fetcher" &&
				serviceName != "telegram" {
				return ctx, status.New(codes.Unauthenticated, unauthenticated).Err()
			}

			return ctx, nil
		}

		return ctx, status.New(codes.Unauthenticated, unauthenticated).Err()
	}
}
