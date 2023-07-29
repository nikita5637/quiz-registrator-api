package authorization

import (
	"context"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/i18n"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/logger"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	users_utils "github.com/nikita5637/quiz-registrator-api/utils/users"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	permissionDeniedLexeme = i18n.Lexeme{
		Key:      "permission_denied",
		FallBack: "Permission denied",
	}
	youAreBannedLexeme = i18n.Lexeme{
		Key:      "you_are_banned",
		FallBack: "You are banned",
	}
)

// Authorization ...
func (m *Middleware) Authorization() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		user := users_utils.UserFromContext(ctx)
		if user.State == model.UserStateBanned {
			st := model.GetStatus(ctx, codes.PermissionDenied, "permission denied", "You are banned", nil, youAreBannedLexeme)

			return nil, st.Err()
		}

		userRoles, err := m.userRolesFacade.GetUserRolesByUserID(ctx, user.ID)
		if err != nil {
			st := status.New(codes.Internal, err.Error())
			return nil, st.Err()
		}

		if roles, ok := grpcRules[info.FullMethod]; ok {
			if _, ok := roles[Public]; ok {
				return handler(ctx, req)
			}

			for _, userRole := range userRoles {
				if _, ok := roles[userRole.Role.String()]; ok {
					return handler(ctx, req)
				}
			}
		} else {
			logger.ErrorKV(ctx, "roles for method not found", "method", info.FullMethod)
		}

		st := model.GetStatus(ctx, codes.PermissionDenied, "", "permission denied", nil, permissionDeniedLexeme)

		return nil, st.Err()
	}
}
