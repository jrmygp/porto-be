package responses

type ProjectResponse struct {
	ID          int                    `json:"id"`
	Title       string                 `json:"title"`
	Description string                 `json:"description"`
	Url         string                 `json:"url"`
	Image       string                 `json:"image"`
	Stacks      []ProjectStackResponse `json:"stacks"`
}

type ProjectStackResponse struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Image string `json:"image"`
}
