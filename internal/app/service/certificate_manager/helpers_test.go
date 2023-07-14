package certificatemanager

import (
	"context"
	"reflect"
	"testing"

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
					SpentOn: model.NewMaybeInt32(100),
					Info:    model.NewMaybeString("{}"),
				},
			},
			want: &certificatemanagerpb.Certificate{
				Id:    1,
				Type:  certificatemanagerpb.CertificateType_CERTIFICATE_TYPE_FREE_PASS,
				WonOn: 2,
				SpentOn: &wrapperspb.Int32Value{
					Value: 100,
				},
				Info: &wrapperspb.StringValue{
					Value: "{}",
				},
			},
		},
		{
			name: "tc2",
			args: args{
				certificate: model.Certificate{
					ID:    1,
					Type:  model.CertificateTypeFreePass,
					WonOn: 2,
					Info:  model.NewMaybeString("{}"),
				},
			},
			want: &certificatemanagerpb.Certificate{
				Id:    1,
				Type:  certificatemanagerpb.CertificateType_CERTIFICATE_TYPE_FREE_PASS,
				WonOn: 2,
				Info: &wrapperspb.StringValue{
					Value: "{}",
				},
			},
		},
		{
			name: "tc3",
			args: args{
				certificate: model.Certificate{
					ID:    1,
					Type:  model.CertificateTypeFreePass,
					WonOn: 2,
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
					ID:    1,
					Type:  model.CertificateTypeFreePass,
					WonOn: 2,
					SpentOn: model.MaybeInt32{
						Valid: false,
						Value: 0,
					},
					Info: model.MaybeString{
						Valid: false,
						Value: "",
					},
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
				value: model.MaybeInt32{
					Valid: true,
					Value: 0,
				},
			},
			wantErr: true,
		},
		{
			name: "no error. eq 0 and not valid",
			args: args{
				value: model.MaybeInt32{
					Valid: false,
					Value: 0,
				},
			},
			wantErr: false,
		},
		{
			name: "error. lt minSpentOn and valid",
			args: args{
				value: model.MaybeInt32{
					Valid: true,
					Value: -1,
				},
			},
			wantErr: true,
		},
		{
			name: "no error. lt minSpentOn and not valid",
			args: args{
				value: model.MaybeInt32{
					Valid: false,
					Value: -1,
				},
			},
			wantErr: false,
		},
		{
			name: "ok",
			args: args{
				value: model.MaybeInt32{
					Valid: true,
					Value: minSpentOn,
				},
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
				value: model.MaybeString{
					Valid: true,
					Value: "",
				},
			},
			wantErr: true,
		},
		{
			name: "no error. empty string and not valid",
			args: args{
				value: model.MaybeString{
					Valid: false,
					Value: "",
				},
			},
			wantErr: false,
		},
		{
			name: "error. invalid JSON value and valid",
			args: args{
				value: model.MaybeString{
					Valid: true,
					Value: "invalid JSON",
				},
			},
			wantErr: true,
		},
		{
			name: "no error. invalid JSON value and not valid",
			args: args{
				value: model.MaybeString{
					Valid: false,
					Value: "invalid JSON",
				},
			},
			wantErr: false,
		},
		{
			name: "error. too long and valid",
			args: args{
				value: model.MaybeString{
					Valid: true,
					Value: "{\"a\":\"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa\"}",
				},
			},
			wantErr: true,
		},
		{
			name: "error. too long and not valid",
			args: args{
				value: model.MaybeString{
					Valid: false,
					Value: "{\"a\":\"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa\"}",
				},
			},
			wantErr: false,
		},
		{
			name: "ok",
			args: args{
				value: model.MaybeString{
					Valid: true,
					Value: "{}",
				},
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
