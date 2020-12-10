package internal

import (
	"fmt"
	"time"
)

// ParseDate парсинг даты с учетом нулевых значений
func ParseDate(birthDate string) time.Time {
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

// CheckTimeZone получает локацию по указанной таймзоне
func CheckTimeZone() (*time.Location, error) {
	timezone := "Europe/Moscow"
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		return nil, fmt.Errorf("invalid timezone: %v", timezone)
	}
	return loc, nil
}
