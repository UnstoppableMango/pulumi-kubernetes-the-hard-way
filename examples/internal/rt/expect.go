package rt

import (
	"testing"

	"github.com/pulumi/pulumi/sdk/v3/go/common/apitype"
	"github.com/stretchr/testify/assert"
)

func ExpectOutput(t *testing.T, res apitype.ResourceV3, key string, value interface{}) {
	ExpectKey(t, res.Outputs, key, value)
}

func ExpectKey[T comparable, V any](t *testing.T, m map[T]V, key T, value V) {
	actual, ok := m[key]
	assert.Truef(t, ok, "Key `%s` was not set", key)
	assert.Equal(t, value, actual)
}
