package main

import (
	"fmt"

	"cloud.google.com/go/civil"
)

type day struct {
	time civil.Date
}

func main() {
	today := civil.Date{2024, 9, 17}
	fmt.Println("calendar")
	fmt.Println(today)
}
