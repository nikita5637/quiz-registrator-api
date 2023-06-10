package adminservice

import (
	"context"
	"errors"
	"fmt"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/i18n"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	adminpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/admin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

var (
	userRoleNotFoundLexeme = i18n.Lexeme{
		Key:      "user_role_not_found",
		FallBack: "User role not found",
	}
)

// DeleteUserRole ...
func (i *Implementation) DeleteUserRole(ctx context.Context, req *adminpb.DeleteUserRoleRequest) (*emptypb.Empty, error) {
	err := i.userRolesFacade.DeleteUserRole(ctx, req.GetId())
	if err != nil {
		st := status.New(codes.Internal, err.Error())
		if errors.Is(err, model.ErrUserRoleNotFound) {
			reason := fmt.Sprintf("user role with ID %d not found", req.GetId())
			st = model.GetStatus(ctx, codes.NotFound, err, reason, userRoleNotFoundLexeme)
		}

		return nil, st.Err()
	}

	return &emptypb.Empty{}, nil
}
