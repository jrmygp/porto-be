package models

import (
	"time"
)

type Project struct {
	ID          int
	Title       string
	Description string
	Url         string
	Image       string
	Stacks      []Skill `gorm:"many2many:project_stacks;"`
	Stack_id    []int   `gorm:"-"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
