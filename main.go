package main

import (
	"fmt"
	"porto-be/config"
	"porto-be/controllers"
	"porto-be/models"
	projectRepository "porto-be/repositories/project"
	techRepository "porto-be/repositories/tech"
	projectService "porto-be/services/project"
	techService "porto-be/services/tech"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Database
	db := config.DatabaseConnection()
	db.AutoMigrate(&models.Project{})

	// Repository
	projectRepository := projectRepository.NewRepository(db)
	techRepository := techRepository.NewRepository(db)

	// Service
	projectService := projectService.NewService(projectRepository)
	techService := techService.NewService(techRepository)

	// Controller
	projectController := controllers.NewProjectController(projectService)
	techController := controllers.NewTechController(techService)

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

	techRouter := router.Group("/tech")
	techRouter.GET("", techController.FindAllTechs)
	techRouter.GET("/:id", techController.FindTechByID)
	techRouter.POST("", techController.CreateNewTech)
	techRouter.PATCH("/:id", techController.EditTech)
	techRouter.DELETE("/:id", techController.DeleteTech)

	fmt.Println("jeremy loves andre to the heart")

	router.Run(":8082")
}
