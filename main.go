package main

import (
	"fmt"
	"net/http"
	"porto-be/config"
	"porto-be/controllers"
	"porto-be/models"
	projectRepository "porto-be/repositories/project"
	"porto-be/routers"
	projectService "porto-be/services/project"
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
	routes := routers.NewRouter(projectController)

	server := &http.Server{
		Addr:    ":8080",
		Handler: routes,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

	fmt.Println("jeremy loves andre to the heart")
}
