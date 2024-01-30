package services

import (
	"porto-be/models"
	requests "porto-be/requests/project"
)

type Service interface {
	FindAll() ([]models.Project, error)
	FindByID(ID int) (models.Project, error)
	Create(project requests.CreateProjectRequest) (models.Project, error)
	Update(ID int, project requests.UpdateProjectRequest) (models.Project, error)
	Delete(ID int) (models.Project, error)
}
