package internal

import (
	"golang.org/x/crypto/bcrypt"
	"strings"
)

// GetString проверка работы join метода с заданным разделителем
func GetString(ids []string) string {
	return strings.Join(ids, "|")
}

// GetPasswordHash формирует хеш для заданного пароля
func GetPasswordHash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func ConvertSymbolsWithRune(str string) string {
	result := make([]rune, len(str))
	for i := range []rune(str) {
		if i == 0 || i == len(str)-1 || rune(str[i]) == '@' {
			result[i] = rune(str[i])
			continue
		}
		result[i] = '*'
	}
	return string(result)
}

func ConvertSymbolsWithBytes(str string) string {
	bytes := []byte(str)
	for i, k := range bytes {
		if i == 0 || i == len(bytes)-1 || string(k) == "@" {
			continue
		}
		bytes[i] = []byte("*")[0]
	}
	return string(bytes)
}

func ConvertSymbolsWithSlicing(str string) string {
	return string(str[0]) + strings.Repeat("*", len(str)-2) + string(str[len(str)-1])
}
