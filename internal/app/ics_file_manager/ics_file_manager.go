//go:generate mockery --case underscore --name ICSFilesFacade --with-expecter

package icsfilemanager

import (
	"context"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	icsfilemanagerpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/ics_file_manager"
)

// ICSFilesFacade ...
type ICSFilesFacade interface {
	CreateICSFile(ctx context.Context, ics model.ICSFile) (model.ICSFile, error)
	DeleteICSFile(ctx context.Context, id int32) error
	GetICSFile(ctx context.Context, id int32) (model.ICSFile, error)
	GetICSFileByGameID(ctx context.Context, gameID int32) (model.ICSFile, error)
	ListICSFiles(ctx context.Context) ([]model.ICSFile, error)
}

// Implementation ...
type Implementation struct {
	icsFilesFacade ICSFilesFacade
	icsfilemanagerpb.UnimplementedServiceServer
}

// Config ...
type Config struct {
	ICSFilesFacade ICSFilesFacade
}

// New ...
func New(cfg Config) *Implementation {
	return &Implementation{
		icsFilesFacade: cfg.ICSFilesFacade,
	}

}
