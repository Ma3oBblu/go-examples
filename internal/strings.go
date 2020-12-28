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
