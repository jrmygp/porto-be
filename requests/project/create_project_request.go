package requests

import "mime/multipart"

type CreateProjectRequest struct {
	Title       string                `form:"title" binding:"required"`
	Description string                `form:"description" binding:"required"`
	Url         string                `form:"url"`
	Image       *multipart.FileHeader `form:"image" binding:"required"`
}
