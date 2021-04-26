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

func TestGenTime(t *testing.T) {
	//loc, _ := time.LoadLocation("Europe/Moscow")
	loc := time.FixedZone("asdasd", 0)
	d := time.Date(2021, time.Month(4), 17, 0, 0, 0, 0, loc)
	fmt.Printf("%v", d.String())
	fmt.Printf("\n%v", d.Unix())
	fmt.Printf("\n%s", d.Format(time.RFC3339))

	t1 := time.Time{}
	var receivedMinute int
	// не все точки обладают информацией о времени получения товара
	receivedHour := 0

	t2 := time.Date(
		t1.Year(),
		t1.Month(),
		t1.Day(),
		receivedHour,
		receivedMinute,
		0,
		0,
		loc,
	)

	fmt.Printf("\n\n\n%v => %s => %v", t1.IsZero(), t2.String(), t2.IsZero())
}

func TestMultiCount(t *testing.T) {
	count := []int{1, 1, 1, 1, 10, 10, 40, 40, 50, 50, 50}
	multi := []int{0, 1, 10, 40, 10, 40, 10, 40, 10, 40, 11}
	for i := range multi {
		result := 0
		if multi[i] != 0 && count[i]%multi[i] != 0 {
			if count[i] <= multi[i] {
				result = multi[i]
				fmt.Printf("if    ===> \t|")
			} else {
				result = count[i] / multi[i] * multi[i]
				fmt.Printf("else  ===> \t|")
			}
		} else {
			fmt.Printf("count ===> \t|")
			result = count[i]
		}
		fmt.Printf("count -> %v multi -> %v result -> %v\n", count[i], multi[i], result)
	}
}
