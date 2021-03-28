package internal

import (
	"fmt"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestTimeZone(t *testing.T) {
	loc, err := CheckTimeZone()
	if err != nil {
		fmt.Printf("Invalid timezone: %s", loc)
	}
	assert.NotNil(t, loc)
	assert.Equal(t, "Europe/Moscow", loc.String())
}

func TestZeroDates(t *testing.T) {
	emptyDateStr := ""
	emptyDateTime := ParseDate(emptyDateStr)
	fmt.Printf("converted date: %s\n", emptyDateTime)
	if emptyDateTime.IsZero() {
		fmt.Printf("this date is zero \n")
	}
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

func TestParseDateAnother(t *testing.T) {
	dateString := "18 May 2020 09:58:5"
	loc, _ := time.LoadLocation("Europe/Moscow")
	date, err := time.ParseInLocation("02 Jan 2006 15:04:5", dateString, loc)
	if err != nil {
		fmt.Printf("%s", err)
	}
	//date = date.Add(3 * time.Hour)
	// parsed string 2020-05-18 09:58:05 +0000 UTC, unix 1589795885, location UTC
	fmt.Printf("parsed string %s, unix %v, location %s", date.String(), date.Unix(), date.Location())
}

func TestParseDateTime(t *testing.T) {
	rawDate := "2021-01-14 00:00:00"
	date, err := time.Parse("2006-01-02 00:00:00", rawDate)
	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Printf("%v", date.Unix())
}
