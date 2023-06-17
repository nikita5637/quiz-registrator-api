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
)

var (
	icsFileForGameAlreadyExistsLexeme = i18n.Lexeme{
		Key:      "ICS_file_for_game_already_exists",
		FallBack: "ICS file for game already exists",
	}
	icsFileNameAlreadyExistsLexeme = i18n.Lexeme{
		Key:      "ICS_file_name_already_exists",
		FallBack: "ICS file name already exists",
	}
)

// CreateICSFile ...
func (i *Implementation) CreateICSFile(ctx context.Context, req *icsfilemanagerpb.CreateICSFileRequest) (*icsfilemanagerpb.ICSFile, error) {
	if err := validateCreateICSRequest(ctx, req); err != nil {
		st := status.New(codes.InvalidArgument, err.Error())
		return nil, st.Err()
	}

	icsFile, err := i.icsFilesFacade.CreateICSFile(ctx, model.ICSFile{
		GameID: req.GetIcsFile().GetGameId(),
		Name:   req.GetIcsFile().GetName(),
	})
	if err != nil {
		st := status.New(codes.Internal, err.Error())
		if errors.Is(err, model.ErrGameNotFound) {
			reason := fmt.Sprintf("game with id %d not found", req.GetIcsFile().GetGameId())
			st = model.GetStatus(ctx, codes.InvalidArgument, err, reason, i18n.GameNotFoundLexeme)
		} else if errors.Is(err, model.ErrICSFileForGameAlreadyExists) {
			reason := fmt.Sprintf("ICS file for game with id %d already exists", req.GetIcsFile().GetGameId())
			st = model.GetStatus(ctx, codes.AlreadyExists, err, reason, icsFileForGameAlreadyExistsLexeme)
		} else if errors.Is(err, model.ErrICSFileNameAlreadyExists) {
			reason := fmt.Sprintf("ICS file name %s already exists", req.GetIcsFile().GetName())
			st = model.GetStatus(ctx, codes.AlreadyExists, err, reason, icsFileNameAlreadyExistsLexeme)
		}

		return nil, st.Err()
	}

	return convertModelICSFileToProtoICSFile(icsFile), nil
}

func validateCreateICSRequest(ctx context.Context, req *icsfilemanagerpb.CreateICSFileRequest) error {
	return nil
}
