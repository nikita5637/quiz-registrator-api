package authentication

import (
	"context"

	"google.golang.org/grpc/metadata"
)

func getAuthenticationType(ctx context.Context) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ""
	}

	telegramClientIDs := md.Get(telegramClientIDHeader)
	if len(telegramClientIDs) > 0 {
		return authenticationTypeTelegramID
	}

	serviceNames := md.Get(serviceNameHeader)
	if len(serviceNames) > 0 {
		return authenticationTypeServiceName
	}

	moduleNames := md.Get(moduleNameHeader)
	if len(moduleNames) > 0 {
		return authenticationTypeServiceName
	}

	return ""
}
