package model

import (
	"testing"

	commonpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/common"
	"github.com/stretchr/testify/assert"
)

func Test_matchGameTypeWithProto(t *testing.T) {
	assert.Equal(t, GameType(commonpb.GameType_GAME_TYPE_INVALID), GameTypeInvalid)
	assert.Equal(t, GameType(commonpb.GameType_GAME_TYPE_CLASSIC), GameTypeClassic)
	assert.Equal(t, GameType(commonpb.GameType_GAME_TYPE_THEMATIC), GameTypeThematic)
	assert.Equal(t, GameType(commonpb.GameType_GAME_TYPE_MOVIES_AND_MUSIC), GameTypeMoviesAndMusic)
	assert.Equal(t, GameType(commonpb.GameType_GAME_TYPE_CLOSED), GameTypeClosed)
	assert.Equal(t, GameType(commonpb.GameType_GAME_TYPE_THEMATIC_MOVIES_AND_MUSIC), GameTypeThematicMoviesAndMusic)
	assert.Len(t, []GameType{
		GameTypeInvalid,
		GameTypeClassic,
		GameTypeThematic,
		GameTypeMoviesAndMusic,
		GameTypeClosed,
		GameTypeThematicMoviesAndMusic,
	}, len(commonpb.GameType_name))
}
