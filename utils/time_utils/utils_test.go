package time_utils

import (
	"testing"
	"time"
)

func TestRoundDuration(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		duration time.Duration
		want     string
	}{
		{
			name:     "less than a second",
			duration: 45*time.Millisecond + 52*time.Microsecond,
			want:     "45.052ms",
		},
		{
			name:     "second with ms",
			duration: 1*time.Second + 45*time.Millisecond,
			want:     "1.045s",
		},
		{
			name:     "less than a minute",
			duration: 45 * time.Second,
			want:     "45s",
		},
		{
			name:     "exactly one minute",
			duration: 1 * time.Minute,
			want:     "1m",
		},
		{
			name:     "minutes and seconds",
			duration: 2*time.Minute + 30*time.Second,
			want:     "2m30s",
		},
		{
			name:     "exactly one hour",
			duration: 1 * time.Hour,
			want:     "1h",
		},
		{
			name:     "hours and minutes",
			duration: 2*time.Hour + 45*time.Minute,
			want:     "2h45m",
		},
		{
			name:     "exactly one day",
			duration: 24 * time.Hour,
			want:     "1d",
		},
		{
			name:     "days and hours",
			duration: 2*24*time.Hour + 12*time.Hour,
			want:     "2d12h",
		},
		{
			name:     "multiple days",
			duration: 3*24*time.Hour + 18*time.Hour + 10*time.Minute + 20*time.Second,
			want:     "3d18h",
		},
		{
			name:     "ms",
			duration: 512*time.Millisecond + 512*time.Microsecond + 512*time.Nanosecond,
			want:     "512.512ms",
		},
		{
			name:     "µ",
			duration: 512*time.Microsecond + 51*time.Nanosecond,
			want:     "512.051µ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := RoundDuration(tt.duration); got != tt.want {
				t.Errorf("RoundDuration() = %v, want %v", got, tt.want)
			}
		})
	}
}
