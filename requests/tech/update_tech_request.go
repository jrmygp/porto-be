package requests

import "mime/multipart"

type UpdateTechRequest struct {
	Title      string                `form:"title"`
	Image      *multipart.FileHeader `form:"image"`
	Percentage int                   `form:"percentage"`
}
