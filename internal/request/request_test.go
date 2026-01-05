package request

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringEquals(t *testing.T) {
	assert.Equal(t, "testText", "testText")
}