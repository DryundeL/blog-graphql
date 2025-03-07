package models

type NewPost struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}
