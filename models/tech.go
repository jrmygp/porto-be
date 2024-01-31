package models

import "time"

type Tech struct {
	ID         int
	Title      string
	Image      string
	Percentage int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
