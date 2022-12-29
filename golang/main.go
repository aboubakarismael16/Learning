package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	// Print the current time
	currentTime := time.Now()
	fmt.Println("Current time:", currentTime.Format(time.RFC1123))

	// Create a new blog post
	post := NewPost("My First Go Blog Post", "Hello, world!", currentTime)

	// Print the post's title and content
	fmt.Println("Title:", post.Title)
	fmt.Println("Content:", post.Content)

	// Start the HTTP server
	http.ListenAndServe(":8080", nil)
}

// Post represents a blog post
type Post struct {
	Title   string
	Content string
	Date    time.Time
}

// NewPost creates a new Post instance with the given title and content
func NewPost(title, content string, date time.Time) *Post {
	return &Post{
		Title:   title,
		Content: content,
		Date:    date,
	}
}
