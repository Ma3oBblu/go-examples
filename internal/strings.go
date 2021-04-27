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
	runeStr := []rune(str)
	for i := range runeStr {
		if i == 0 || i == len(runeStr)-1 {
			continue
		}
		runeStr[i] = '*'
	}
	return string(runeStr)
}

func ConvertSymbolsWithBytes(str string) string {
	bytes := make([]byte, len(str))
	for i, k := range []byte(str) {
		if i == 0 || i == len(str)-1 {
			bytes[i] = k
			continue
		}
		bytes[i] = []byte("*")[0]
	}
	return string(bytes)
}

func ConvertSymbolsWithSlicing(str string) string {
	return string(str[0]) + strings.Repeat("*", len(str)-2) + string(str[len(str)-1])
}
