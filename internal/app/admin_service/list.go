package adminservice

import (
	"context"

	adminpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/admin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

// ListUserRoles ...
func (i *Implementation) ListUserRoles(ctx context.Context, _ *emptypb.Empty) (*adminpb.ListUserRolesResponse, error) {
	userRoles, err := i.userRolesFacade.ListUserRoles(ctx)
	if err != nil {
		st := status.New(codes.Internal, err.Error())
		return nil, st.Err()
	}

	respUserRoles := make([]*adminpb.UserRole, 0, len(userRoles))
	for _, userRole := range userRoles {
		respUserRoles = append(respUserRoles, &adminpb.UserRole{
			Id:     userRole.ID,
			UserId: userRole.UserID,
			Role:   adminpb.Role(userRole.Role),
		})
	}

	return &adminpb.ListUserRolesResponse{
		UserRoles: respUserRoles,
	}, nil
}
