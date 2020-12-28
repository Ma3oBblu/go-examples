package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetString(t *testing.T) {
	ids := []string{"1139310", "1369917"}
	idsStr := GetString(ids)
	assert.Equal(t, "1139310|1369917", idsStr)
}

func TestGetPasswordHash(t *testing.T) {
	hash, err := GetPasswordHash("old_password123")
	assert.Nil(t, err)
	println(string(hash))
}