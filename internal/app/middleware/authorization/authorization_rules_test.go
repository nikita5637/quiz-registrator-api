package authorization

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuthorizationGRPCRulesNumber(t *testing.T) {
	assert.Len(t, grpcRules, 41)
}
