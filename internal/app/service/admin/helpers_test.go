package admin

import (
	"context"
	"reflect"
	"testing"

	"github.com/nikita5637/quiz-registrator-api/internal/app/service/admin/mocks"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	adminpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/admin"
)

type fixture struct {
	ctx context.Context

	userRolesFacade *mocks.UserRolesFacade

	implementation *Implementation
}

func tearUp(t *testing.T) *fixture {
	fx := &fixture{
		ctx: context.Background(),

		userRolesFacade: mocks.NewUserRolesFacade(t),
	}

	fx.implementation = &Implementation{
		userRolesFacade: fx.userRolesFacade,
	}

	t.Cleanup(func() {})

	return fx
}

func Test_convertModelUserRoleToProtoUserRole(t *testing.T) {
	type args struct {
		userRole model.UserRole
	}
	tests := []struct {
		name string
		args args
		want *adminpb.UserRole
	}{
		{
			name: "tc1",
			args: args{
				userRole: model.UserRole{
					ID:     1,
					UserID: 1,
					Role:   model.RoleAdmin,
				},
			},
			want: &adminpb.UserRole{
				Id:     1,
				UserId: 1,
				Role:   adminpb.Role_ROLE_ADMIN,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertModelUserRoleToProtoUserRole(tt.args.userRole); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertModelUserRoleToProtoUserRole() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertProtoUserRoleToModelUserRole(t *testing.T) {
	type args struct {
		userRole *adminpb.UserRole
	}
	tests := []struct {
		name string
		args args
		want model.UserRole
	}{
		{
			name: "tc1",
			args: args{
				userRole: &adminpb.UserRole{
					Id:     1,
					UserId: 1,
					Role:   adminpb.Role_ROLE_ADMIN,
				},
			},
			want: model.UserRole{
				ID:     1,
				UserID: 1,
				Role:   model.RoleAdmin,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertProtoUserRoleToModelUserRole(tt.args.userRole); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertProtoUserRoleToModelUserRole() = %v, want %v", got, tt.want)
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
			name: "UserID",
			args: args{
				keys: []string{"UserID"},
			},
			want: &errorDetails{
				Reason: reasonInvalidUserID,
				Lexeme: invalidUserIDLexeme,
			},
		},
		{
			name: "Role",
			args: args{
				keys: []string{"Role"},
			},
			want: &errorDetails{
				Reason: reasonInvalidUserRole,
				Lexeme: invalidUserRoleLexeme,
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
