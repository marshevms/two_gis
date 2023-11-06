package model

import "time"

type Order struct {
	ID    int64     `json:"id"`
	Room  string    `json:"room"`
	Email string    `json:"email"`
	From  time.Time `json:"from"`
	To    time.Time `json:"to"`
}
