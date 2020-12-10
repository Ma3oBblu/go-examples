package internal

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLanguageCheck(t *testing.T) {
	locale := "ru"
	tag, err := CheckLanguageTag(locale)
	if err != nil {
		fmt.Printf("Invalid locale: %s", locale)
	}
	assert.Equal(t, "ru", tag.String())
	region, _ := tag.Region()
	assert.Equal(t, "RU", region.String())
}
