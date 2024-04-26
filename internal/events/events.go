package events

import (
	"event-management-system/model"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

func AddEvent(users []model.User, userName, title, description string, startTime, endTime, date time.Time) ([]model.User, error) {
	newEvent := model.Event{
		EventID:     generateID(),
		Title:       title,
		Description: description,
		StartTime:   startTime,
		EndTime:     endTime,
		Date:        date,
	}

	found := false
	for i, user := range users {
		if user.Name == userName {
			users[i].Events = append(users[i].Events, newEvent)
			found = true
			break
		}
	}

	if !found {
		newUser := model.User{
			UserID: generateID(),
			Name:   userName,
			Events: []model.Event{newEvent},
		}
		users = append(users, newUser)
	}

	return users, nil
}

func DeleteEvent(users []model.User, userName, eventID string) ([]model.User, error) {
	for i, user := range users {
		if user.Name == userName {
			for j, event := range user.Events {
				if event.EventID == eventID {
					users[i].Events = append(users[i].Events[:j], users[i].Events[j+1:]...)
					return users, nil
				}
			}
		}
	}
	return users, fmt.Errorf("event not found")
}

func SearchEventsByTitle(users []model.User, title string) []model.Event {
	var foundEvents []model.Event
	for _, user := range users {
		for _, event := range user.Events {
			if strings.Contains(strings.ToLower(event.Title), strings.ToLower(title)) {
				foundEvents = append(foundEvents, event)
			}
		}
	}
	return foundEvents
}

func ModifyEvent(users []model.User, userName, eventID string, newEvent model.Event) ([]model.User, error) {
	for i, user := range users {
		if user.Name == userName {
			for j, event := range user.Events {
				if event.EventID == eventID {
					users[i].Events[j] = newEvent
					return users, nil
				}
			}
		}
	}
	return users, fmt.Errorf("event not found")
}

func ListEvents(users []model.User) {
	for _, user := range users {
		fmt.Printf("\nEvents for %s:\n", user.Name)
		for _, event := range user.Events {
			fmt.Printf("ID: %s, Title: %s, Date: %s\n", event.EventID, event.Title, event.Date.Format("2006-01-02"))
		}
	}
}

func generateID() string {
	return uuid.New().String()
}
