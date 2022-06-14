package database

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	CreatedAt time.Time `json:"createdAt"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Name      string    `json:"name"`
	Age       int       `json:"age"`
}

// CreateUser We're using email as a primary key, that is, we can't have two users with the same email.
// If you look at the databaseSchema struct you'll notice that the Users field is a map: map[string]User.
// The string key in the map will be the user's email.
// This function should read the current state of the database, create a new user struct, add it to the Users map in the schema, then update the data on disk.
// Don't forget to set the CreatedAt field to time.Now().UTC().
func (c Client) CreateUser(email, password, name string, age int) (User, error) {
	db, err := c.readDB()
	if err != nil {
		return User{}, fmt.Errorf("readDB: %v", err)
	}
	_, ok := db.Users[email]
	if ok {
		return User{}, errors.New("user already exists")
	}

	user := User{
		CreatedAt: time.Now().UTC(),
		Email:     email,
		Password:  password,
		Name:      name,
		Age:       age,
	}
	db.Users[email] = user

	err = c.updateDB(db)
	if err != nil {
		return User{}, fmt.Errorf("updateDB: %v", err)
	}
	return user, err
}

//This function will be similar to CreateUser, but if the user doesn't already exist,
//it should return an error: "user doesn't exist".
//It also won't update the CreatedAt timestamp, that should be left alone.
func (c Client) UpdateUser(email, password, name string, age int) (User, error) {
	db, err := c.readDB()
	if err != nil {
		return User{}, err
	}
	user, ok := db.Users[email]
	if !ok {
		return User{}, fmt.Errorf("user doesn't exist")
	}

	user.Name = name
	user.Password = password
	user.Age = age
	db.Users[email] = user
	err = c.updateDB(db)
	if err != nil {
		return User{}, err
	}
	return user, err
}

// Given the user's email, find the user in the database and return it.
func (c Client) GetUser(email string) (User, error) {
	db, err := c.readDB()
	if err != nil {
		return User{}, err
	}
	user, ok := db.Users[email]
	if !ok {
		return User{}, fmt.Errorf("user doesn't exist")
	}

	return user, nil
}

// Given the user's email, delete the user from the database if it exists.
func (c Client) DeleteUser(email string) error {
	db, err := c.readDB()
	if err != nil {
		return fmt.Errorf("deleteUser: %v", err)
	}
	_, ok := db.Users[email]
	if !ok {
		return fmt.Errorf("user doesn't exist")
	}
	delete(db.Users, email)
	return nil
}
