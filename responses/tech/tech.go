package responses

type TechResponse struct {
	ID         int    `json:"id"`
	Title      string `json:"title"`
	Image      string `json:"image"`
	Percentage int    `json:"percentage"`
}
