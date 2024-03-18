package models

import "time"

type Notification struct {
	Id     int       `json:"id"`
	Title  string    `json:"title"`
	Date   time.Time `json:"date"`
	UserId int       `json:"userId"`
}
