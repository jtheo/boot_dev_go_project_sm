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

// CreateUser -
func (c Client) CreateUser(email, password, name string, age int) (User, error) {
	if err := userIsEligible(email, password, age); err != nil {
		return User{}, err
	}
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
	err = c.updateDB(db)
	if err != nil {
		return err
	}
	return nil
}

func userIsEligible(email, password string, age int) error {
	if email == "" {
		return fmt.Errorf("email can't be empty")
	}

	if password == "" {
		return fmt.Errorf("password can't be empty")
	}

	if age < 18 {
		return fmt.Errorf("age must be at least %d years old", 18)
	}
	return nil
}
