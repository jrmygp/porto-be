package main

import (
	"fmt"
	"os"
	"porto-be/controllers"
	"porto-be/models"
	"porto-be/repositories"
	"porto-be/services"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.Project{})

	projectRepository := repositories.NewRepository(db)
	projectService := services.NewService(projectRepository)
	projectController := controllers.NewController(projectService)

	router := gin.Default()
	projectRoutes := router.Group("/project")

	projectRoutes.GET("/", projectController.FindAllProjects)
	projectRoutes.GET("/:id", projectController.FindProjectByID)
	projectRoutes.POST("/", projectController.CreateNewProject)
	projectRoutes.PATCH("/:id", projectController.EditProject)
	projectRoutes.DELETE("/:id", projectController.DeleteProject)

	router.Run()
}
