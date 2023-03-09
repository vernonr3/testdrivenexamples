package clockface_test

import (
	"math"
	"time"
)

func secondsInRadians(mtime time.Time) float64 {
	var seconds = float64(mtime.Second())
	return seconds / 60 * 2 * math.Pi
}

func minutesInRadians(mtime time.Time) float64 {
	var minutes = secondsInRadians(mtime)/60 + math.Pi/(30/float64(mtime.Minute()))
	return minutes
}
