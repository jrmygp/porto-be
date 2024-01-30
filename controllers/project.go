package controllers

import (
	"fmt"
	"net/http"
	"path/filepath"
	"porto-be/models"
	requests "porto-be/requests/project"
	"porto-be/responses"
	projectResponse "porto-be/responses/project"
	services "porto-be/services/project"

	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ProjectController struct {
	service services.Service
}

func NewController(service services.Service) *ProjectController {
	return &ProjectController{service}
}

// Private function
func convertResponse(o models.Project) projectResponse.ProjectResponse {
	return projectResponse.ProjectResponse{
		ID:          o.ID,
		Title:       o.Title,
		Description: o.Description,
		Url:         o.Url,
		Image:       o.Image,
	}
}

// Find All Project
func (h *ProjectController) FindAllProjects(c *gin.Context) {
	projects, err := h.service.FindAll()
	if err != nil {
		webResponse := responses.Response{
			Code:   http.StatusBadRequest,
			Status: "ERROR",
			Data:   err,
		}
		c.JSON(http.StatusBadRequest, webResponse)
		return
	}

	var projectResponses []projectResponse.ProjectResponse

	for _, project := range projects {
		response := convertResponse(project)

		projectResponses = append(projectResponses, response)
	}

	webResponse := responses.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   projectResponses,
	}

	c.JSON(http.StatusOK, webResponse)
}

// Find Project By ID
func (h *ProjectController) FindProjectByID(c *gin.Context) {
	idString := c.Param("id")
	// convert id from string to int
	id, _ := strconv.Atoi(idString)

	project, err := h.service.FindByID(id)
	if err != nil {
		webResponse := responses.Response{
			Code:   http.StatusBadRequest,
			Status: "ERROR",
			Data:   err,
		}
		c.JSON(http.StatusBadRequest, webResponse)
		return
	}

	webResponse := responses.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   convertResponse(project),
	}

	c.JSON(http.StatusOK, webResponse)
}

// Create New Project
func (h *ProjectController) CreateNewProject(c *gin.Context) {
	var projectForm requests.CreateProjectRequest

	err := c.ShouldBind(&projectForm)
	if err != nil {
		webResponse := responses.Response{
			Code:   http.StatusBadRequest,
			Status: "ERROR",
			Data:   err,
		}
		c.JSON(http.StatusBadRequest, webResponse)
		return
	}

	// Handle file upload
	file, err := c.FormFile("image")
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

	webResponse := responses.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   convertResponse(project),
	}

	c.JSON(http.StatusOK, webResponse)
}

// Edit Project
func (h *ProjectController) EditProject(c *gin.Context) {
	var projectForm requests.CreateProjectRequest

	err := c.ShouldBindJSON(&projectForm)
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			message := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, message)
		}

		webResponse := responses.Response{
			Code:   http.StatusBadRequest,
			Status: "ERROR",
			Data:   errorMessages,
		}
		c.JSON(http.StatusBadRequest, webResponse)
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

	webResponse := responses.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   convertResponse(project),
	}

	c.JSON(http.StatusOK, webResponse)
}

// Delete Project
func (h *ProjectController) DeleteProject(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	project, err := h.service.Delete(id)
	if err != nil {
		webResponse := responses.Response{
			Code:   http.StatusBadRequest,
			Status: "ERROR",
			Data:   err,
		}
		c.JSON(http.StatusBadRequest, webResponse)
		return
	}

	webResponse := responses.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   convertResponse(project),
	}

	c.JSON(http.StatusOK, webResponse)
}
