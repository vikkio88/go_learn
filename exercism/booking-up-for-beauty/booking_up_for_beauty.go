package booking

import (
	"fmt"
	"time"
)

const (
	baseFormat        = "1/2/2006 15:04:05"
	complexFormat     = "January 2, 2006 15:04:05"
	completeFormat    = "Monday, January 2, 2006 15:04:05"
	descriptionFormat = "Monday, January 2, 2006, at 15:04"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	result, _ := time.Parse(baseFormat, date)

	return result
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	// HasPassed("July 25, 2019 13:45:00")
	toCheck, _ := time.Parse(complexFormat, date)

	return toCheck.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	// "Thursday, May 13, 2010 20:32:00"
	toCheck, _ := time.Parse(completeFormat, date)
	hour := toCheck.Hour()
	return hour >= 12 && hour <= 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	appDate := Schedule(date)

	return fmt.Sprintf("You have an appointment on %s.", appDate.Format(descriptionFormat))
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	thisYear := time.Now().Year()
	return time.Date(thisYear, time.September, 15, 0, 0, 0, 0, time.UTC)
}
