package utils

import (
	"time"
)

func ParseShortTime(timeStr string) time.Time {
	time_layout := "2006-01-02"
	lt, _ := time.ParseInLocation(time_layout, timeStr, time.Local)
	return lt
}

func ParseLongTime(timeStr string) time.Time {
	time_layout := "2006-01-02  15:04:05"
	lt, _ := time.ParseInLocation(time_layout, timeStr, time.Local)
	return lt
}