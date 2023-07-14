package authentication

import (
	"context"
	"testing"

	"google.golang.org/grpc/metadata"
)

func Test_getAuthenticationType(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "there are not metadata",
			args: args{
				ctx: context.Background(),
			},
			want: "",
		},
		{
			name: "authentication type Telegram ID",
			args: args{
				ctx: metadata.NewIncomingContext(context.Background(), metadata.New(
					map[string]string{
						telegramClientIDHeader: "1",
					},
				)),
			},
			want: authenticationTypeTelegramID,
		},
		{
			name: "authentication type service name via serviceNameHeader",
			args: args{
				ctx: metadata.NewIncomingContext(context.Background(), metadata.New(
					map[string]string{
						serviceNameHeader: "service name",
					},
				)),
			},
			want: authenticationTypeServiceName,
		},
		{
			name: "authentication type service name via moduleNameHeader",
			args: args{
				ctx: metadata.NewIncomingContext(context.Background(), metadata.New(
					map[string]string{
						moduleNameHeader: "module name",
					},
				)),
			},
			want: authenticationTypeServiceName,
		},
		{
			name: "default",
			args: args{
				ctx: metadata.NewIncomingContext(context.Background(), metadata.New(
					map[string]string{},
				)),
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getAuthenticationType(tt.args.ctx); got != tt.want {
				t.Errorf("getAuthenticationType() = %v, want %v", got, tt.want)
			}
		})
	}
}
