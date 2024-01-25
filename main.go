package main

import (
	projectcontroller "porto-be/controllers/projectController"
	"porto-be/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.DatabaseConnection()

	r.GET("/api/projects", projectcontroller.FindAll)

	r.Run()
}
