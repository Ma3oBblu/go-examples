package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetMaxInt(t *testing.T) {
	maxInt := GetMaxInt()
	assert.Equal(t, 9223372036854775807, maxInt)
}
