package forms

import "mime/multipart"

type ProjectForm struct {
	Title       string                `form:"title" binding:"required"`
	Description string                `form:"description" binding:"required"`
	Url         *multipart.FileHeader `form:"url" binding:"required"`
}
