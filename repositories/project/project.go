package repositories

import (
	"porto-be/models"
)

type Repository interface {
	FindAll() ([]models.Project, error)
	FindByID(ID int) (models.Project, error)
	Create(project models.Project) (models.Project, error)
	Update(project models.Project) (models.Project, error)
	Delete(project models.Project) (models.Project, error)
}
