package certificatemanager

import (
	"context"
	"testing"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	certificatemanagerpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/certificate_manager"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

func TestRegistrator_PatchCertificate(t *testing.T) {
	t.Run("validate error. invalid info json value", func(t *testing.T) {
		fx := tearUp(t)

		got, err := fx.certificateManager.PatchCertificate(fx.ctx, &certificatemanagerpb.PatchCertificateRequest{
			Certificate: &certificatemanagerpb.Certificate{
				Id:      1,
				Type:    certificatemanagerpb.CertificateType_CERTIFICATE_TYPE_BAR_BILL_PAYMENT,
				WonOn:   10,
				SpentOn: 190,
				Info:    "invalid JSON",
			},
			UpdateMask: &fieldmaskpb.FieldMask{
				Paths: []string{
					"type",
					"spent_on",
				},
			},
		})

		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.InvalidArgument, st.Code())
		assert.Len(t, st.Details(), 2)
	})

	t.Run("validate error. invalid certificate type", func(t *testing.T) {
		fx := tearUp(t)

		got, err := fx.certificateManager.PatchCertificate(fx.ctx, &certificatemanagerpb.PatchCertificateRequest{
			Certificate: &certificatemanagerpb.Certificate{
				Id:      1,
				Type:    certificatemanagerpb.CertificateType_CERTIFICATE_TYPE_INVALID,
				WonOn:   10,
				SpentOn: 190,
				Info:    "{}",
			},
			UpdateMask: &fieldmaskpb.FieldMask{
				Paths: []string{
					"type",
					"spent_on",
				},
			},
		})

		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.InvalidArgument, st.Code())
		assert.Len(t, st.Details(), 2)
	})

	t.Run("error ErrCertificateNotFound", func(t *testing.T) {
		fx := tearUp(t)

		fx.certificatesFacade.EXPECT().PatchCertificate(fx.ctx, model.Certificate{
			ID:      1,
			Type:    model.CertificateTypeBarBillPayment,
			WonOn:   10,
			SpentOn: model.NewMaybeInt32(190),
			Info:    model.NewMaybeString("{}"),
		}, []string{"type", "spent_on"}).Return(model.Certificate{}, model.ErrCertificateNotFound)

		got, err := fx.certificateManager.PatchCertificate(fx.ctx, &certificatemanagerpb.PatchCertificateRequest{
			Certificate: &certificatemanagerpb.Certificate{
				Id:      1,
				Type:    certificatemanagerpb.CertificateType_CERTIFICATE_TYPE_BAR_BILL_PAYMENT,
				WonOn:   10,
				SpentOn: 190,
				Info:    "{}",
			},
			UpdateMask: &fieldmaskpb.FieldMask{
				Paths: []string{
					"type",
					"spent_on",
				},
			},
		})

		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.NotFound, st.Code())
		assert.Len(t, st.Details(), 2)
	})

	t.Run("error ErrWonOnGameNotFound", func(t *testing.T) {
		fx := tearUp(t)

		fx.certificatesFacade.EXPECT().PatchCertificate(fx.ctx, model.Certificate{
			ID:      1,
			Type:    model.CertificateTypeBarBillPayment,
			WonOn:   10,
			SpentOn: model.NewMaybeInt32(190),
			Info:    model.NewMaybeString("{}"),
		}, []string{"type", "spent_on"}).Return(model.Certificate{}, model.ErrWonOnGameNotFound)

		got, err := fx.certificateManager.PatchCertificate(fx.ctx, &certificatemanagerpb.PatchCertificateRequest{
			Certificate: &certificatemanagerpb.Certificate{
				Id:      1,
				Type:    certificatemanagerpb.CertificateType_CERTIFICATE_TYPE_BAR_BILL_PAYMENT,
				WonOn:   10,
				SpentOn: 190,
				Info:    "{}",
			},
			UpdateMask: &fieldmaskpb.FieldMask{
				Paths: []string{
					"type",
					"spent_on",
				},
			},
		})

		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.InvalidArgument, st.Code())
		assert.Len(t, st.Details(), 2)
	})

	t.Run("error ErrSpentOnGameNotFound", func(t *testing.T) {
		fx := tearUp(t)

		fx.certificatesFacade.EXPECT().PatchCertificate(fx.ctx, model.Certificate{
			ID:      1,
			Type:    model.CertificateTypeBarBillPayment,
			WonOn:   10,
			SpentOn: model.NewMaybeInt32(190),
			Info:    model.NewMaybeString("{}"),
		}, []string{"type", "spent_on"}).Return(model.Certificate{}, model.ErrSpentOnGameNotFound)

		got, err := fx.certificateManager.PatchCertificate(fx.ctx, &certificatemanagerpb.PatchCertificateRequest{
			Certificate: &certificatemanagerpb.Certificate{
				Id:      1,
				Type:    certificatemanagerpb.CertificateType_CERTIFICATE_TYPE_BAR_BILL_PAYMENT,
				WonOn:   10,
				SpentOn: 190,
				Info:    "{}",
			},
			UpdateMask: &fieldmaskpb.FieldMask{
				Paths: []string{
					"type",
					"spent_on",
				},
			},
		})

		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.InvalidArgument, st.Code())
		assert.Len(t, st.Details(), 2)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		fx.certificatesFacade.EXPECT().PatchCertificate(fx.ctx, model.Certificate{
			ID:      1,
			Type:    model.CertificateTypeBarBillPayment,
			WonOn:   10,
			SpentOn: model.NewMaybeInt32(190),
			Info:    model.NewMaybeString("{}"),
		}, []string{"type", "spent_on"}).Return(model.Certificate{
			ID:      1,
			Type:    model.CertificateTypeBarBillPayment,
			WonOn:   100,
			SpentOn: model.NewMaybeInt32(190),
			Info:    model.NewMaybeString("some valid json"),
		}, nil)

		got, err := fx.certificateManager.PatchCertificate(fx.ctx, &certificatemanagerpb.PatchCertificateRequest{
			Certificate: &certificatemanagerpb.Certificate{
				Id:      1,
				Type:    certificatemanagerpb.CertificateType_CERTIFICATE_TYPE_BAR_BILL_PAYMENT,
				WonOn:   10,
				SpentOn: 190,
				Info:    "{}",
			},
			UpdateMask: &fieldmaskpb.FieldMask{
				Paths: []string{
					"type",
					"spent_on",
				},
			},
		})

		assert.Equal(t, &certificatemanagerpb.Certificate{
			Id:      1,
			Type:    certificatemanagerpb.CertificateType_CERTIFICATE_TYPE_BAR_BILL_PAYMENT,
			WonOn:   100,
			SpentOn: 190,
			Info:    "some valid json",
		}, got)
		assert.NoError(t, err)
	})
}

func Test_validatePatchCertificateRequest(t *testing.T) {
	type args struct {
		ctx context.Context
		req *certificatemanagerpb.PatchCertificateRequest
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
				req: &certificatemanagerpb.PatchCertificateRequest{
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
				req: &certificatemanagerpb.PatchCertificateRequest{
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
				req: &certificatemanagerpb.PatchCertificateRequest{
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
				req: &certificatemanagerpb.PatchCertificateRequest{
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
				req: &certificatemanagerpb.PatchCertificateRequest{
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
				req: &certificatemanagerpb.PatchCertificateRequest{
					Certificate: &certificatemanagerpb.Certificate{
						Type: certificatemanagerpb.CertificateType(100),
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
				req: &certificatemanagerpb.PatchCertificateRequest{
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
			if err := validatePatchCertificateRequest(tt.args.ctx, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("validatePatchCertificateRequest() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
