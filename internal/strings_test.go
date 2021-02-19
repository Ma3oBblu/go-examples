package internal

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
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

func TestInterfaceSlice(t *testing.T) {
	resp := make([]interface{}, 0)
	if len(resp) > 0 {
		t.Errorf("not empty string")
	}
}

func TestString(t *testing.T) {
	str := "91 123 34 73 100 34 58 34 49 50 51 34 44 34 85 115 101 114 73 100 34 58 34 34 44 34 80 114 111 100 117 99 116 73 100 34 58 34 34 44 34 84 105 116 108 101 34 58 34 34 44 34 84 101 120 116 34 58 34 34 44 34 83 116 97 116 117 115 34 58 48 44 34 82 97 116 105 110 103 34 58 48 44 34 82 97 116 105 110 103 67 111 117 110 116 34 58 48 44 34 67 111 109 109 101 110 116 67 111 117 110 116 34 58 48 44 34 86 105 101 119 115 67 111 117 110 116 34 58 48 44 34 67 114 101 97 116 105 111 110 84 105 109 101 34 58 110 117 108 108 44 34 73 115 66 108 111 99 107 101 100 34 58 102 97 108 115 101 44 34 84 111 116 97 108 34 58 48 125 44 123 34 73 100 34 58 34 49 50 51 52 34 44 34 85 115 101 114 73 100 34 58 34 34 44 34 80 114 111 100 117 99 116 73 100 34 58 34 34 44 34 84 105 116 108 101 34 58 34 34 44 34 84 101 120 116 34 58 34 34 44 34 83 116 97 116 117 115 34 58 48 44 34 82 97 116 105 110 103 34 58 48 44 34 82 97 116 105 110 103 67 111 117 110 116 34 58 48 44 34 67 111 109 109 101 110 116 67 111 117 110 116 34 58 48 44 34 86 105 101 119 115 67 111 117 110 116 34 58 48 44 34 67 114 101 97 116 105 111 110 84 105 109 101 34 58 110 117 108 108 44 34 73 115 66 108 111 99 107 101 100 34 58 102 97 108 115 101 44 34 84 111 116 97 108 34 58 48 125 93"
	replaced := strings.ReplaceAll(str, " ", ",")
	fmt.Printf("%s", replaced)
}

func TestTrim(t *testing.T) {
	invalidString := "     Ruslan      "
	fmt.Printf("Trim ====>%s<=\n", strings.Trim(invalidString, " "))
	fmt.Printf("TrimLeft ====>%s<=\n", strings.TrimLeft(invalidString, " "))
	fmt.Printf("TrimRight ====>%s<=\n", strings.TrimRight(invalidString, " "))
}
