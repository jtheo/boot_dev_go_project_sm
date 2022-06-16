package database

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

// Post -
type Post struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UserEmail string    `json:"userEmail"`
	Text      string    `json:"text"`
}

func (c Client) CreatePost(userEmail, text string) (Post, error) {
	db, err := c.readDB()
	if err != nil {
		return Post{}, err
	}

	_, ok := db.Users[userEmail]
	if !ok {
		return Post{}, fmt.Errorf("user doesn't exist")
	}

	post := Post{
		CreatedAt: time.Now().UTC(),
		ID:        uuid.New().String(),
		UserEmail: userEmail,
		Text:      text,
	}

	db.Posts[post.ID] = post

	err = c.updateDB(db)
	if err != nil {
		return Post{}, fmt.Errorf("createPost: %v", err)
	}
	return post, nil
}

func (c Client) GetPosts(userEmail string) ([]Post, error) {
	db, err := c.readDB()
	if err != nil {
		return []Post{}, err
	}

	posts := []Post{}

	for _, p := range db.Posts {
		fmt.Printf("p.UserEmail >%s< | userEmail >%s< | equal %v\n", p.UserEmail, userEmail, p.UserEmail == userEmail)
		if p.UserEmail == userEmail {

			posts = append(posts, p)
		}
	}

	return posts, nil
}

func (c Client) DeletePost(id string) error {
	db, err := c.readDB()
	if err != nil {
		return err
	}

	_, ok := db.Posts[id]
	if !ok {
		return fmt.Errorf("the post id doesn't exist")
	}

	delete(db.Posts, id)
	err = c.updateDB(db)
	return err
}
