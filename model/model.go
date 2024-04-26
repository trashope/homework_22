package model

import "time"

type User struct {
	UserID string  `json:"userID"`
	Name   string  `json:"name"`
	Events []Event `json:"events"`
}

type Event struct {
	EventID     string    `json:"eventID"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	StartTime   time.Time `json:"startTime"`
	EndTime     time.Time `json:"endTime"`
	Date        time.Time `json:"date"`
}
