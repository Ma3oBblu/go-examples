package internal

import "strings"

// GetString проверка работы join метода с заданным разделителем
func GetString(ids []string) string {
	return strings.Join(ids, "|")
}
