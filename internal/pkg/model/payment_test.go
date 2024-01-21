package model

import (
	"reflect"
	"testing"
)

func TestPayment_String(t *testing.T) {
	tests := []struct {
		name string
		p    Payment
		want string
	}{
		{
			name: "cash",
			p:    PaymentCash,
			want: "cash",
		},
		{
			name: "certificate",
			p:    PaymentCertificate,
			want: "certificate",
		},
		{
			name: "mixed",
			p:    PaymentMixed,
			want: "mixed",
		},
		{
			name: "invalid",
			p:    numberOfPayments,
			want: "invalid payment",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.String(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Payment.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
