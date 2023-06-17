package icsfilemanager

import (
	"context"

	icsfilemanagerpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/ics_file_manager"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

// ListICSFiles ...
func (i *Implementation) ListICSFiles(ctx context.Context, _ *emptypb.Empty) (*icsfilemanagerpb.ListICSFilesResponse, error) {
	icsFiles, err := i.icsFilesFacade.ListICSFiles(ctx)
	if err != nil {
		st := status.New(codes.Internal, err.Error())
		return nil, st.Err()
	}

	respICSFiles := make([]*icsfilemanagerpb.ICSFile, 0, len(icsFiles))
	for _, icsFile := range icsFiles {
		respICSFiles = append(respICSFiles, &icsfilemanagerpb.ICSFile{
			Id:     icsFile.ID,
			GameId: icsFile.GameID,
			Name:   icsFile.Name,
		})
	}

	return &icsfilemanagerpb.ListICSFilesResponse{
		IcsFiles: respICSFiles,
	}, nil
}
