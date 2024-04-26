package storage

import (
	"encoding/json"
	"event-management-system/model"
	"os"
)

const FilePath = "events.json"

func LoadEvents() ([]model.User, error) {
	if _, err := os.Stat(FilePath); os.IsNotExist(err) {
		users := []model.User{}
		if err := saveEventsToFile(users); err != nil {
			return nil, err
		}
	}

	data, err := os.ReadFile(FilePath)
	if err != nil {
		return nil, err
	}

	var users []model.User
	if err := json.Unmarshal(data, &users); err != nil {
		return nil, err
	}

	return users, nil
}

func SaveEvents(users []model.User) error {
	return saveEventsToFile(users)
}

func saveEventsToFile(users []model.User) error {
	data, err := json.Marshal(users)
	if err != nil {
		return err
	}

	return os.WriteFile(FilePath, data, 0644)
}
