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
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func TestRegistrator_PatchCertificate(t *testing.T) {
	t.Run("error. original certificate not found", func(t *testing.T) {
		fx := tearUp(t)

		fx.certificatesFacade.EXPECT().GetCertificate(fx.ctx, int32(1)).Return(model.Certificate{}, certificates.ErrCertificateNotFound)

		got, err := fx.certificateManager.PatchCertificate(fx.ctx, &certificatemanagerpb.PatchCertificateRequest{
			Certificate: &certificatemanagerpb.Certificate{
				Id:      1,
				Type:    certificatemanagerpb.CertificateType_CERTIFICATE_TYPE_BAR_BILL_PAYMENT,
				WonOn:   10,
				SpentOn: wrapperspb.Int32(190),
				Info:    wrapperspb.String("invalid JSON"),
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

	t.Run("internal error while get original certificate", func(t *testing.T) {
		fx := tearUp(t)

		fx.certificatesFacade.EXPECT().GetCertificate(fx.ctx, int32(1)).Return(model.Certificate{}, errors.New("some error"))

		got, err := fx.certificateManager.PatchCertificate(fx.ctx, &certificatemanagerpb.PatchCertificateRequest{
			Certificate: &certificatemanagerpb.Certificate{
				Id:      1,
				Type:    certificatemanagerpb.CertificateType_CERTIFICATE_TYPE_BAR_BILL_PAYMENT,
				WonOn:   10,
				SpentOn: wrapperspb.Int32(190),
				Info:    wrapperspb.String("invalid JSON"),
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
		assert.Equal(t, codes.Internal, st.Code())
		assert.Len(t, st.Details(), 0)
	})

	t.Run("validate error. invalid certificate type", func(t *testing.T) {
		fx := tearUp(t)

		fx.certificatesFacade.EXPECT().GetCertificate(fx.ctx, int32(1)).Return(model.Certificate{
			ID:      1,
			Type:    model.CertificateTypeFreePass,
			WonOn:   10,
			SpentOn: maybe.Just(int32(1)),
			Info:    maybe.Just("{}"),
		}, nil)

		got, err := fx.certificateManager.PatchCertificate(fx.ctx, &certificatemanagerpb.PatchCertificateRequest{
			Certificate: &certificatemanagerpb.Certificate{
				Id:      1,
				Type:    certificatemanagerpb.CertificateType_CERTIFICATE_TYPE_INVALID,
				WonOn:   10,
				SpentOn: wrapperspb.Int32(190),
				Info:    wrapperspb.String("{}"),
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

	t.Run("error ErrWonOnGameNotFound", func(t *testing.T) {
		fx := tearUp(t)

		fx.certificatesFacade.EXPECT().GetCertificate(fx.ctx, int32(1)).Return(model.Certificate{
			ID:      1,
			Type:    model.CertificateTypeFreePass,
			WonOn:   10,
			SpentOn: maybe.Just(int32(1)),
			Info:    maybe.Just("{}"),
		}, nil)

		fx.certificatesFacade.EXPECT().PatchCertificate(fx.ctx, model.Certificate{
			ID:      1,
			Type:    model.CertificateTypeBarBillPayment,
			WonOn:   10,
			SpentOn: maybe.Just(int32(190)),
			Info:    maybe.Just("{}"),
		}).Return(model.Certificate{}, certificates.ErrWonOnGameNotFound)

		got, err := fx.certificateManager.PatchCertificate(fx.ctx, &certificatemanagerpb.PatchCertificateRequest{
			Certificate: &certificatemanagerpb.Certificate{
				Id:      1,
				Type:    certificatemanagerpb.CertificateType_CERTIFICATE_TYPE_BAR_BILL_PAYMENT,
				WonOn:   10,
				SpentOn: wrapperspb.Int32(190),
				Info:    wrapperspb.String("{}"),
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

		fx.certificatesFacade.EXPECT().GetCertificate(fx.ctx, int32(1)).Return(model.Certificate{
			ID:      1,
			Type:    model.CertificateTypeFreePass,
			WonOn:   10,
			SpentOn: maybe.Just(int32(1)),
			Info:    maybe.Just("{}"),
		}, nil)

		fx.certificatesFacade.EXPECT().PatchCertificate(fx.ctx, model.Certificate{
			ID:      1,
			Type:    model.CertificateTypeBarBillPayment,
			WonOn:   10,
			SpentOn: maybe.Just(int32(190)),
			Info:    maybe.Just("{}"),
		}).Return(model.Certificate{}, certificates.ErrSpentOnGameNotFound)

		got, err := fx.certificateManager.PatchCertificate(fx.ctx, &certificatemanagerpb.PatchCertificateRequest{
			Certificate: &certificatemanagerpb.Certificate{
				Id:      1,
				Type:    certificatemanagerpb.CertificateType_CERTIFICATE_TYPE_BAR_BILL_PAYMENT,
				WonOn:   10,
				SpentOn: wrapperspb.Int32(190),
				Info:    wrapperspb.String("{}"),
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

		fx.certificatesFacade.EXPECT().GetCertificate(fx.ctx, int32(1)).Return(model.Certificate{
			ID:      1,
			Type:    model.CertificateTypeFreePass,
			WonOn:   1,
			SpentOn: maybe.Just(int32(1)),
			Info:    maybe.Just("{}"),
		}, nil)

		fx.certificatesFacade.EXPECT().PatchCertificate(fx.ctx, model.Certificate{
			ID:      1,
			Type:    model.CertificateTypeBarBillPayment,
			WonOn:   10,
			SpentOn: maybe.Just(int32(190)),
			Info:    maybe.Just("{\"a\":1}"),
		}).Return(model.Certificate{
			ID:      1,
			Type:    model.CertificateTypeBarBillPayment,
			WonOn:   100,
			SpentOn: maybe.Just(int32(190)),
			Info:    maybe.Just("{\"a\":1}"),
		}, nil)

		got, err := fx.certificateManager.PatchCertificate(fx.ctx, &certificatemanagerpb.PatchCertificateRequest{
			Certificate: &certificatemanagerpb.Certificate{
				Id:      1,
				Type:    certificatemanagerpb.CertificateType_CERTIFICATE_TYPE_BAR_BILL_PAYMENT,
				WonOn:   10,
				SpentOn: wrapperspb.Int32(190),
				Info:    wrapperspb.String("{\"a\":1}"),
			},
			UpdateMask: &fieldmaskpb.FieldMask{
				Paths: []string{
					"type",
					"won_on",
					"spent_on",
					"info",
				},
			},
		})

		assert.Equal(t, &certificatemanagerpb.Certificate{
			Id:      1,
			Type:    certificatemanagerpb.CertificateType_CERTIFICATE_TYPE_BAR_BILL_PAYMENT,
			WonOn:   100,
			SpentOn: wrapperspb.Int32(190),
			Info:    wrapperspb.String("{\"a\":1}"),
		}, got)
		assert.NoError(t, err)
	})
}

func Test_validatePatchedCertificate(t *testing.T) {
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
			wantErr: true,
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
			if err := validatePatchedCertificate(tt.args.certificate); (err != nil) != tt.wantErr {
				t.Errorf("validatePatchedCertificate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
