package routers

import (
	"net/http"
	"porto-be/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewRouter(projectController *controllers.ProjectController) *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())

	router.GET("", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Welcome Home")
	})

	projectRoutes := router.Group("/project")

	projectRoutes.GET("/", projectController.FindAllProjects)
	projectRoutes.GET("/:id", projectController.FindProjectByID)
	projectRoutes.POST("/", projectController.CreateNewProject)
	projectRoutes.PATCH("/:id", projectController.EditProject)
	projectRoutes.DELETE("/:id", projectController.DeleteProject)

	return router

}
