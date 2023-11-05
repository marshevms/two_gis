package model

import "time"

type Order struct {
	Room  string
	Email string
	From  time.Time
	To    time.Time
}
