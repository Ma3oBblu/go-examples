package internal

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"golang.org/x/text/language"
	"testing"
	"time"
)

func TestConvertIdsToXmlForOrder(t *testing.T) {
	raw, err := ConvertIdsToXmlForNews([]string{"a", "b", "c"})

	assert.NoError(t, err)
	assert.Equal(t, "<news><id>a</id><id>b</id><id>c</id></news>", raw)
}

func TestGetString(t *testing.T) {
	productIds := []string{"1139310", "1369917"}
	productIdsString := GetString(productIds)

	assert.Equal(t, "1139310|1369917", productIdsString)
}

func TestUrl(t *testing.T) {
	maxInt := int(^uint(0) >> 1)
	fmt.Println("MAX_INT: 9223372036854775807")
	fmt.Printf("maxInt go: %v", maxInt)
}

func TestLanguageCheck(t *testing.T) {
	locale := "ru"
	tag, err := language.Parse(locale)
	fmt.Printf("%s ==> ", tag)
	if err != nil {
		fmt.Println("Invalid locale: " + locale)
	}
}

func TestTimeZone(t *testing.T) {
	timezone := "Europe/Moscow"
	loc, err := time.LoadLocation(timezone)
	fmt.Printf("%s ==> ", loc)
	if err != nil {
		fmt.Println("Invalid timezone: " + timezone)
	}
}
