package responses

import "porto-be/models"

type ProjectResponse struct {
	ID          int            `json:"id"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Url         string         `json:"url"`
	Image       string         `json:"image"`
	Stacks      []models.Skill `json:"stacks"`
}
