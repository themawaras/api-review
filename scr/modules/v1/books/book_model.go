package books

import "time"

type Book struct {
	Book_id       uint      `gorm:"primaryKey" json:"id"`
	Book_title    string    `json:"name"`
	Book_author   string    `json:"author"`
	Book_category string    `json:"category"`
	Book_image    string    `json:"image"`
	Created_at    time.Time `json:"created_at"`
	Updated_at    time.Time `json:"updated_at"`
}

type Books []Book
