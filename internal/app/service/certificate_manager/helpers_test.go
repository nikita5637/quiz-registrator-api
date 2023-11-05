package certificatemanager

import (
	"context"
	"reflect"
	"testing"

	"github.com/mono83/maybe"
	"github.com/nikita5637/quiz-registrator-api/internal/app/service/certificate_manager/mocks"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	certificatemanagerpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/certificate_manager"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type fixture struct {
	ctx context.Context

	certificatesFacade *mocks.CertificatesFacade

	certificateManager *CertificateManager
}

func tearUp(t *testing.T) *fixture {
	fx := &fixture{
		ctx: context.Background(),

		certificatesFacade: mocks.NewCertificatesFacade(t),
	}

	fx.certificateManager = New(Config{
		CertificatesFacade: fx.certificatesFacade,
	})

	t.Cleanup(func() {})

	return fx
}

func Test_convertModelCertificateToProtoCertificate(t *testing.T) {
	type args struct {
		certificate model.Certificate
	}
	tests := []struct {
		name string
		args args
		want *certificatemanagerpb.Certificate
	}{
		{
			name: "tc1",
			args: args{
				certificate: model.Certificate{
					ID:      1,
					Type:    model.CertificateTypeFreePass,
					WonOn:   2,
					SpentOn: maybe.Just(int32(100)),
					Info:    maybe.Just("{}"),
				},
			},
			want: &certificatemanagerpb.Certificate{
				Id:      1,
				Type:    certificatemanagerpb.CertificateType_CERTIFICATE_TYPE_FREE_PASS,
				WonOn:   2,
				SpentOn: wrapperspb.Int32(100),
				Info:    wrapperspb.String("{}"),
			},
		},
		{
			name: "tc2",
			args: args{
				certificate: model.Certificate{
					ID:      1,
					Type:    model.CertificateTypeFreePass,
					WonOn:   2,
					SpentOn: maybe.Nothing[int32](),
					Info:    maybe.Just("{}"),
				},
			},
			want: &certificatemanagerpb.Certificate{
				Id:    1,
				Type:  certificatemanagerpb.CertificateType_CERTIFICATE_TYPE_FREE_PASS,
				WonOn: 2,
				Info:  wrapperspb.String("{}"),
			},
		},
		{
			name: "tc3",
			args: args{
				certificate: model.Certificate{
					ID:      1,
					Type:    model.CertificateTypeFreePass,
					WonOn:   2,
					SpentOn: maybe.Nothing[int32](),
					Info:    maybe.Nothing[string](),
				},
			},
			want: &certificatemanagerpb.Certificate{
				Id:    1,
				Type:  certificatemanagerpb.CertificateType_CERTIFICATE_TYPE_FREE_PASS,
				WonOn: 2,
			},
		},
		{
			name: "tc4",
			args: args{
				certificate: model.Certificate{
					ID:      1,
					Type:    model.CertificateTypeFreePass,
					WonOn:   2,
					SpentOn: maybe.Nothing[int32](),
					Info:    maybe.Nothing[string](),
				},
			},
			want: &certificatemanagerpb.Certificate{
				Id:      1,
				Type:    certificatemanagerpb.CertificateType_CERTIFICATE_TYPE_FREE_PASS,
				WonOn:   2,
				SpentOn: nil,
				Info:    nil,
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

func Test_convertProtoCertificateToModelCertificate(t *testing.T) {
	type args struct {
		certificate *certificatemanagerpb.Certificate
	}
	tests := []struct {
		name string
		args args
		want model.Certificate
	}{
		{
			name: "tc1",
			args: args{
				certificate: &certificatemanagerpb.Certificate{
					Id:      1,
					Type:    certificatemanagerpb.CertificateType_CERTIFICATE_TYPE_BAR_BILL_PAYMENT,
					WonOn:   1,
					SpentOn: wrapperspb.Int32(2),
					Info:    wrapperspb.String("{}"),
				},
			},
			want: model.Certificate{
				ID:      1,
				Type:    model.CertificateTypeBarBillPayment,
				WonOn:   1,
				SpentOn: maybe.Just(int32(2)),
				Info:    maybe.Just("{}"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertProtoCertificateToModelCertificate(tt.args.certificate); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertProtoCertificateToModelCertificate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getErrorDetails(t *testing.T) {
	type args struct {
		keys []string
	}
	tests := []struct {
		name string
		args args
		want *errorDetails
	}{
		{
			name: "keys is nil",
			args: args{
				keys: nil,
			},
			want: nil,
		},
		{
			name: "keys is empty",
			args: args{
				keys: []string{},
			},
			want: nil,
		},
		{
			name: "Type",
			args: args{
				keys: []string{"Type"},
			},
			want: &errorDetails{
				Reason: reasonInvalidCertificateType,
				Lexeme: invalidCertificateTypeLexeme,
			},
		},
		{
			name: "WonOn",
			args: args{
				keys: []string{"WonOn"},
			},
			want: &errorDetails{
				Reason: reasonInvalidWonOnGameID,
				Lexeme: invalidWonOnGameIDLexeme,
			},
		},
		{
			name: "SpentOn",
			args: args{
				keys: []string{"SpentOn"},
			},
			want: &errorDetails{
				Reason: reasonInvalidSpentOnGameID,
				Lexeme: invalidSpentOnGameIDLexeme,
			},
		},
		{
			name: "Info",
			args: args{
				keys: []string{"Info"},
			},
			want: &errorDetails{
				Reason: reasonInvalidInfo,
				Lexeme: invalidInfoLexeme,
			},
		},
		{
			name: "not found",
			args: args{
				keys: []string{"not found"},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getErrorDetails(tt.args.keys); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getErrorDetails() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validateSpentOn(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "error. not MaybeInt32",
			args: args{
				value: "not MaybeInt32",
			},
			wantErr: true,
		},
		{
			name: "error. eq 0 and valid",
			args: args{
				value: maybe.Just(0),
			},
			wantErr: true,
		},
		{
			name: "no error. eq 0 and not valid",
			args: args{
				value: maybe.Nothing[int32](),
			},
			wantErr: false,
		},
		{
			name: "error. lt minSpentOn and valid",
			args: args{
				value: maybe.Just(-1),
			},
			wantErr: true,
		},
		{
			name: "ok",
			args: args{
				value: maybe.Just(minSpentOn),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateSpentOn(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("validateSpentOn() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_validateCertificateInfo(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "error. not MaybeString",
			args: args{
				value: "not MaybeString",
			},
			wantErr: true,
		},
		{
			name: "error. empty string and valid",
			args: args{
				value: maybe.Just(""),
			},
			wantErr: true,
		},
		{
			name: "no error. empty string and not valid",
			args: args{
				value: maybe.Nothing[string](),
			},
			wantErr: false,
		},
		{
			name: "error. invalid JSON value and valid",
			args: args{
				value: maybe.Just("invalid JSON"),
			},
			wantErr: true,
		},
		{
			name: "error. too long and valid",
			args: args{
				value: maybe.Just("{\"a\":\"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa\"}"),
			},
			wantErr: true,
		},
		{
			name: "ok",
			args: args{
				value: maybe.Just("{}"),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateCertificateInfo(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("validateCertificateInfo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
