package models

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
