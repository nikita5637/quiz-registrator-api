package admin

import (
	"context"

	adminpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/admin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetUserRolesByUserID ...
func (i *Implementation) GetUserRolesByUserID(ctx context.Context, req *adminpb.GetUserRolesByUserIDRequest) (*adminpb.GetUserRolesByUserIDResponse, error) {
	userRoles, err := i.userRolesFacade.GetUserRolesByUserID(ctx, req.GetUserId())
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

	return &adminpb.GetUserRolesByUserIDResponse{
		UserRoles: respUserRoles,
	}, nil
}
