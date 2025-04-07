package time_utils

import (
	"fmt"
	"time"
)

func RoundDuration(d time.Duration) string {
	const (
		secondsInMinute = 60.0
		minutesInHour   = 60.0
		hoursInDay      = 24.0
	)

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
