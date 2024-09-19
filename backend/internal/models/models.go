package models

type TimeStruct struct {
	year, month, day, hour, minutes int
}

type Location struct {
	room, building, address, postcode, city *string
}

type Event struct {
	ID          int
	name, group string
	start, end  TimeStruct
	location    *Location
}

type Deadline struct {
	ID          int
	name, group string
	time        TimeStruct
	location    *Location
}
