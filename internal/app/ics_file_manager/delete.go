package icsfilemanager

import (
	"context"
	"errors"
	"fmt"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/i18n"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	icsfilemanagerpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/ics_file_manager"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

var (
	icsFileNotFoundLexeme = i18n.Lexeme{
		Key:      "ICS_file_not_found",
		FallBack: "ICS file not found",
	}
)

// DeleteICSFile ...
func (i *Implementation) DeleteICSFile(ctx context.Context, req *icsfilemanagerpb.DeleteICSFileRequest) (*emptypb.Empty, error) {
	err := i.icsFilesFacade.DeleteICSFile(ctx, req.GetId())
	if err != nil {
		st := status.New(codes.Internal, err.Error())
		if errors.Is(err, model.ErrICSFileNotFound) {
			reason := fmt.Sprintf("ICS file with ID %d not found", req.GetId())
			st = model.GetStatus(ctx, codes.NotFound, err, reason, icsFileNotFoundLexeme)
		}

		return nil, st.Err()
	}

	return &emptypb.Empty{}, nil
}
