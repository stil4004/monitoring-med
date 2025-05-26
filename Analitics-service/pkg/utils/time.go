package utils

import "time"

func GetMoscowTime() time.Time {
	loc, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		return time.Now().UTC().Add(time.Hour * 3)
	}
	return time.Now().In(loc)
}
