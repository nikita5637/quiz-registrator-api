package usermanager

import (
	"context"
	"errors"

	users "github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/users"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	usermanagerpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/user_manager"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetUser ...
func (i *Implementation) GetUser(ctx context.Context, req *usermanagerpb.GetUserRequest) (*usermanagerpb.User, error) {
	user, err := i.usersFacade.GetUser(ctx, req.GetId())
	if err != nil {
		st := status.New(codes.Internal, err.Error())
		if errors.Is(err, users.ErrUserNotFound) {
			st = model.GetStatus(ctx, codes.NotFound, err, reasonUserNotFound, users.UserNotFoundLexeme)
		}

		return nil, st.Err()
	}

	return convertModelUserToProtoUser(user), nil
}

// GetUserByTelegramID ...
func (i *Implementation) GetUserByTelegramID(ctx context.Context, req *usermanagerpb.GetUserByTelegramIDRequest) (*usermanagerpb.User, error) {
	user, err := i.usersFacade.GetUserByTelegramID(ctx, req.GetTelegramId())
	if err != nil {
		st := status.New(codes.Internal, err.Error())
		if errors.Is(err, users.ErrUserNotFound) {
			st = model.GetStatus(ctx, codes.NotFound, err, reasonUserNotFound, users.UserNotFoundLexeme)
		}

		return nil, st.Err()
	}

	return convertModelUserToProtoUser(user), nil
}
