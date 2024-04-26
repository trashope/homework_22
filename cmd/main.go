package main

import (
	"bufio"
	"event-management-system/internal/events"
	"event-management-system/internal/storage"
	"event-management-system/model"
	"fmt"
	"os"
	"time"
)

const TimeFormat = "2006-01-02 15:04"

func main() {
	users, err := storage.LoadEvents()
	if err != nil {
		fmt.Printf("Error loading events: %v\n", err)
		return
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("\nEvent Management System")
		fmt.Println("1. Add Event")
		fmt.Println("2. Delete Event")
		fmt.Println("3. Modify Event")
		fmt.Println("4. List Events")
		fmt.Println("5. Exit")
		fmt.Print("Enter your choice: ")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			users = addEventCLI(users, scanner)
		case "2":
			users = deleteEventCLI(users, scanner)
		case "3":
			users = modifyEventCLI(users, scanner)
		case "4":
			events.ListEvents(users)
		case "5":
			fmt.Println("Exiting...")
			err := storage.SaveEvents(users)
			if err != nil {
				fmt.Printf("Error saving events: %v\n", err)
			}
			return
		default:
			fmt.Println("Invalid choice, please enter a number between 1 and 5.")
		}
	}
}

func addEventCLI(users []model.User, scanner *bufio.Scanner) []model.User {
	fmt.Print("Enter user name: ")
	scanner.Scan()
	userName := scanner.Text()

	fmt.Print("Enter event title: ")
	scanner.Scan()
	title := scanner.Text()

	fmt.Print("Enter event description: ")
	scanner.Scan()
	description := scanner.Text()

	fmt.Print("Enter start time (YYYY-MM-DD HH:MM): ")
	scanner.Scan()
	startTimeStr := scanner.Text()
	startTime, _ := time.Parse(TimeFormat, startTimeStr)

	fmt.Print("Enter end time (YYYY-MM-DD HH:MM): ")
	scanner.Scan()
	endTimeStr := scanner.Text()
	endTime, _ := time.Parse(TimeFormat, endTimeStr)

	fmt.Print("Enter date (YYYY-MM-DD): ")
	scanner.Scan()
	dateStr := scanner.Text()
	date, _ := time.Parse("2006-01-02", dateStr)

	users, err := events.AddEvent(users, userName, title, description, startTime, endTime, date)
	if err != nil {
		fmt.Printf("Error adding event: %v\n", err)
	}
	return users
}

func deleteEventCLI(users []model.User, scanner *bufio.Scanner) []model.User {
	fmt.Print("Enter user name: ")
	scanner.Scan()
	userName := scanner.Text()

	fmt.Print("Enter event ID: ")
	scanner.Scan()
	eventID := scanner.Text()

	users, err := events.DeleteEvent(users, userName, eventID)
	if err != nil {
		fmt.Printf("Error deleting event: %v\n", err)
	}
	return users
}

func modifyEventCLI(users []model.User, scanner *bufio.Scanner) []model.User {
	fmt.Print("Enter user name: ")
	scanner.Scan()
	userName := scanner.Text()

	fmt.Print("Enter event ID to modify: ")
	scanner.Scan()
	eventID := scanner.Text()

	existingEvent, found := findEventByID(users, userName, eventID)
	if !found {
		fmt.Println("Event not found.")
		return users
	}

	modifyEventDetails(scanner, &existingEvent)

	users, err := events.ModifyEvent(users, userName, eventID, existingEvent)
	if err != nil {
		fmt.Printf("Error modifying event: %v\n", err)
	} else {
		fmt.Println("Event modified successfully.")
	}

	return users
}

func findEventByID(users []model.User, userName, eventID string) (model.Event, bool) {
	var existingEvent model.Event
	found := false
	for _, user := range users {
		if user.Name == userName {
			for _, event := range user.Events {
				if event.EventID == eventID {
					existingEvent = event
					found = true
					break
				}
			}
		}
		if found {
			break
		}
	}
	return existingEvent, found
}

func modifyEventDetails(scanner *bufio.Scanner, event *model.Event) {
	fmt.Print("Enter new event title (leave blank to keep current): ")
	scanner.Scan()
	title := scanner.Text()
	if title != "" {
		event.Title = title
	}

	fmt.Print("Enter new event description (leave blank to keep current): ")
	scanner.Scan()
	description := scanner.Text()
	if description != "" {
		event.Description = description
	}

	fmt.Print("Enter new start time (YYYY-MM-DD HH:MM) or leave blank to keep current: ")
	scanner.Scan()
	startTimeStr := scanner.Text()
	if startTimeStr != "" {
		startTime, err := time.Parse(TimeFormat, startTimeStr)
		if err == nil {
			event.StartTime = startTime
		}
	}

	fmt.Print("Enter new end time (YYYY-MM-DD HH:MM) or leave blank to keep current: ")
	scanner.Scan()
	endTimeStr := scanner.Text()
	if endTimeStr != "" {
		endTime, err := time.Parse(TimeFormat, endTimeStr)
		if err == nil {
			event.EndTime = endTime
		}
	}

	fmt.Print("Enter new date (YYYY-MM-DD) or leave blank to keep current: ")
	scanner.Scan()
	dateStr := scanner.Text()
	if dateStr != "" {
		date, err := time.Parse("2006-01-02", dateStr)
		if err == nil {
			event.Date = date
		}
	}
}
