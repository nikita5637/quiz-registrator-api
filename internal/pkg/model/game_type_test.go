package model

import (
	"testing"

	gamepb "github.com/nikita5637/quiz-registrator-api/pkg/pb/game"
	"github.com/stretchr/testify/assert"
)

func Test_matchGameTypeWithProto(t *testing.T) {
	assert.Equal(t, GameType(gamepb.GameType_GAME_TYPE_CLASSIC), GameTypeClassic)
	assert.Equal(t, GameType(gamepb.GameType_GAME_TYPE_THEMATIC), GameTypeThematic)
	assert.Equal(t, GameType(gamepb.GameType_GAME_TYPE_MOVIES_AND_MUSIC), GameTypeMoviesAndMusic)
	assert.Equal(t, GameType(gamepb.GameType_GAME_TYPE_CLOSED), GameTypeClosed)
	assert.Equal(t, GameType(gamepb.GameType_GAME_TYPE_THEMATIC_MOVIES_AND_MUSIC), GameTypeThematicMoviesAndMusic)
	assert.Len(t, []GameType{
		GameTypeClassic,
		GameTypeThematic,
		GameTypeMoviesAndMusic,
		GameTypeClosed,
		GameTypeThematicMoviesAndMusic,
	}, len(gamepb.GameType_name)-1)
}

func TestValidateGameType(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "is not GameType",
			args: args{
				"is not GameType",
			},
			wantErr: true,
		},
		{
			name: "lt 0",
			args: args{
				GameType(-1),
			},
			wantErr: true,
		},
		{
			name: "eq 0",
			args: args{
				GameType(0),
			},
			wantErr: true,
		},
		{
			name: "GameTypeClassic",
			args: args{
				GameTypeClassic,
			},
			wantErr: false,
		},
		{
			name: "GameTypeThematic",
			args: args{
				GameTypeThematic,
			},
			wantErr: false,
		},
		{
			name: "GameTypeEnglish",
			args: args{
				GameTypeEnglish,
			},
			wantErr: false,
		},
		{
			name: "GameTypeMoviesAndMusic",
			args: args{
				GameTypeMoviesAndMusic,
			},
			wantErr: false,
		},
		{
			name: "GameTypeClosed",
			args: args{
				GameTypeClosed,
			},
			wantErr: false,
		},
		{
			name: "GameTypeThematicMoviesAndMusic",
			args: args{
				GameTypeThematicMoviesAndMusic,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateGameType(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("ValidateGameType() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
