package internal

import (
	"golang.org/x/text/language"
)

func CheckLanguageTag(locale string) (language.Tag, error) {
	return language.Parse(locale)
}
