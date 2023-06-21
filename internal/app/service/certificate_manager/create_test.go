package certificatemanager

import (
	"context"
	"errors"
	"testing"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	pkgmodel "github.com/nikita5637/quiz-registrator-api/pkg/model"
	certificatemanagerpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/certificate_manager"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestRegistrator_CreateCertificate(t *testing.T) {
	t.Run("request validation error. invalid info JSON value", func(t *testing.T) {
		fx := tearUp(t)

		got, err := fx.certificateManager.CreateCertificate(fx.ctx, &certificatemanagerpb.CreateCertificateRequest{
			Certificate: &certificatemanagerpb.Certificate{
				Type: certificatemanagerpb.CertificateType_CERTIFICATE_TYPE_FREE_PASS,
				Info: "invalid json",
			},
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.InvalidArgument, st.Code())
		assert.Len(t, st.Details(), 2)
	})

	t.Run("request validation error. invalid certificate type", func(t *testing.T) {
		fx := tearUp(t)

		got, err := fx.certificateManager.CreateCertificate(fx.ctx, &certificatemanagerpb.CreateCertificateRequest{
			Certificate: &certificatemanagerpb.Certificate{
				Type: certificatemanagerpb.CertificateType(pkgmodel.NumberOfCertificateTypes),
				Info: "{}",
			},
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.InvalidArgument, st.Code())
		assert.Len(t, st.Details(), 2)
	})

	t.Run("error while create certificate. won_on game not found", func(t *testing.T) {
		fx := tearUp(t)

		fx.certificatesFacade.EXPECT().CreateCertificate(fx.ctx, model.Certificate{
			Type:    pkgmodel.CertificateTypeFreePass,
			WonOn:   1,
			SpentOn: model.NewMaybeInt32(2),
			Info:    model.NewMaybeString("{\"sum\": 5000, \"expired\": \"2023-06-16\"}"),
		}).Return(model.Certificate{}, model.ErrWonOnGameNotFound)

		got, err := fx.certificateManager.CreateCertificate(fx.ctx, &certificatemanagerpb.CreateCertificateRequest{
			Certificate: &certificatemanagerpb.Certificate{
				Type:    certificatemanagerpb.CertificateType_CERTIFICATE_TYPE_FREE_PASS,
				WonOn:   1,
				SpentOn: 2,
				Info:    "{\"sum\": 5000, \"expired\": \"2023-06-16\"}",
			},
		})

		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.InvalidArgument, st.Code())
		assert.Len(t, st.Details(), 2)
	})

	t.Run("error while create certificate. spent_on game not found", func(t *testing.T) {
		fx := tearUp(t)

		fx.certificatesFacade.EXPECT().CreateCertificate(fx.ctx, model.Certificate{
			Type:    pkgmodel.CertificateTypeFreePass,
			WonOn:   1,
			SpentOn: model.NewMaybeInt32(2),
			Info:    model.NewMaybeString("{\"sum\": 5000, \"expired\": \"2023-06-16\"}"),
		}).Return(model.Certificate{}, model.ErrSpentOnGameNotFound)

		got, err := fx.certificateManager.CreateCertificate(fx.ctx, &certificatemanagerpb.CreateCertificateRequest{
			Certificate: &certificatemanagerpb.Certificate{
				Type:    certificatemanagerpb.CertificateType_CERTIFICATE_TYPE_FREE_PASS,
				WonOn:   1,
				SpentOn: 2,
				Info:    "{\"sum\": 5000, \"expired\": \"2023-06-16\"}",
			},
		})

		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.InvalidArgument, st.Code())
		assert.Len(t, st.Details(), 2)
	})

	t.Run("error while create certificate. other error", func(t *testing.T) {
		fx := tearUp(t)

		fx.certificatesFacade.EXPECT().CreateCertificate(fx.ctx, model.Certificate{
			Type:    pkgmodel.CertificateTypeFreePass,
			WonOn:   1,
			SpentOn: model.NewMaybeInt32(2),
			Info:    model.NewMaybeString("{\"sum\": 5000, \"expired\": \"2023-06-16\"}"),
		}).Return(model.Certificate{}, errors.New("some error"))

		got, err := fx.certificateManager.CreateCertificate(fx.ctx, &certificatemanagerpb.CreateCertificateRequest{
			Certificate: &certificatemanagerpb.Certificate{
				Type:    certificatemanagerpb.CertificateType_CERTIFICATE_TYPE_FREE_PASS,
				WonOn:   1,
				SpentOn: 2,
				Info:    "{\"sum\": 5000, \"expired\": \"2023-06-16\"}",
			},
		})

		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.Internal, st.Code())
		assert.Len(t, st.Details(), 0)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		fx.certificatesFacade.EXPECT().CreateCertificate(fx.ctx, model.Certificate{
			Type:    pkgmodel.CertificateTypeFreePass,
			WonOn:   1,
			SpentOn: model.NewMaybeInt32(2),
			Info:    model.NewMaybeString("{\"sum\": 5000, \"expired\": \"2023-06-16\"}"),
		}).Return(model.Certificate{
			ID:      777,
			Type:    pkgmodel.CertificateTypeFreePass,
			WonOn:   1,
			SpentOn: model.NewMaybeInt32(2),
			Info:    model.NewMaybeString("{\"sum\": 5000, \"expired\": \"2023-06-16\"}"),
		}, nil)

		got, err := fx.certificateManager.CreateCertificate(fx.ctx, &certificatemanagerpb.CreateCertificateRequest{
			Certificate: &certificatemanagerpb.Certificate{
				Type:    certificatemanagerpb.CertificateType_CERTIFICATE_TYPE_FREE_PASS,
				WonOn:   1,
				SpentOn: 2,
				Info:    "{\"sum\": 5000, \"expired\": \"2023-06-16\"}",
			},
		})

		assert.Equal(t, &certificatemanagerpb.Certificate{
			Id:      777,
			Type:    certificatemanagerpb.CertificateType_CERTIFICATE_TYPE_FREE_PASS,
			WonOn:   1,
			SpentOn: 2,
			Info:    "{\"sum\": 5000, \"expired\": \"2023-06-16\"}",
		}, got)
		assert.NoError(t, err)
	})
}

