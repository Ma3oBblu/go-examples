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
