package internal

import (
	"mime"
)

// CheckMimeByExtension получение mime типа по расширению файла
func CheckMimeByExtension(extension string) string {
	return mime.TypeByExtension(extension)
}
