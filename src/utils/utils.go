package utils

import (
	"strconv"
	"time"

	"github.com/multiplay/go-ts3"
)

func GetHours(client *ts3.Client) (string, error) {
	hours := time.Now().Local().Hour()

	if hours < 10 {
		return "0" + strconv.Itoa(hours), nil
	}

	return strconv.Itoa(hours), nil
}

func GetMinutes(client *ts3.Client) (string, error) {
	minutes := time.Now().Local().Minute()

	if minutes < 10 {
		return "0" + strconv.Itoa(minutes), nil
	}

	return strconv.Itoa(minutes), nil
}

func GetSeconds(client *ts3.Client) (string, error) {
	seconds := time.Now().Local().Second()

	if seconds < 10 {
		return "0" + strconv.Itoa(seconds), nil
	}

	return strconv.Itoa(seconds), nil
}
