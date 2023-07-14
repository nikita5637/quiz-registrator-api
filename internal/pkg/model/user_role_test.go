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
			name: "invalid",
			r:    RoleInvalid,
			want: invalid,
		},
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
			r:    Role(4),
			want: invalid,
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
