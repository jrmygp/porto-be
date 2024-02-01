package requests

import "mime/multipart"

type UpdateSkillRequest struct {
	Title string                `form:"title"`
	Image *multipart.FileHeader `form:"image"`
}
