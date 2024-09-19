package main

import (
	"fmt"
	"time"

	"github.com/AK4LAM/calendar/backend/internal/models"
)

func main() {
	newDeadline := models.Deadline{ID: 0, name: "first deadline", group: 0,
		models.TimeStruct{}}

	time := time.Now()
	fmt.Println("calendar")
	fmt.Println(time)
	fmt.Println(newDeadline)
}
