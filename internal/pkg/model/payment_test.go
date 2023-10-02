package model

import (
	"testing"
)

func TestValidatePayment(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "is not payment",
			args: args{
				value: "not payment",
			},
			wantErr: true,
		},
		{
			name: "lt 0",
			args: args{
				value: Payment(-1),
			},
			wantErr: true,
		},
		{
			name: "eq 0",
			args: args{
				value: Payment(0),
			},
			wantErr: true,
		},
		{
			name: "eq numberOfPayments",
			args: args{
				value: numberOfPayments,
			},
			wantErr: true,
		},
		{
			name: "PaymentCash",
			args: args{
				value: PaymentCash,
			},
			wantErr: false,
		},
		{
			name: "PaymentCertificate",
			args: args{
				value: PaymentCertificate,
			},
			wantErr: false,
		},
		{
			name: "PaymentMixed",
			args: args{
				value: PaymentMixed,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidatePayment(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("ValidatePayment() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
