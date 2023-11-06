package model

import "time"

type Order struct {
	ID    int64
	Room  string
	Email string
	From  time.Time
	To    time.Time
}
