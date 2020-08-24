package internal

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"strings"
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
