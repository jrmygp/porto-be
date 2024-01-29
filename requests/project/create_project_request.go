package requests

import "mime/multipart"

type CreateProjectRequest struct {
	Title       string                `form:"title" binding:"required"`
	Description string                `form:"description" binding:"required"`
	Url         *multipart.FileHeader `form:"url" binding:"required"`
}
