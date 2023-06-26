package model

import "testing"

func TestValidateLeague(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "error. not int",
			args: args{
				value: "not int value",
			},
			wantErr: true,
		},
		{
			name: "error. gt max value",
			args: args{
				value: int(numberOfLeagues),
			},
			wantErr: true,
		},
		{
			name: "error. eq 0",
			args: args{
				value: int(0),
			},
			wantErr: true,
		},
		{
			name: "ok",
			args: args{
				value: LeagueQuizPlease,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateLeague(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("ValidateLeague() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
