package internal

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestGetMaxInt(t *testing.T) {
	maxInt := GetMaxInt()
	assert.Equal(t, 9223372036854775807, maxInt)
}

func TestItoa(t *testing.T) {
	number := 123
	itoaResult := strconv.Itoa(number)
	fmt.Printf("itoa result => %s, %T\n", itoaResult, itoaResult)
	formatIntResult := strconv.FormatInt(int64(number), 10)
	fmt.Printf("formatInt result => %s, %T\n", formatIntResult, formatIntResult)
}
