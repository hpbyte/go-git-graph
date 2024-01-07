package main

import (
	"time"
)

func startOfDay(d time.Time) time.Time {
	year, month, day := d.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, d.Location())
}
