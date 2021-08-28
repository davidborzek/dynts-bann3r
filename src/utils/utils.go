package utils

import (
	"strconv"
	"time"
)

func GetHours() string {
	hours := time.Now().Local().Hour()

	if hours < 10 {
		return "0" + strconv.Itoa(hours)
	}

	return strconv.Itoa(hours)
}

func GetMinutes() string {
	minutes := time.Now().Local().Minute()

	if minutes < 10 {
		return "0" + strconv.Itoa(minutes)
	}

	return strconv.Itoa(minutes)
}

func GetSeconds() string {
	seconds := time.Now().Local().Second()

	if seconds < 10 {
		return "0" + strconv.Itoa(seconds)
	}

	return strconv.Itoa(seconds)
}
