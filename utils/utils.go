package utils

import "time"

func DaysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}
