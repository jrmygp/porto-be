package controllers

import (
	"fmt"
	"net/http"
	"path/filepath"
	"porto-be/forms"
	"porto-be/models"
	"porto-be/responses"
	"porto-be/services"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type controller struct {
	service services.Service
}

func NewController(service services.Service) *controller {
	return &controller{service}
}

// Private function
func convertResponse(o models.Project) responses.ProjectResponse {
	return responses.ProjectResponse{
		ID:          o.ID,
		Title:       o.Title,
		Description: o.Description,
		Url:         o.Url,
	}
}

func (h *controller) FindAllProjects(c *gin.Context) {
	projects, err := h.service.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	var projectResponses []responses.ProjectResponse

	for _, project := range projects {
		response := convertResponse(project)

		projectResponses = append(projectResponses, response)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": projectResponses,
	})
}

func (h *controller) FindProjectByID(c *gin.Context) {
	idString := c.Param("id")
	// convert id from string to int
	id, _ := strconv.Atoi(idString)

	project, err := h.service.FindByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	convertedProject := convertResponse(project)

	c.JSON(http.StatusOK, gin.H{
		"data": convertedProject,
	})
}

func (h *controller) CreateNewProject(c *gin.Context) {
	var projectForm forms.ProjectForm

	err := c.ShouldBind(&projectForm)
	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Handle file upload
	file, err := c.FormFile("url")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "File upload failed",
		})
		return
	}

	// Save the file to the server
	destination := "public/project/"
	filePath := filepath.Join(destination, file.Filename)
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to save file",
		})
		return
	}

	project, err := h.service.Create(projectForm)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": convertResponse(project),
	})
}

func (h *controller) EditProject(c *gin.Context) {
	var projectForm forms.ProjectForm

	err := c.ShouldBindJSON(&projectForm)
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			message := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, message)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}

	idString := c.Param("id")
	// convert id from string to int
	id, _ := strconv.Atoi(idString)
	project, err := h.service.Update(id, projectForm)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": convertResponse(project),
	})
}

func (h *controller) DeleteProject(c *gin.Context) {
	id, _ := strconv.Atoi("id")
	project, err := h.service.Delete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": convertResponse(project),
	})
}
