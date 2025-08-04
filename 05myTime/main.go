package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("welcome to go time tracker app")

	presentTime := time.Now()
	fmt.Println("Current time is:", presentTime)
	// the way we give the string in the format the result will be such :- "01/02/06 Monday 15:04:05", etc.
	fmt.Println(presentTime.Format("01-02-2006 Monday 15:04:05"))

	createdDate := time.Date(2002, time.December, 23, 0, 0, 0, 0, time.UTC)
	fmt.Println("Manually created date:", createdDate.Format("02 Janu 2006"))

	// converting a string into date
	layoutStr := "06-01-02"
	dateStr := "25-05-02"
	formatted, _ := time.Parse(layoutStr, dateStr)
	fmt.Println("converted from string to date : ", formatted)
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
