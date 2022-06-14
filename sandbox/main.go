package main

import (
	"fmt"
	"log"

	"github.com/jtheo/socialmedia/internal/database"
)

func main() {
	log.Println("Create DB")
	c := database.NewClient("db.json")
	err := c.EnsureDB()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Create User")
	user, err := c.CreateUser("test@example.com", "password", "john doe", 18)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("user created", user)

	log.Println("update User")
	updatedUser, err := c.UpdateUser("test@example.com", "password", "john doe", 18)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("user updated", updatedUser)

	log.Println("get User")
	gotUser, err := c.GetUser("test@example.com")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("user got", gotUser)

	log.Println("delete User")
	err = c.DeleteUser("test@example.com")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("user deleted")

	log.Println("Get User")
	_, err = c.GetUser("test@example.com")
	if err == nil {
		log.Fatal("shouldn't be able to get user that was deleted")
	}
	fmt.Println("user confirmed deleted")
}
