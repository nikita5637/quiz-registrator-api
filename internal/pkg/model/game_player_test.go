package model

import (
	"testing"
)

func TestValidateDegree(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "not degree",
			args: args{
				value: "not degree",
			},
			wantErr: true,
		},
		{
			name: "invalid value",
			args: args{
				value: DegreeInvalid,
			},
			wantErr: true,
		},
		{
			name: "lt 0",
			args: args{
				value: Degree(-1),
			},
			wantErr: true,
		},
		{
			name: "eq numberOfDegrees",
			args: args{
				value: numberOfDegrees,
			},
			wantErr: true,
		},
		{
			name: "ok",
			args: args{
				value: DegreeLikely,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateDegree(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("ValidateDegree() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
