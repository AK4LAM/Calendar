package main

import (
	"fmt"
	"time"
)

type TimeStruct struct {
	year   int
	month  int
	day    int
	hour   int
	minute int
}

type Location struct {
	name     *string
	postcode *string
	city     *string
}

type Event struct {
	ID       int
	name     string
	group    string
	start    TimeStruct
	end      TimeStruct
	location Location
}

type Deadline struct {
	ID    int
	name  string
	group string
	time  TimeStruct
}

func main() {
	time := time.Now()
	fmt.Println("calendar")
	fmt.Println(time)
}