func Test_validateCreateCertificateRequest(t *testing.T) {
	type args struct {
		ctx context.Context
		req *certificatemanagerpb.CreateCertificateRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "invalid JSON",
			args: args{
				ctx: context.Background(),
				req: &certificatemanagerpb.CreateCertificateRequest{
					Certificate: &certificatemanagerpb.Certificate{
						Info: "invalid JSON",
					},
				},
			},
			wantErr: true,
		},
		{
			name: "empty json string",
			args: args{
				ctx: context.Background(),
				req: &certificatemanagerpb.CreateCertificateRequest{
					Certificate: &certificatemanagerpb.Certificate{
						Info: "",
					},
				},
			},
			wantErr: true,
		},
		{
			name: "len of JSON string gt 256",
			args: args{
				ctx: context.Background(),
				req: &certificatemanagerpb.CreateCertificateRequest{
					Certificate: &certificatemanagerpb.Certificate{
						Type: certificatemanagerpb.CertificateType_CERTIFICATE_TYPE_FREE_PASS,
						Info: "{\"a\":\"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa\"}",
					},
				},
			},
			wantErr: true,
		},
		{
			name: "invalid certifcate type. eq 0",
			args: args{
				ctx: context.Background(),
				req: &certificatemanagerpb.CreateCertificateRequest{
					Certificate: &certificatemanagerpb.Certificate{
						Type: certificatemanagerpb.CertificateType_CERTIFICATE_TYPE_INVALID,
						Info: "{}",
					},
				},
			},
			wantErr: true,
		},
		{
			name: "invalid certifcate type. lt 1 and neq 0",
			args: args{
				ctx: context.Background(),
				req: &certificatemanagerpb.CreateCertificateRequest{
					Certificate: &certificatemanagerpb.Certificate{
						Type: -1,
						Info: "{}",
					},
				},
			},
			wantErr: true,
		},
		{
			name: "invalid certifcate type. gt number of certificate types",
			args: args{
				ctx: context.Background(),
				req: &certificatemanagerpb.CreateCertificateRequest{
					Certificate: &certificatemanagerpb.Certificate{
						Type: certificatemanagerpb.CertificateType(pkgmodel.NumberOfCertificateTypes),
						Info: "{}",
					},
				},
			},
			wantErr: true,
		},
		{
			name: "ok",
			args: args{
				ctx: context.Background(),
				req: &certificatemanagerpb.CreateCertificateRequest{
					Certificate: &certificatemanagerpb.Certificate{
						Type: certificatemanagerpb.CertificateType_CERTIFICATE_TYPE_FREE_PASS,
						Info: "{}",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateCreateCertificateRequest(tt.args.ctx, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("validateCreateCertificateRequest() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
