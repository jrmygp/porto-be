package controllers

import (
	"net/http"
	"path/filepath"
	"porto-be/models"
	requests "porto-be/requests/skill"
	"porto-be/responses"
	skillResponse "porto-be/responses/skill"
	services "porto-be/services/skill"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SkillController struct {
	service services.Service
}

func NewSkillController(service services.Service) *SkillController {
	return &SkillController{service}
}

func convertSkillResponse(o models.Skill) skillResponse.SkillResponse {
	return skillResponse.SkillResponse{
		ID:    o.ID,
		Title: o.Title,
		Image: o.Image,
	}
}

// Find All Skill
func (c *SkillController) FindAllSkills(ctx *gin.Context) {
	skills, err := c.service.FindAll()
	if err != nil {
		webResponse := responses.Response{
			Code:   http.StatusBadRequest,
			Status: "ERROR",
			Data:   err,
		}
		ctx.JSON(http.StatusBadRequest, webResponse)
	}

	var skillResponses []skillResponse.SkillResponse

	for _, skill := range skills {
		response := convertSkillResponse(skill)

		skillResponses = append(skillResponses, response)
	}

	webResponse := responses.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   skillResponses,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

// Find Skill By ID
func (c *SkillController) FindSkillByID(ctx *gin.Context) {
	idString := ctx.Param("id")
	id, _ := strconv.Atoi(idString)

	skill, err := c.service.FindByID(id)
	if err != nil {
		webResponse := responses.Response{
			Code:   http.StatusBadRequest,
			Status: "ERROR",
			Data:   err,
		}
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}

	webResponse := responses.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   convertSkillResponse(skill),
	}

	ctx.JSON(http.StatusOK, webResponse)
}

// Create New Skill
func (c *SkillController) CreateNewSkill(ctx *gin.Context) {
	var skillRequest requests.CreateSkillRequest
	err := ctx.ShouldBind(&skillRequest)
	if err != nil {
		webResponse := responses.Response{
			Code:   http.StatusBadRequest,
			Status: "ERROR",
			Data:   err,
		}
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}

	file, err := ctx.FormFile("image")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "File upload failed",
		})
		return
	}

	destination := "public/skill/"
	filePath := filepath.Join(destination, file.Filename)
	if err := ctx.SaveUploadedFile(file, filePath); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to save file",
		})
		return
	}

	skill, err := c.service.Create(skillRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	webResponse := responses.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   convertSkillResponse(skill),
	}

	ctx.JSON(http.StatusOK, webResponse)
}

// Edit Skill
func (c *SkillController) EditSkill(ctx *gin.Context) {
	var skillRequest requests.UpdateSkillRequest

	err := ctx.ShouldBind(&skillRequest)
	if err != nil {
		webResponse := responses.Response{
			Code:   http.StatusBadRequest,
			Status: "ERROR",
			Data:   err,
		}
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}

	_, imageHeader, err := ctx.Request.FormFile("image")
	if err == nil && imageHeader != nil {
		file, err := ctx.FormFile("image")
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "File upload failed",
			})
			return
		}

		destination := "public/skill/"
		filePath := filepath.Join(destination, file.Filename)
		if err := ctx.SaveUploadedFile(file, filePath); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to save file",
			})
			return
		}
	}

	idString := ctx.Param("id")
	id, _ := strconv.Atoi(idString)

	skill, err := c.service.Update(id, skillRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	webResponse := responses.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   convertSkillResponse(skill),
	}

	ctx.JSON(http.StatusOK, webResponse)
}

// Delete Skill
func (c *SkillController) DeleteSkill(ctx *gin.Context) {
	idString := ctx.Param("id")
	id, _ := strconv.Atoi(idString)

	skill, err := c.service.Delete(id)
	if err != nil {
		webResponse := responses.Response{
			Code:   http.StatusBadRequest,
			Status: "ERROR",
			Data:   err,
		}
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}

	webResponse := responses.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   convertSkillResponse(skill),
	}

	ctx.JSON(http.StatusOK, webResponse)
}
