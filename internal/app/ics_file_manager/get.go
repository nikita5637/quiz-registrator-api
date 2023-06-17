package icsfilemanager

import (
	"context"
	"errors"
	"fmt"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	icsfilemanagerpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/ics_file_manager"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetICSFile ...
func (i *Implementation) GetICSFile(ctx context.Context, req *icsfilemanagerpb.GetICSFileRequest) (*icsfilemanagerpb.ICSFile, error) {
	icsFile, err := i.icsFilesFacade.GetICSFile(ctx, req.GetId())
	if err != nil {
		st := status.New(codes.Internal, err.Error())
		if errors.Is(err, model.ErrICSFileNotFound) {
			reason := fmt.Sprintf("ICS file with ID %d not found", req.GetId())
			st = model.GetStatus(ctx, codes.NotFound, err, reason, icsFileNotFoundLexeme)
		}

		return nil, st.Err()
	}

	return convertModelICSFileToProtoICSFile(icsFile), nil
}

// GetICSFileByGameID ...
func (i *Implementation) GetICSFileByGameID(ctx context.Context, req *icsfilemanagerpb.GetICSFileByGameIDRequest) (*icsfilemanagerpb.ICSFile, error) {
	icsFile, err := i.icsFilesFacade.GetICSFileByGameID(ctx, req.GetGameId())
	if err != nil {
		st := status.New(codes.Internal, err.Error())
		if errors.Is(err, model.ErrICSFileNotFound) {
			reason := fmt.Sprintf("ICS file with game ID %d not found", req.GetGameId())
			st = model.GetStatus(ctx, codes.NotFound, err, reason, icsFileNotFoundLexeme)
		}

		return nil, st.Err()
	}

	return convertModelICSFileToProtoICSFile(icsFile), nil
}
