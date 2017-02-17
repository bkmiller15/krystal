package helper

import (
	"time"
)

func Get_Current_Date() (year, month, day string) {
	time := time.Now()
	year = time.Format("2006")
	month = time.Format("01")
	day = time.Format("02")

	return year, month, day
}
