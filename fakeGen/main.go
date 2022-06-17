package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/jaswdr/faker"
)

func main() {
	u := flag.Int("n", 1, "number of fake identities")
	p := flag.Int("p", 1, "posts per user")
	addr := flag.String("addr", "localhost", "host to send the requests")
	port := flag.Int("port", 8080, "port to send the requests")
	flag.Parse()

	fullAddr := fmt.Sprintf("http://%s:%d", *addr, *port)
	f := faker.New()
	for i := 0; i < *u; i++ {
		user := genUser(f)
		send(fullAddr, "users", user)

		posts := genPost(f, user.Email, *p)
		for _, post := range posts {
			send(fullAddr, "posts", post)
		}
	}
}

func genUser(f faker.Faker) user {
	rand.Seed(time.Now().UnixNano())
	min := 18
	max := 100
	r1 := rand.Intn(max-min+1) + min

	email := f.Person().Contact().Email
	return user{
		Name:     f.Person().Name(),
		Age:      r1,
		Password: f.Internet().Password(),
		Email:    email,
	}
}

func genPost(f faker.Faker, email string, num int) []post {
	posts := []post{}
	for i := 0; i < num; i++ {
		quote := f.Person().Faker.Lorem().Sentence(15)
		posts = append(posts, post{
			UserEmail: email,
			Text:      quote,
		})
	}
	return posts
}

type user struct {
	Name     string
	Age      int
	Password string
	Email    string
}

type post struct {
	UserEmail string
	Text      string
}

func send(addr, e string, f interface{}) {
	endp := fmt.Sprintf("%s/%s/", addr, e)

	payload, err := json.Marshal(f)
	if err != nil {
		log.Fatalf("da fuq? %v => %v", f, err)
	}
	responseBody := bytes.NewBuffer(payload)
	resp, err := http.Post(endp, "application/json", responseBody)
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	sb := string(body)
	log.Printf("[%s]: %s,", e, sb)
}
