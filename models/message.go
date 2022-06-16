package models

import "time"

type Message struct {
	Id       string
	Time     time.Time
	Body     string
	Username string
}
