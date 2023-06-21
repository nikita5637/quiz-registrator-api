package certificatemanager

import (
	"context"
	"reflect"
	"testing"

	"github.com/nikita5637/quiz-registrator-api/internal/app/service/certificate_manager/mocks"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	pkgmodel "github.com/nikita5637/quiz-registrator-api/pkg/model"
	certificatemanagerpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/certificate_manager"
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
					Type:    pkgmodel.CertificateTypeFreePass,
					WonOn:   2,
					SpentOn: model.NewMaybeInt32(100),
					Info:    model.NewMaybeString("{}"),
				},
			},
			want: &certificatemanagerpb.Certificate{
				Id:      1,
				Type:    certificatemanagerpb.CertificateType_CERTIFICATE_TYPE_FREE_PASS,
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
			want: &certificatemanagerpb.Certificate{
				Id:    1,
				Type:  certificatemanagerpb.CertificateType_CERTIFICATE_TYPE_FREE_PASS,
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
			want: &certificatemanagerpb.Certificate{
				Id:    1,
				Type:  certificatemanagerpb.CertificateType_CERTIFICATE_TYPE_FREE_PASS,
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
