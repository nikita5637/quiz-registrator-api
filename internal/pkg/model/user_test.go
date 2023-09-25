package model

import (
	"testing"
)

func TestUserState_ToSQL(t *testing.T) {
	tests := []struct {
		name string
		s    UserState
		want int
	}{
		{
			name: "tc1",
			s:    UserStateChangingName,
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.ToSQL(); got != tt.want {
				t.Errorf("UserState.ToSQL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidateUserState(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "error. not UserState",
			args: args{
				value: "not UserState value",
			},
			wantErr: true,
		},
		{
			name: "error. eq 0",
			args: args{
				value: UserState(0),
			},
			wantErr: true,
		},
		{
			name: "error. eq numberOfUserStates",
			args: args{
				value: numberOfUserStates,
			},
			wantErr: true,
		},
		{
			name: "ok",
			args: args{
				value: UserStateRegistered,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateUserState(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("ValidateUserState() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateSex(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "error. not Sex",
			args: args{
				value: "not sex value",
			},
			wantErr: true,
		},
		{
			name: "error. eq 0",
			args: args{
				value: Sex(0),
			},
			wantErr: true,
		},
		{
			name: "error. eq numberOfSexes",
			args: args{
				value: numberOfSexes,
			},
			wantErr: true,
		},
		{
			name: "ok",
			args: args{
				value: SexMale,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateSex(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("ValidateSex() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
