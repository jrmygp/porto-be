package requests

import "mime/multipart"

type CreateTechRequest struct {
	Title string                `form:"title" binding:"required"`
	Image *multipart.FileHeader `form:"image" binding:"required"`
}
