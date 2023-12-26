package api

import (
	"fmt"
	"time"
)

type Day struct {
	Date string
	DayOfWeek string
	Info string
}

type Week []Day

var daysOfWeek = []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday"}

// GenerateMockData generates mock data for the given date and number of weeks


func GenerateMockData(startingDate string, numWeeks int) []Week {
	formattedDate, err := time.Parse("2006-01-02", startingDate)
	if err != nil {
		fmt.Printf("Error parsing date: %s", err.Error())
	}
	var weeks []Week
	for w := 0; w < numWeeks; w++ {
		var week Week
		for d := 0; d < 7; d++ {
			date := formattedDate.Add(time.Duration(w*7+d) * 24 * time.Hour)
			week = append(week, Day{
				Date: date.Format("02 Jan"),
				DayOfWeek: daysOfWeek[d % 5],
				Info: "info for" + date.Format("02 Jan"),
			})
		}
		weeks = append(weeks, week)
	}
	return weeks
}
