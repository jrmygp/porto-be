package services

import (
	"porto-be/models"
	requests "porto-be/requests/skill"
)

type Service interface {
	FindAll() ([]models.Skill, error)
	FindByID(ID int) (models.Skill, error)
	Create(skill requests.CreateSkillRequest) (models.Skill, error)
	Update(ID int, skill requests.UpdateSkillRequest) (models.Skill, error)
	Delete(ID int) (models.Skill, error)
}
