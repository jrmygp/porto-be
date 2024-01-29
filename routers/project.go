package routers

import (
	"net/http"
	"porto-be/controllers"

	"github.com/gin-gonic/gin"
)

func NewRouter(projectController *controllers.ProjectController) *gin.Engine {
	router := gin.Default()

	router.GET("", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Welcome Home")
	})

	router.Static("/public", "./public")
	router.MaxMultipartMemory = 8 << 20 // 8 MiB

	projectRoutes := router.Group("/project")
	projectRoutes.GET("/", projectController.FindAllProjects)
	projectRoutes.GET("/:id", projectController.FindProjectByID)
	projectRoutes.POST("/", projectController.CreateNewProject)
	projectRoutes.PATCH("/:id", projectController.EditProject)
	projectRoutes.DELETE("/:id", projectController.DeleteProject)

	return router

}
