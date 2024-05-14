package util

import (
	"mini_project/constant"
	"time"
)

func StringToDate(date string) (time.Time, error) {
	newDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		return time.Time{}, constant.ErrInputDate
	}
	return newDate, nil
}
