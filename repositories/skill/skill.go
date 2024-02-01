package repositories

import (
	"porto-be/models"
)

type Repository interface {
	FindAll() ([]models.Skill, error)
	FindByID(ID int) (models.Skill, error)
	Create(skill models.Skill) (models.Skill, error)
	Update(skill models.Skill) (models.Skill, error)
	Delete(skill models.Skill) (models.Skill, error)
}
