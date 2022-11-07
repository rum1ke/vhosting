package timedate

import (
	"fmt"
	"strings"
	"time"
)

func GetTimestamp() string {
	// get timestamp, separate gmt and time
	time := time.Now().Round(time.Microsecond).String()
	time = time[:len(time)-6]
	gmt := time[len(time)-3:]
	timeWithoutGMT := time[:len(time)-4]

	// adds a zero to the end of time until ms digits is 6
	digits := len(timeWithoutGMT) - 1 - strings.LastIndex(timeWithoutGMT, ".")
	for i := digits; i < 6; i++ {
		timeWithoutGMT += "0"
	}

	// assemble and return
	return timeWithoutGMT + gmt
}

func IsDateExpired(date string, expirationHours int) bool {
	expirationDate, err := time.Parse(time.RFC3339Nano, date)
	if err != nil {
		fmt.Printf("Cannot parse date. Date: %s.\n", date)
	}

	expirationDate = expirationDate.Add(time.Duration(expirationHours) * time.Hour)

	if expirationDate.After(time.Now()) {
		return false
	}

	return true
}
