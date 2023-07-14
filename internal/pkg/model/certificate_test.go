package model

import (
	"testing"
)

func TestCertificateType_ToSQL(t *testing.T) {
	tests := []struct {
		name string
		tr   CertificateType
		want uint8
	}{
		{
			name: "tc1",
			tr:   CertificateTypeInvalid,
			want: 0,
		},
		{
			name: "tc2",
			tr:   CertificateTypeFreePass,
			want: 1,
		},
		{
			name: "tc3",
			tr:   CertificateTypeBarBillPayment,
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tr.ToSQL(); got != tt.want {
				t.Errorf("CertificateType.ToSQL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidateCertificateType(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "error. not CertificateType",
			args: args{
				value: "not CertificateType",
			},
			wantErr: true,
		},
		{
			name: "error. gt max value",
			args: args{
				value: numberOfCertificateTypes,
			},
			wantErr: true,
		},
		{
			name: "invalid certificate type",
			args: args{
				value: CertificateTypeInvalid,
			},
			wantErr: false,
		},
		{
			name: "ok",
			args: args{
				value: CertificateType(CertificateTypeBarBillPayment),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateCertificateType(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("ValidateCertificateType() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
