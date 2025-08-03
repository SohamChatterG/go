package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("welcome to my time tracker app")

	presentTime := time.Now()
	fmt.Println("Current time is:", presentTime)

	fmt.Println(presentTime.Format("01-02-2006 Monday 15:04:05 "))

	createdDate := time.Date(2002, time.December, 23, 0, 0, 0, 0, time.UTC)
	fmt.Println("Manually created date:", createdDate.Format("02 Jan 2006"))

	// User input
	var day, year int
	var monthStr string

	fmt.Println("\nEnter a date:")
	fmt.Print("Day: ")
	fmt.Scanln(&day)

	fmt.Print("Month (e.g. January): ")
	fmt.Scanln(&monthStr)

	fmt.Print("Year: ")
	fmt.Scanln(&year)

	// Convert month string to time.Month
	monthMap := map[string]time.Month{
		"January":   time.January,
		"February":  time.February,
		"March":     time.March,
		"April":     time.April,
		"May":       time.May,
		"June":      time.June,
		"July":      time.July,
		"August":    time.August,
		"September": time.September,
		"October":   time.October,
		"November":  time.November,
		"December":  time.December,
	}

	month, ok := monthMap[monthStr]
	if !ok {
		fmt.Println("Invalid month entered.")
		return
	}

	userDate := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
	fmt.Println("User-entered date:", userDate.Format("02 Jan 2006 Monday"))
}
