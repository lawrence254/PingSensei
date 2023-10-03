package utils

import "time"

func ConvertDurationToMs(input time.Duration) time.Duration{
	result := time.Duration(input) * time.Nanosecond
	return result
}