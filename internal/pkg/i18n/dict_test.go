package i18n

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_DictLength(t *testing.T) {
	assert.Len(t, dictionary, 1)
	assert.Len(t, dictionary["ru"], 47)
}
