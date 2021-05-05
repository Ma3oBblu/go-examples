package internal

import (
	"golang.org/x/crypto/bcrypt"
	"regexp"
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

// Регулярное выражения для поиска последовательности длиной от 3-х цифр обрамленных спец. символами
var maskRegexp = regexp.MustCompile(`(?:[\s,;]|^)(\d{3,})(?:[\s,;.?!]|$)`)

func MaskSensitiveDataRunes(text string) string {
	groupsPositionsList := maskRegexp.FindAllStringSubmatchIndex(text, -1)
	for i := len(groupsPositionsList) - 1; i >= 0; i-- {
		groupsPositions := groupsPositionsList[i]

		head := text[:groupsPositions[2]]
		tail := text[groupsPositions[3]:]
		group1 := text[groupsPositions[2]:groupsPositions[3]]
		group1Runes := []rune(group1)
		if len(group1Runes) < 3 {
			continue
		}

		dots := strings.Repeat("*", len(group1Runes)-2)
		text = head + string(group1Runes[0]) + dots + string(group1Runes[len(group1)-1]) + tail
	}

	return text
}
