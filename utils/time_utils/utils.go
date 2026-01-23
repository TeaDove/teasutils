package time_utils

import (
	"fmt"
	"time"
)

// RoundDuration
// Returns pretty rounded duration string, i.e. 1h, 2d3h etc.
func RoundDuration(d time.Duration) string {
	const (
		thousand        = 1000
		secondsInMinute = 60.0
		minutesInHour   = 60.0
		hoursInDay      = 24.0
	)

	if d.Microseconds() < thousand {
		if int(d.Nanoseconds())%thousand == 0 {
			return fmt.Sprintf("%dms", int(d.Microseconds()))
		}

		return fmt.Sprintf("%d.%03dµ", int(d.Microseconds()), int(d.Nanoseconds())%thousand)
	}

	if d.Milliseconds() < thousand {
		if int(d.Microseconds())%thousand == 0 {
			return fmt.Sprintf("%dms", int(d.Milliseconds()))
		}

		return fmt.Sprintf("%d.%03dms", int(d.Milliseconds()), int(d.Microseconds())%thousand)
	}

	if d.Seconds() < secondsInMinute {
		return d.String()
	}

	if d.Minutes() < minutesInHour {
		if int(d.Seconds())%secondsInMinute == 0 {
			return fmt.Sprintf("%dm", int(d.Minutes()))
		}

		return fmt.Sprintf("%dm%ds", int(d.Minutes()), int(d.Seconds())%secondsInMinute)
	}

	if d.Hours() < hoursInDay {
		if int(d.Minutes())%minutesInHour == 0 {
			return fmt.Sprintf("%dh", int(d.Hours()))
		}

		return fmt.Sprintf("%dh%dm", int(d.Hours()), int(d.Minutes())%minutesInHour)
	}

	if int(d.Hours())%hoursInDay == 0 {
		return fmt.Sprintf("%dd", int(d.Hours())/hoursInDay)
	}

	return fmt.Sprintf("%dd%dh", int(d.Hours())/hoursInDay, int(d.Hours())%hoursInDay)
}
