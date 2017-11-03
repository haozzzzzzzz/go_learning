package main

import (
	"time"
	"fmt"
	"log"
)

func main() {
	now := time.Now()
	fmt.Println(now)

	strTodayStart := now.Format("2006-01-02 00:00:00 -0700 MST")
	fmt.Println(strTodayStart)

	timeTodayStart, err := time.Parse("2006-01-02 15:04:05 -0700 MST", strTodayStart)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(timeTodayStart)
	fmt.Println(timeTodayStart.Unix())
}
