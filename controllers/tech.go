package controllers

import (
	"net/http"
	"path/filepath"
	"porto-be/models"
	requests "porto-be/requests/tech"
	"porto-be/responses"
	techResponse "porto-be/responses/tech"
	services "porto-be/services/tech"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TechController struct {
	service services.Service
}

func NewTechController(service services.Service) *TechController {
	return &TechController{service}
}

func convertTechResponse(o models.Tech) techResponse.TechResponse {
	return techResponse.TechResponse{
		ID:    o.ID,
		Title: o.Title,
		Image: o.Image,
	}
}

// Find All Tech
func (h *TechController) FindAllTechs(c *gin.Context) {
	techs, err := h.service.FindAll()
	if err != nil {
		webResponse := responses.Response{
			Code:   http.StatusBadRequest,
			Status: "ERROR",
			Data:   err,
		}
		c.JSON(http.StatusBadRequest, webResponse)
	}

	var techResponses []techResponse.TechResponse

	for _, tech := range techs {
		response := convertTechResponse(tech)

		techResponses = append(techResponses, response)
	}

	webResponse := responses.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   techResponses,
	}

	c.JSON(http.StatusOK, webResponse)
}

// Find Tech By ID
func (h *TechController) FindTechByID(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	tech, err := h.service.FindByID(id)
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
		Data:   convertTechResponse(tech),
	}

	c.JSON(http.StatusOK, webResponse)
}

// Create New Tech
func (h *TechController) CreateNewTech(c *gin.Context) {
	var techForm requests.CreateTechRequest

	err := c.ShouldBind(&techForm)
	if err != nil {
		webResponse := responses.Response{
			Code:   http.StatusBadRequest,
			Status: "ERROR",
			Data:   err,
		}
		c.JSON(http.StatusBadRequest, webResponse)
		return
	}

	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "File upload failed",
		})
		return
	}

	destination := "public/tech/"
	filePath := filepath.Join(destination, file.Filename)
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to save file",
		})
		return
	}

	tech, err := h.service.Create(techForm)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	webResponse := responses.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   convertTechResponse(tech),
	}

	c.JSON(http.StatusOK, webResponse)
}

// Edit Project
func (h *TechController) EditTech(c *gin.Context) {
	var techForm requests.UpdateTechRequest

	err := c.ShouldBind(&techForm)
	if err != nil {
		webResponse := responses.Response{
			Code:   http.StatusBadRequest,
			Status: "ERROR",
			Data:   err,
		}
		c.JSON(http.StatusBadRequest, webResponse)
		return
	}

	_, imageHeader, err := c.Request.FormFile("image")
	if err == nil && imageHeader != nil {
		file, err := c.FormFile("image")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "File upload failed",
			})
			return
		}

		destination := "public/tech/"
		filePath := filepath.Join(destination, file.Filename)
		if err := c.SaveUploadedFile(file, filePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to save file",
			})
			return
		}
	}

	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)
	tech, err := h.service.Update(id, techForm)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	webResponse := responses.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   convertTechResponse(tech),
	}

	c.JSON(http.StatusOK, webResponse)
}

// Delete Tech
func (h *TechController) DeleteTech(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	tech, err := h.service.Delete(id)
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
		Data:   convertTechResponse(tech),
	}

	c.JSON(http.StatusOK, webResponse)
}
