package models

import "time"

type Event struct {
	ID                   int
	Title                string
	DateTimeStart        time.Time
	DateTimeEnd          time.Time
	Description          string
	UserId               int
	NotificationDuration *time.Duration
}
