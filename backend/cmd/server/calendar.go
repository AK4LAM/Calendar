package main

import (
	"fmt"
	"time"

	"github.com/AK4LAM/calendar/backend/internal/models"
)

func main() {
	newDeadline := models.Deadline{ID: 0, Name: "first deadline", Group: "0", Time: models.GetTime()}

	time := time.Now()
	fmt.Println("calendar")
	fmt.Println(time)
	fmt.Println(newDeadline)
}
