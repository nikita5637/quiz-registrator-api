package certificatemanager

import (
	"errors"
	"testing"

	"github.com/mono83/maybe"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/certificates"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	certificatemanagerpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/certificate_manager"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func TestRegistrator_CreateCertificate(t *testing.T) {
	t.Run("validation error. invalid certificate type", func(t *testing.T) {
		fx := tearUp(t)

		got, err := fx.certificateManager.CreateCertificate(fx.ctx, &certificatemanagerpb.CreateCertificateRequest{
			Certificate: &certificatemanagerpb.Certificate{
				Type: certificatemanagerpb.CertificateType_CERTIFICATE_TYPE_INVALID,
			},
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.InvalidArgument, st.Code())
		assert.Len(t, st.Details(), 2)
	})

	t.Run("validation error. invalid certificate type", func(t *testing.T) {
		fx := tearUp(t)

		got, err := fx.certificateManager.CreateCertificate(fx.ctx, &certificatemanagerpb.CreateCertificateRequest{
			Certificate: &certificatemanagerpb.Certificate{
				Type: certificatemanagerpb.CertificateType_CERTIFICATE_TYPE_BAR_BILL_PAYMENT + 1,
			},
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.InvalidArgument, st.Code())
		assert.Len(t, st.Details(), 2)
	})

	t.Run("validation error. invalid won on", func(t *testing.T) {
		fx := tearUp(t)

		got, err := fx.certificateManager.CreateCertificate(fx.ctx, &certificatemanagerpb.CreateCertificateRequest{
			Certificate: &certificatemanagerpb.Certificate{
				Type:  certificatemanagerpb.CertificateType_CERTIFICATE_TYPE_BAR_BILL_PAYMENT,
				WonOn: 0,
			},
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.InvalidArgument, st.Code())
		assert.Len(t, st.Details(), 2)
	})

	t.Run("validation error. invalid won on", func(t *testing.T) {
		fx := tearUp(t)

		got, err := fx.certificateManager.CreateCertificate(fx.ctx, &certificatemanagerpb.CreateCertificateRequest{
			Certificate: &certificatemanagerpb.Certificate{
				Type:  certificatemanagerpb.CertificateType_CERTIFICATE_TYPE_BAR_BILL_PAYMENT,
				WonOn: -1,
			},
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.InvalidArgument, st.Code())
		assert.Len(t, st.Details(), 2)
	})

	t.Run("validation error. invalid spent on", func(t *testing.T) {
		fx := tearUp(t)

		got, err := fx.certificateManager.CreateCertificate(fx.ctx, &certificatemanagerpb.CreateCertificateRequest{
			Certificate: &certificatemanagerpb.Certificate{
				Type:    certificatemanagerpb.CertificateType_CERTIFICATE_TYPE_BAR_BILL_PAYMENT,
				WonOn:   1,
				SpentOn: &wrapperspb.Int32Value{},
			},
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.InvalidArgument, st.Code())
		assert.Len(t, st.Details(), 2)
	})

	t.Run("validation error. invalid spent on", func(t *testing.T) {
		fx := tearUp(t)

		got, err := fx.certificateManager.CreateCertificate(fx.ctx, &certificatemanagerpb.CreateCertificateRequest{
			Certificate: &certificatemanagerpb.Certificate{
				Type:  certificatemanagerpb.CertificateType_CERTIFICATE_TYPE_BAR_BILL_PAYMENT,
				WonOn: 1,
				SpentOn: &wrapperspb.Int32Value{
					Value: -1,
				},
			},
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.InvalidArgument, st.Code())
		assert.Len(t, st.Details(), 2)
	})

	t.Run("request validation error. invalid info", func(t *testing.T) {
		fx := tearUp(t)

		got, err := fx.certificateManager.CreateCertificate(fx.ctx, &certificatemanagerpb.CreateCertificateRequest{
			Certificate: &certificatemanagerpb.Certificate{
				Type:  certificatemanagerpb.CertificateType_CERTIFICATE_TYPE_FREE_PASS,
				WonOn: 1,
				Info:  &wrapperspb.StringValue{},
			},
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.InvalidArgument, st.Code())
		assert.Len(t, st.Details(), 2)
	})

	t.Run("request validation error. invalid info", func(t *testing.T) {
		fx := tearUp(t)

		got, err := fx.certificateManager.CreateCertificate(fx.ctx, &certificatemanagerpb.CreateCertificateRequest{
			Certificate: &certificatemanagerpb.Certificate{
				Type:  certificatemanagerpb.CertificateType_CERTIFICATE_TYPE_FREE_PASS,
				WonOn: 1,
				Info: &wrapperspb.StringValue{
					Value: "invalid JSON",
				},
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
			Type:    model.CertificateTypeFreePass,
			WonOn:   1,
			SpentOn: maybe.Just(int32(2)),
			Info:    maybe.Just("{\"sum\": 5000, \"expired\": \"2023-06-16\"}"),
		}).Return(model.Certificate{}, certificates.ErrWonOnGameNotFound)

		got, err := fx.certificateManager.CreateCertificate(fx.ctx, &certificatemanagerpb.CreateCertificateRequest{
			Certificate: &certificatemanagerpb.Certificate{
				Type:  certificatemanagerpb.CertificateType_CERTIFICATE_TYPE_FREE_PASS,
				WonOn: 1,
				SpentOn: &wrapperspb.Int32Value{
					Value: 2,
				},
				Info: &wrapperspb.StringValue{
					Value: "{\"sum\": 5000, \"expired\": \"2023-06-16\"}",
				},
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
			Type:    model.CertificateTypeFreePass,
			WonOn:   1,
			SpentOn: maybe.Just(int32(2)),
			Info:    maybe.Just("{\"sum\": 5000, \"expired\": \"2023-06-16\"}"),
		}).Return(model.Certificate{}, certificates.ErrSpentOnGameNotFound)

		got, err := fx.certificateManager.CreateCertificate(fx.ctx, &certificatemanagerpb.CreateCertificateRequest{
			Certificate: &certificatemanagerpb.Certificate{
				Type:  certificatemanagerpb.CertificateType_CERTIFICATE_TYPE_FREE_PASS,
				WonOn: 1,
				SpentOn: &wrapperspb.Int32Value{
					Value: 2,
				},
				Info: &wrapperspb.StringValue{
					Value: "{\"sum\": 5000, \"expired\": \"2023-06-16\"}",
				},
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
			Type:    model.CertificateTypeFreePass,
			WonOn:   1,
			SpentOn: maybe.Just(int32(2)),
			Info:    maybe.Just("{\"sum\": 5000, \"expired\": \"2023-06-16\"}"),
		}).Return(model.Certificate{}, errors.New("some error"))

		got, err := fx.certificateManager.CreateCertificate(fx.ctx, &certificatemanagerpb.CreateCertificateRequest{
			Certificate: &certificatemanagerpb.Certificate{
				Type:  certificatemanagerpb.CertificateType_CERTIFICATE_TYPE_FREE_PASS,
				WonOn: 1,
				SpentOn: &wrapperspb.Int32Value{
					Value: 2,
				},
				Info: &wrapperspb.StringValue{
					Value: "{\"sum\": 5000, \"expired\": \"2023-06-16\"}",
				},
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
			Type:    model.CertificateTypeFreePass,
			WonOn:   1,
			SpentOn: maybe.Just(int32(2)),
			Info:    maybe.Just("{\"sum\": 5000, \"expired\": \"2023-06-16\"}"),
		}).Return(model.Certificate{
			ID:      777,
			Type:    model.CertificateTypeFreePass,
			WonOn:   1,
			SpentOn: maybe.Just(int32(2)),
			Info:    maybe.Just("{\"sum\": 5000, \"expired\": \"2023-06-16\"}"),
		}, nil)

		got, err := fx.certificateManager.CreateCertificate(fx.ctx, &certificatemanagerpb.CreateCertificateRequest{
			Certificate: &certificatemanagerpb.Certificate{
				Type:  certificatemanagerpb.CertificateType_CERTIFICATE_TYPE_FREE_PASS,
				WonOn: 1,
				SpentOn: &wrapperspb.Int32Value{
					Value: 2,
				},
				Info: &wrapperspb.StringValue{
					Value: "{\"sum\": 5000, \"expired\": \"2023-06-16\"}",
				},
			},
		})

		assert.Equal(t, &certificatemanagerpb.Certificate{
			Id:    777,
			Type:  certificatemanagerpb.CertificateType_CERTIFICATE_TYPE_FREE_PASS,
			WonOn: 1,
			SpentOn: &wrapperspb.Int32Value{
				Value: 2,
			},
			Info: &wrapperspb.StringValue{
				Value: "{\"sum\": 5000, \"expired\": \"2023-06-16\"}",
			},
		}, got)
		assert.NoError(t, err)
	})
}

func Test_validateCreatedCertificate(t *testing.T) {
	type args struct {
		certificate model.Certificate
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "no ID",
			args: args{
				certificate: model.Certificate{
					Type:    model.CertificateTypeFreePass,
					WonOn:   1,
					SpentOn: maybe.Nothing[int32](),
					Info:    maybe.Nothing[string](),
				},
			},
			wantErr: false,
		},
		{
			name: "no type",
			args: args{
				certificate: model.Certificate{
					ID:      1,
					WonOn:   1,
					SpentOn: maybe.Nothing[int32](),
					Info:    maybe.Nothing[string](),
				},
			},
			wantErr: true,
		},
		{
			name: "no won_on",
			args: args{
				certificate: model.Certificate{
					ID:      1,
					Type:    model.CertificateTypeFreePass,
					SpentOn: maybe.Nothing[int32](),
					Info:    maybe.Nothing[string](),
				},
			},
			wantErr: true,
		},
		{
			name: "won_on eq minWonOn and valid",
			args: args{
				certificate: model.Certificate{
					ID:      1,
					Type:    model.CertificateTypeFreePass,
					WonOn:   minWonOn,
					SpentOn: maybe.Nothing[int32](),
					Info:    maybe.Nothing[string](),
				},
			},
			wantErr: false,
		},
		{
			name: "no spent_on",
			args: args{
				certificate: model.Certificate{
					ID:      1,
					Type:    model.CertificateTypeFreePass,
					WonOn:   1,
					SpentOn: maybe.Nothing[int32](),
					Info:    maybe.Nothing[string](),
				},
			},
			wantErr: false,
		},
		{
			name: "spent_on eq 0 and valid",
			args: args{
				certificate: model.Certificate{
					ID:      1,
					Type:    model.CertificateTypeFreePass,
					WonOn:   1,
					SpentOn: maybe.Just(int32(0)),
					Info:    maybe.Nothing[string](),
				},
			},
			wantErr: true,
		},
		{
			name: "spent_on eq minSpentOn and valid",
			args: args{
				certificate: model.Certificate{
					ID:      1,
					Type:    model.CertificateTypeFreePass,
					WonOn:   1,
					SpentOn: maybe.Just(minSpentOn),
					Info:    maybe.Nothing[string](),
				},
			},
			wantErr: false,
		},
		{
			name: "spent_on eq 0 and not valid",
			args: args{
				certificate: model.Certificate{
					ID:      1,
					Type:    model.CertificateTypeFreePass,
					WonOn:   1,
					SpentOn: maybe.Nothing[int32](),
					Info:    maybe.Nothing[string](),
				},
			},
			wantErr: false,
		},
		{
			name: "info is empty and valid",
			args: args{
				certificate: model.Certificate{
					ID:      1,
					Type:    model.CertificateTypeFreePass,
					WonOn:   1,
					SpentOn: maybe.Just(minSpentOn),
					Info:    maybe.Just(""),
				},
			},
			wantErr: true,
		},
		{
			name: "info is empty and not valid",
			args: args{
				certificate: model.Certificate{
					ID:      1,
					Type:    model.CertificateTypeFreePass,
					WonOn:   1,
					SpentOn: maybe.Just(minSpentOn),
					Info:    maybe.Nothing[string](),
				},
			},
			wantErr: false,
		},
		{
			name: "ok",
			args: args{
				certificate: model.Certificate{
					ID:      1,
					Type:    model.CertificateTypeFreePass,
					WonOn:   1,
					SpentOn: maybe.Just(minSpentOn),
					Info:    maybe.Just("{}"),
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateCreatedCertificate(tt.args.certificate); (err != nil) != tt.wantErr {
				t.Errorf("validateCreatedCertificate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
