package internal

import (
	"fmt"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/stretchr/testify/assert"
	"golang.org/x/text/language"
	"regexp"
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

func TestRegexp(t *testing.T) {
	var emailRegex = regexp.MustCompile("^(?i)(?:[a-zа-я0-9!#$%&'*+/=?^_`{|}~-]+(?:\\.[a-zа-я0-9!#$%&'*+/=?^_`{|}~-]+)*|\"(?:[\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x21\\x23-\\x5b\\x5d-\\x7f]|\\\\[\\x01-\\x09\\x0b\\x0c\\x0e-\\x7f])*\")@(?:(?:[a-zа-я0-9](?:[a-zа-я0-9-]*[a-zа-я0-9])?\\.)+[a-zа-я0-9](?:[a-zа-я0-9-]*[a-zа-я0-9])?|\\[(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?|[a-zа-я0-9-]*[a-zа-я0-9]:(?:[\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x21-\\x5a\\x53-\\x7f]|\\\\[\\x01-\\x09\\x0b\\x0c\\x0e-\\x7f])+)\\])$")
	fmt.Printf("%s", emailRegex)
}

//BenchmarkBool
//BenchmarkBool/check_text1_is_exist_in_bool_var
//BenchmarkBool/check_text1_is_exist_in_bool_var-8         	181123724	         6.56 ns/op	       0 B/op	       0 allocs/op
//BenchmarkBool/check_text3_is_not_exist_in_bool_var
//BenchmarkBool/check_text3_is_not_exist_in_bool_var-8     	75207015	        15.4 ns/op	       0 B/op	       0 allocs/op
func BenchmarkBool(b *testing.B) {
	boolVar := map[string]bool{
		"text1": true,
		"text2": true,
	}
	testsForBool := []struct {
		name         string
		checkedValue string
		boolVar      map[string]bool
	}{
		{
			"check text1 is exist in bool var",
			"text1",
			boolVar,
		},
		{
			"check text3 is not exist in bool var",
			"text3",
			boolVar,
		},
	}
	for _, tt := range testsForBool {
		b.Run(tt.name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				if _, ok := tt.boolVar[tt.checkedValue]; ok {
					continue
				}
			}
		})
	}
}

//BenchmarkStruct
//BenchmarkStruct/check_text1_is_exist_in_struct_var
//BenchmarkStruct/check_text1_is_exist_in_struct_var-8         	183780481	         6.50 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStruct/check_text3_is_not_exist_in_struct_var
//BenchmarkStruct/check_text3_is_not_exist_in_struct_var-8     	74461951	        15.3 ns/op	       0 B/op	       0 allocs/op
func BenchmarkStruct(b *testing.B) {
	structVar := map[string]struct{}{
		"text1": {},
		"text2": {},
	}
	testsForStruct := []struct {
		name         string
		checkedValue string
		structVar    map[string]struct{}
	}{
		{
			"check text1 is exist in struct var",
			"text1",
			structVar,
		},
		{
			"check text3 is not exist in struct var",
			"text3",
			structVar,
		},
	}
	for _, tt := range testsForStruct {
		b.Run(tt.name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				if _, ok := tt.structVar[tt.checkedValue]; ok {
					continue
				}
			}
		})
	}
}
