package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheckMimeByExtension(t *testing.T) {
	extension := ".jpg"
	mimeType := CheckMimeByExtension(extension)
	assert.Equal(t, "image/jpeg", mimeType)

	extension = "jpg"
	mimeType = CheckMimeByExtension(extension)
	assert.NotEqual(t, "image/jpeg", mimeType)
	assert.Empty(t, mimeType)
}
