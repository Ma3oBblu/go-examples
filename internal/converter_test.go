package internal

import (
	"fmt"
	"github.com/golang/protobuf/ptypes/timestamp"
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

func TestZeroDates(t *testing.T) {
	emptyDateStr := ""
	emptyDateTime := BirthDate(emptyDateStr)
	fmt.Printf("converted date: %s\n", emptyDateTime)
	fmt.Printf("formatted date: %s\n", emptyDateTime.Format("2006-01-02T00:00:00Z"))
	fmt.Printf("formatted unix: %v\n", emptyDateTime.Unix())
	timestampConverted := &timestamp.Timestamp{Seconds: emptyDateTime.Unix(), Nanos: int32(emptyDateTime.UnixNano())}
	fmt.Printf("timestamp: %s\n", timestampConverted)

	zeroTimeStamp := &timestamp.Timestamp{Seconds: 0, Nanos: 0}
	zeroTimeFromUnix := time.Unix(zeroTimeStamp.Seconds, int64(zeroTimeStamp.Nanos))
	fmt.Printf("zeroTimeFromUnix: %s\n", zeroTimeFromUnix)
	fmt.Printf("zeroTimeFromUnix formatted: %s\n", zeroTimeFromUnix.Format("2006-01-02T00:00:00Z"))

	zeroTime := time.Time{}
	fmt.Printf("zeroTime formatted: %s\n", zeroTime.Format("2006-01-02T00:00:00Z"))
}
