package database

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Client struct {
	path string
}

func NewClient(path string) Client {
	return Client{path: path}
}

type databaseSchema struct {
	Users map[string]User `json:"users"`
	Posts map[string]Post `json:"posts"`
}

// Post -
type Post struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UserEmail string    `json:"userEmail"`
	Text      string    `json:"text"`
}

func (c Client) createDB() error {
	d := databaseSchema{}

	payload, err := json.Marshal(d)
	if err != nil {
		return fmt.Errorf("error marshaling the new DB %v", err)
	}

	err = os.WriteFile(c.path, payload, 0600)
	if err != nil {
		return fmt.Errorf("problem opening the file %s: %v", c.path, err)
	}

	return nil
}

func (c Client) EnsureDB() error {
	_, err := os.ReadFile(c.path)
	if err != nil {
		err = c.createDB()
		if err != nil {
			return fmt.Errorf("error opening %s", c.path)
		}
	}
	return nil
}

func (c Client) updateDB(db databaseSchema) error {
	payload, err := json.Marshal(db)
	if err != nil {
		return fmt.Errorf("error marshaling the new DB %v", err)
	}

	err = os.WriteFile(c.path, payload, 0600)
	if err != nil {
		return fmt.Errorf("problem opening the file %s: %v", c.path, err)
	}
	return nil
}

func (c Client) readDB() (databaseSchema, error) {
	payload, err := os.ReadFile(c.path)
	if err != nil {
		return databaseSchema{}, fmt.Errorf("error opening %s", c.path)
	}

	db := databaseSchema{}

	err = json.Unmarshal(payload, &db)

	return db, fmt.Errorf("unmarshalDB: %v", err)
}
