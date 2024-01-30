package services

import (
	"porto-be/models"
	requests "porto-be/requests/tech"
)

type Service interface {
	FindAll() ([]models.Tech, error)
	FindByID(ID int) (models.Tech, error)
	Create(tech requests.CreateTechRequest) (models.Tech, error)
	Update(ID int, tech requests.UpdateTechRequest) (models.Tech, error)
	Delete(ID int) (models.Tech, error)
}
