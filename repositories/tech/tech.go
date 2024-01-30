package repositories

import (
	"porto-be/models"
)

type Repository interface {
	FindAll() ([]models.Tech, error)
	FindByID(ID int) (models.Tech, error)
	Create(tech models.Tech) (models.Tech, error)
	Update(tech models.Tech) (models.Tech, error)
	Delete(tech models.Tech) (models.Tech, error)
}
