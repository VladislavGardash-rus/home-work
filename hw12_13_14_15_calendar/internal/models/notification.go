package models

import "time"

type Notification struct {
	Id     int
	Title  string
	Date   time.Time
	UserId int
}
