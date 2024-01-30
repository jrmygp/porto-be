package requests

import "mime/multipart"

type UpdateProjectRequest struct {
	Title       string                `form:"title"`
	Description string                `form:"description"`
	Url         string                `form:"url"`
	Image       *multipart.FileHeader `form:"image"`
}
