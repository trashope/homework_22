package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	customDate := time.Date(2023, time.November, 4, 0, 0, 0, 0, time.UTC)
	diff := now.Sub(customDate)

	days := int(diff.Hours() / 24)
	hours := int(diff.Hours()) % 24
	minutes := int(diff.Minutes()) % 60
	seconds := int(diff.Seconds()) % 60

	fmt.Printf("Farq: %d kun, %d soat, %d daqiqa, %d soniya\n", days, hours, minutes, seconds)
}
