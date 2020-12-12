package data

import "time"

type Thought struct {
	ItemID string    `json:"itemId"`
	Author string    `json:"author"`
	Date   time.Time `json:"date"`
	Image  string    `json:"image"`
	Text   string    `json:"text"`
	Title  string    `json:"title"`
}

type ThoughtsResponse struct {
	Status  int      `json:"status"`
	Message string   `json:"message"`
	Data    []Thought `json:"data"`
}

