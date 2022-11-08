package models

import "time"

type Book struct {
	BookID        uint      `gorm:"primaryKey" json:"id"`
	Book_title    string    `json:"title"`
	Book_author   string    `json:"author"`
	Book_category string    `json:"category"`
	Book_image    string    `json:"image"`
	Created_at    time.Time `json:"created_at"`
	Updated_at    time.Time `json:"updated_at"`
}

type Books []Book
