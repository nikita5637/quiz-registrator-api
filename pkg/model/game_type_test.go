package model

import (
	"testing"

	commonpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/common"
	"github.com/stretchr/testify/assert"
)

func Test_matchGameTypeWithProto(t *testing.T) {
	assert.Equal(t, int32(commonpb.GameType_GAME_TYPE_CLASSIC), GameTypeClassic)
	assert.Equal(t, int32(commonpb.GameType_GAME_TYPE_THEMATIC), GameTypeThematic)
	assert.Equal(t, int32(commonpb.GameType_GAME_TYPE_MOVIES_AND_MUSIC), GameTypeMoviesAndMusic)
	assert.Equal(t, int32(commonpb.GameType_GAME_TYPE_CLOSED), GameTypeClosed)
	assert.Equal(t, int32(commonpb.GameType_GAME_TYPE_THEMATIC_MOVIES_AND_MUSIC), GameTypeThematicMoviesAndMusic)
	assert.Len(t, []int32{
		GameTypeClassic,
		GameTypeThematic,
		GameTypeMoviesAndMusic,
		GameTypeClosed,
		GameTypeThematicMoviesAndMusic,
	}, len(commonpb.GameType_name)-1)
}