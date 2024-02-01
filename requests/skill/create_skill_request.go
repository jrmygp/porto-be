package requests

import "mime/multipart"

type CreateSkillRequest struct {
	Title string                `form:"title" binding:"required"`
	Image *multipart.FileHeader `form:"image" binding:"required"`
}
