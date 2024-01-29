package requests

import "mime/multipart"

type UpdateProjectRequest struct {
	Title       string                `form:"title"`
	Description string                `form:"description"`
	Url         *multipart.FileHeader `form:"url"`
}
