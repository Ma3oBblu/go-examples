package internal

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"strings"
	"time"
)

// ConvertIdsToXmlForNews конвертирует последовательность строк вида []string{"a", "b", "c"} в xml вида <news><id>a</id><id>b</id><id>c</id></news>
func ConvertIdsToXmlForNews(ids []string) (string, error) {
	buffer := bytes.NewBuffer(nil)
	encoder := xml.NewEncoder(buffer)
	idWrapper := struct {
		XMLName struct{} `xml:"id"`
		Id      string   `xml:",chardata"`
	}{}

	for _, id := range ids {
		idWrapper.Id = id

		err := encoder.Encode(idWrapper)
		if err != nil {
			return "", fmt.Errorf("can't marshall struct to xml: %s", err)
		}
	}

	return "<news>" + buffer.String() + "</news>", nil
}

func GetString(productIds []string) string {
	return strings.Join(productIds, "|")
}

// birthDate дата рождения
func BirthDate(birthDate string) time.Time {
	date, err := time.Parse("2006-01-02T00:00:00Z", birthDate)
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return time.Time{}
	}

	if date.Year() < 1901 {
		fmt.Printf("erro: < 1901\n")
		return time.Time{}
	}

	return date
}
