package main

import (
	"fmt"
	"porto-be/config"
	"porto-be/controllers"
	"porto-be/models"
	projectRepository "porto-be/repositories/project"
	projectService "porto-be/services/project"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Database
	db := config.DatabaseConnection()
	db.AutoMigrate(&models.Project{})

	// Repository
	projectRepository := projectRepository.NewRepository(db)

	// Service
	projectService := projectService.NewService(projectRepository)

	// Controller
	projectController := controllers.NewController(projectService)

	// Router
	router := gin.Default()
	router.Use(cors.Default())

	router.Static("/public", "./public")
	router.MaxMultipartMemory = 8 << 20 // 8 MiB

	projectRouter := router.Group("/project")
	projectRouter.GET("", projectController.FindAllProjects)
	projectRouter.GET("/:id", projectController.FindProjectByID)
	projectRouter.POST("", projectController.CreateNewProject)
	projectRouter.PATCH("/:id", projectController.EditProject)
	projectRouter.DELETE("/:id", projectController.DeleteProject)

	fmt.Println("jeremy loves andre to the heart")

	router.Run(":8082")
}
