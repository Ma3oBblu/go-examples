package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvertIdsToXmlForOrder(t *testing.T) {
	raw, err := ConvertIdsToXmlForNews([]string{"a", "b", "c"})

	assert.NoError(t, err)
	assert.Equal(t, "<news><id>a</id><id>b</id><id>c</id></news>", raw)
}

func TestGetString(t *testing.T) {
	productIds := []string{"1139310", "1369917"}
	productIdsString := GetString(productIds)

	assert.Equal(t, "1139310|1369917", productIdsString)
}
