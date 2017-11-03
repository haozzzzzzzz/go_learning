package main

import (
	"time"
	"fmt"
)

func DayStartTime(t time.Time) (dayStartTime time.Time, err error) {
	strDayStartTime := t.Format("2006-01-02 00:00:00 -0700 MST")
	dayStartTime, err = time.Parse("2006-01-02 15:04:05 -0700 MST", strDayStartTime)
	return
}

func WeekStartTime(t time.Time) (weekStartTime time.Time, err error){
	dayStart, err := DayStartTime(t)
	if err != nil {
		return
	}

	dayStartUnix := dayStart.Unix()
	weekDay := dayStart.Weekday()

	var secondsPerDay int64 = 60*60*24
	offset := int64(int(weekDay) - 1) * secondsPerDay

	weekStartUnix := dayStartUnix - offset
	weekStartTime = time.Unix(weekStartUnix, 0)
	return
}

func main() {
	now := time.Now()
	dayStartTime, _ := DayStartTime(now)
	weekStartTime, _ := WeekStartTime(now)

	fmt.Println(dayStartTime)
	fmt.Println(weekStartTime)

	ident := weekStartTime.Format("w_2006_01_02")
	fmt.Println(ident)
}
