package main

import (
	"fmt"
	"time"
)

/*
 * time.go demonstrates how to work with time in go
 */
func main() {
	/* Now() returns the current local time
	in the format: 2019-10-30 17:29:02.6455037 -0500 CDT m=+0.004001001 */
	now := time.Now()
	fmt.Println(now)

	// parse time
	date := "November 28, 2019 5:00 PM CST"
	dateFormat := "January 2, 2006 3:04 PM MST"
	t, _ := time.Parse(dateFormat, date)
	fmt.Println(t)

	// format time using 'RFC822' predefined format
	formattedTime := t.Format(time.RFC822)
	fmt.Println(formattedTime)

	// Add 1 year to the current time
	later := now.AddDate(1, 0, 0)
	fmt.Println(later)

	// Subtract 1 month from current date
	past := now.AddDate(0, -1, 0)
	fmt.Println(past)

	// Add 1 day to current time
	later = now.AddDate(0, 0, 1)
	fmt.Println(later)

	// Add 5 hours to the current time
	later = now.Add(5 * time.Hour)
	fmt.Println(later)

	// Subtract 30 minutes to the current time
	past = now.Add(-30 * time.Minute)
	fmt.Println(past)

	// Add 30 seconds to the current time
	later = now.Add(30 * time.Second)
	fmt.Println(later)

	// Print different parts of the date
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	fmt.Println(now.Zone())

	// wait 2 seconds
	time.Sleep(2 * time.Second)

	/* calculate the duration
	Since() returns the time elapsed*/
	duration := time.Since(now)
	fmt.Println(duration)

	/* The unit of time for duration
	is nanoseconds by default */
	fmt.Println(int64(duration)) // nanoseconds by default

	// convert duration to a different unit of time
	fmt.Println(duration.Microseconds()) // get duration in microseconds
	fmt.Println(duration.Milliseconds()) // get duration in milliseconds
}
