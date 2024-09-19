package models

import "time"

type TimeStruct struct {
	Year, Month, Day, Hour, Minutes int
}

type Location struct {
	Room, Building, Address, Postcode, City *string
}

type Event struct {
	ID          int
	Name, Group string
	Start, End  TimeStruct
	Location    *Location
}

type Deadline struct {
	ID          int
	Name, Group string
	Time        TimeStruct
	Location    *Location
}

func GetTime() TimeStruct {
	time := time.Now()
	months := map[string]int{
		"January":   1,
		"February":  2,
		"March":     3,
		"April":     4,
		"May":       5,
		"June":      6,
		"July":      7,
		"August":    8,
		"September": 9,
		"October":   10,
		"November":  11,
		"December":  12,
	}
	return TimeStruct{time.Year(), months[time.Month().String()], time.Day(), time.Hour(), time.Minute()}
}
