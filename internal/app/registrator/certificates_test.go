package registrator

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	pkgmodel "github.com/nikita5637/quiz-registrator-api/pkg/model"
	"github.com/nikita5637/quiz-registrator-api/pkg/pb/registrator"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

func TestRegistrator_CreateCertificate(t *testing.T) {
	t.Run("request validation error. invalid info JSON value", func(t *testing.T) {
		fx := tearUp(t)

		got, err := fx.registrator.CreateCertificate(fx.ctx, &registrator.CreateCertificateRequest{
			Certificate: &registrator.Certificate{
				Type: registrator.CertificateType_CERTIFICATE_TYPE_FREE_PASS,
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

		got, err := fx.registrator.CreateCertificate(fx.ctx, &registrator.CreateCertificateRequest{
			Certificate: &registrator.Certificate{
				Type: registrator.CertificateType(pkgmodel.NumberOfCertificateTypes),
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

		got, err := fx.registrator.CreateCertificate(fx.ctx, &registrator.CreateCertificateRequest{
			Certificate: &registrator.Certificate{
				Type:    registrator.CertificateType_CERTIFICATE_TYPE_FREE_PASS,
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

		got, err := fx.registrator.CreateCertificate(fx.ctx, &registrator.CreateCertificateRequest{
			Certificate: &registrator.Certificate{
				Type:    registrator.CertificateType_CERTIFICATE_TYPE_FREE_PASS,
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

		got, err := fx.registrator.CreateCertificate(fx.ctx, &registrator.CreateCertificateRequest{
			Certificate: &registrator.Certificate{
				Type:    registrator.CertificateType_CERTIFICATE_TYPE_FREE_PASS,
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

		got, err := fx.registrator.CreateCertificate(fx.ctx, &registrator.CreateCertificateRequest{
			Certificate: &registrator.Certificate{
				Type:    registrator.CertificateType_CERTIFICATE_TYPE_FREE_PASS,
				WonOn:   1,
				SpentOn: 2,
				Info:    "{\"sum\": 5000, \"expired\": \"2023-06-16\"}",
			},
		})

		assert.Equal(t, &registrator.Certificate{
			Id:      777,
			Type:    registrator.CertificateType_CERTIFICATE_TYPE_FREE_PASS,
			WonOn:   1,
			SpentOn: 2,
			Info:    "{\"sum\": 5000, \"expired\": \"2023-06-16\"}",
		}, got)
		assert.NoError(t, err)
	})
}

func TestRegistrator_DeleteCertificate(t *testing.T) {
	t.Run("error certificate not found", func(t *testing.T) {
		fx := tearUp(t)

		fx.certificatesFacade.EXPECT().DeleteCertificate(fx.ctx, int32(1)).Return(model.ErrCertificateNotFound)

		got, err := fx.registrator.DeleteCertificate(fx.ctx, &registrator.DeleteCertificateRequest{
			Id: 1,
		})

		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.NotFound, st.Code())
		assert.Len(t, st.Details(), 2)
	})

	t.Run("some internal error", func(t *testing.T) {
		fx := tearUp(t)

		fx.certificatesFacade.EXPECT().DeleteCertificate(fx.ctx, int32(1)).Return(errors.New("some error"))

		got, err := fx.registrator.DeleteCertificate(fx.ctx, &registrator.DeleteCertificateRequest{
			Id: 1,
		})

		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.Internal, st.Code())
		assert.Len(t, st.Details(), 0)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		fx.certificatesFacade.EXPECT().DeleteCertificate(fx.ctx, int32(1)).Return(nil)

		got, err := fx.registrator.DeleteCertificate(fx.ctx, &registrator.DeleteCertificateRequest{
			Id: 1,
		})

		assert.Equal(t, &emptypb.Empty{}, got)
		assert.NoError(t, err)
	})
}

func TestRegistrator_ListCertificates(t *testing.T) {
	t.Run("error while list certificates", func(t *testing.T) {
		fx := tearUp(t)

		fx.certificatesFacade.EXPECT().ListCertificates(fx.ctx).Return(nil, errors.New("some error"))

		got, err := fx.registrator.ListCertificates(fx.ctx, &emptypb.Empty{})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.Internal, st.Code())
		assert.Len(t, st.Details(), 0)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		fx.certificatesFacade.EXPECT().ListCertificates(fx.ctx).Return([]model.Certificate{
			{
				ID:      1,
				Type:    pkgmodel.CertificateTypeFreePass,
				WonOn:   1,
				SpentOn: model.NewMaybeInt32(2),
				Info:    model.NewMaybeString("{}"),
			},
			{
				ID:      2,
				Type:    pkgmodel.CertificateTypeBarBillPayment,
				WonOn:   3,
				SpentOn: model.NewMaybeInt32(2),
				Info:    model.NewMaybeString("{}"),
			},
		}, nil)

		got, err := fx.registrator.ListCertificates(fx.ctx, &emptypb.Empty{})
		assert.ElementsMatch(t,
			[]*registrator.Certificate{
				{
					Id:      1,
					Type:    registrator.CertificateType_CERTIFICATE_TYPE_FREE_PASS,
					WonOn:   1,
					SpentOn: 2,
					Info:    "{}",
				},
				{
					Id:      2,
					Type:    registrator.CertificateType_CERTIFICATE_TYPE_BAR_BILL_PAYMENT,
					WonOn:   3,
					SpentOn: 2,
					Info:    "{}",
				},
			},
			got.GetCertificates())
		assert.NoError(t, err)
	})
}

func TestRegistrator_PatchCertificate(t *testing.T) {
	t.Run("validate error. invalid info json value", func(t *testing.T) {
		fx := tearUp(t)

		got, err := fx.registrator.PatchCertificate(fx.ctx, &registrator.PatchCertificateRequest{
			Certificate: &registrator.Certificate{
				Id:      1,
				Type:    registrator.CertificateType_CERTIFICATE_TYPE_BAR_BILL_PAYMENT,
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

		got, err := fx.registrator.PatchCertificate(fx.ctx, &registrator.PatchCertificateRequest{
			Certificate: &registrator.Certificate{
				Id:      1,
				Type:    registrator.CertificateType_CERTIFICATE_TYPE_INVALID,
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
			Type:    pkgmodel.CertificateTypeBarBillPayment,
			WonOn:   10,
			SpentOn: model.NewMaybeInt32(190),
			Info:    model.NewMaybeString("{}"),
		}, []string{"type", "spent_on"}).Return(model.Certificate{}, model.ErrCertificateNotFound)

		got, err := fx.registrator.PatchCertificate(fx.ctx, &registrator.PatchCertificateRequest{
			Certificate: &registrator.Certificate{
				Id:      1,
				Type:    registrator.CertificateType_CERTIFICATE_TYPE_BAR_BILL_PAYMENT,
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
			Type:    pkgmodel.CertificateTypeBarBillPayment,
			WonOn:   10,
			SpentOn: model.NewMaybeInt32(190),
			Info:    model.NewMaybeString("{}"),
		}, []string{"type", "spent_on"}).Return(model.Certificate{}, model.ErrWonOnGameNotFound)

		got, err := fx.registrator.PatchCertificate(fx.ctx, &registrator.PatchCertificateRequest{
			Certificate: &registrator.Certificate{
				Id:      1,
				Type:    registrator.CertificateType_CERTIFICATE_TYPE_BAR_BILL_PAYMENT,
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
			Type:    pkgmodel.CertificateTypeBarBillPayment,
			WonOn:   10,
			SpentOn: model.NewMaybeInt32(190),
			Info:    model.NewMaybeString("{}"),
		}, []string{"type", "spent_on"}).Return(model.Certificate{}, model.ErrSpentOnGameNotFound)

		got, err := fx.registrator.PatchCertificate(fx.ctx, &registrator.PatchCertificateRequest{
			Certificate: &registrator.Certificate{
				Id:      1,
				Type:    registrator.CertificateType_CERTIFICATE_TYPE_BAR_BILL_PAYMENT,
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
			Type:    pkgmodel.CertificateTypeBarBillPayment,
			WonOn:   10,
			SpentOn: model.NewMaybeInt32(190),
			Info:    model.NewMaybeString("{}"),
		}, []string{"type", "spent_on"}).Return(model.Certificate{
			ID:      1,
			Type:    pkgmodel.CertificateTypeBarBillPayment,
			WonOn:   100,
			SpentOn: model.NewMaybeInt32(190),
			Info:    model.NewMaybeString("some valid json"),
		}, nil)

		got, err := fx.registrator.PatchCertificate(fx.ctx, &registrator.PatchCertificateRequest{
			Certificate: &registrator.Certificate{
				Id:      1,
				Type:    registrator.CertificateType_CERTIFICATE_TYPE_BAR_BILL_PAYMENT,
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

		assert.Equal(t, &registrator.Certificate{
			Id:      1,
			Type:    registrator.CertificateType_CERTIFICATE_TYPE_BAR_BILL_PAYMENT,
			WonOn:   100,
			SpentOn: 190,
			Info:    "some valid json",
		}, got)
		assert.NoError(t, err)
	})
}

func Test_convertModelCertificateToProtoCertificate(t *testing.T) {
	type args struct {
		certificate model.Certificate
	}
	tests := []struct {
		name string
		args args
		want *registrator.Certificate
	}{
		{
			name: "tc1",
			args: args{
				certificate: model.Certificate{
					ID:      1,
					Type:    pkgmodel.CertificateTypeFreePass,
					WonOn:   2,
					SpentOn: model.NewMaybeInt32(100),
					Info:    model.NewMaybeString("{}"),
				},
			},
			want: &registrator.Certificate{
				Id:      1,
				Type:    registrator.CertificateType_CERTIFICATE_TYPE_FREE_PASS,
				WonOn:   2,
				SpentOn: 100,
				Info:    "{}",
			},
		},
		{
			name: "tc2",
			args: args{
				certificate: model.Certificate{
					ID:    1,
					Type:  pkgmodel.CertificateTypeFreePass,
					WonOn: 2,
					Info:  model.NewMaybeString("{}"),
				},
			},
			want: &registrator.Certificate{
				Id:    1,
				Type:  registrator.CertificateType_CERTIFICATE_TYPE_FREE_PASS,
				WonOn: 2,
				Info:  "{}",
			},
		},
		{
			name: "tc3",
			args: args{
				certificate: model.Certificate{
					ID:    1,
					Type:  pkgmodel.CertificateTypeFreePass,
					WonOn: 2,
				},
			},
			want: &registrator.Certificate{
				Id:    1,
				Type:  registrator.CertificateType_CERTIFICATE_TYPE_FREE_PASS,
				WonOn: 2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertModelCertificateToProtoCertificate(tt.args.certificate); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertModelCertificateToProtoCertificate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validateCreateCertificateRequest(t *testing.T) {
	type args struct {
		ctx context.Context
		req *registrator.CreateCertificateRequest
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
				req: &registrator.CreateCertificateRequest{
					Certificate: &registrator.Certificate{
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
				req: &registrator.CreateCertificateRequest{
					Certificate: &registrator.Certificate{
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
				req: &registrator.CreateCertificateRequest{
					Certificate: &registrator.Certificate{
						Type: registrator.CertificateType_CERTIFICATE_TYPE_FREE_PASS,
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
				req: &registrator.CreateCertificateRequest{
					Certificate: &registrator.Certificate{
						Type: registrator.CertificateType_CERTIFICATE_TYPE_INVALID,
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
				req: &registrator.CreateCertificateRequest{
					Certificate: &registrator.Certificate{
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
				req: &registrator.CreateCertificateRequest{
					Certificate: &registrator.Certificate{
						Type: registrator.CertificateType(pkgmodel.NumberOfCertificateTypes),
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
				req: &registrator.CreateCertificateRequest{
					Certificate: &registrator.Certificate{
						Type: registrator.CertificateType_CERTIFICATE_TYPE_FREE_PASS,
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

func Test_validatePatchCertificateRequest(t *testing.T) {
	type args struct {
		ctx context.Context
		req *registrator.PatchCertificateRequest
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
				req: &registrator.PatchCertificateRequest{
					Certificate: &registrator.Certificate{
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
				req: &registrator.PatchCertificateRequest{
					Certificate: &registrator.Certificate{
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
				req: &registrator.PatchCertificateRequest{
					Certificate: &registrator.Certificate{
						Type: registrator.CertificateType_CERTIFICATE_TYPE_FREE_PASS,
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
				req: &registrator.PatchCertificateRequest{
					Certificate: &registrator.Certificate{
						Type: registrator.CertificateType_CERTIFICATE_TYPE_INVALID,
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
				req: &registrator.PatchCertificateRequest{
					Certificate: &registrator.Certificate{
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
				req: &registrator.PatchCertificateRequest{
					Certificate: &registrator.Certificate{
						Type: registrator.CertificateType(pkgmodel.NumberOfCertificateTypes),
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
				req: &registrator.PatchCertificateRequest{
					Certificate: &registrator.Certificate{
						Type: registrator.CertificateType_CERTIFICATE_TYPE_FREE_PASS,
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
