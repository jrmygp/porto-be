package models

import "time"

type Project struct {
	ID          int
	Title       string
	Description string
	Url         string
	Image       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
