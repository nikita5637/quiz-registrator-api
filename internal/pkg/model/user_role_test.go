package model

import (
	"testing"
)

func TestRole_String(t *testing.T) {
	tests := []struct {
		name string
		r    Role
		want string
	}{
		{
			name: "admin",
			r:    RoleAdmin,
			want: admin,
		},
		{
			name: "management",
			r:    RoleManagement,
			want: management,
		},
		{
			name: "user",
			r:    RoleUser,
			want: user,
		},
		{
			name: "not existed",
			r:    numberOfRoles,
			want: "invalid",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.String(); got != tt.want {
				t.Errorf("Role.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidateRole(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "is not Role",
			args: args{
				value: "not Role",
			},
			wantErr: true,
		},
		{
			name: "error. gt max value",
			args: args{
				value: numberOfRoles,
			},
			wantErr: true,
		},
		{
			name: "eq 0",
			args: args{
				value: Role(0),
			},
			wantErr: true,
		},
		{
			name: "ok",
			args: args{
				value: RoleAdmin,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateRole(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("ValidateRole() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
