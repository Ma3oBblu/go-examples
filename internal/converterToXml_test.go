package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvertIdsToXml(t *testing.T) {
	raw, err := ConvertIdsToXml([]string{"a", "b", "c"})
	assert.NoError(t, err)
	assert.Equal(t, "<smth><id>a</id><id>b</id><id>c</id></smth>", raw)
}
