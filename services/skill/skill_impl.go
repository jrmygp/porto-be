package services

import (
	"mime/multipart"
	"path/filepath"
	"porto-be/models"
	repositories "porto-be/repositories/skill"
	requests "porto-be/requests/skill"
)

type service struct {
	repository repositories.Repository
}

func NewService(repository repositories.Repository) *service {
	return &service{repository}
}

// Private function
func convertFileToPath(file *multipart.FileHeader) string {
	baseDirectory := "public/skill/"

	filePath := filepath.Join(baseDirectory, file.Filename)

	return filePath
}

func (s *service) FindAll() ([]models.Skill, error) {
	skills, err := s.repository.FindAll()
	return skills, err
}

func (s *service) FindByID(ID int) (models.Skill, error) {
	skill, err := s.repository.FindByID(ID)
	return skill, err
}

func (s *service) Create(skillRequest requests.CreateSkillRequest) (models.Skill, error) {
	skill := models.Skill{
		Title: skillRequest.Title,
		Image: convertFileToPath(skillRequest.Image),
	}

	newSkill, err := s.repository.Create(skill)
	return newSkill, err
}

func (s *service) Update(ID int, skillRequest requests.UpdateSkillRequest) (models.Skill, error) {
	found, _ := s.repository.FindByID(ID)

	if skillRequest.Title != "" {
		found.Title = skillRequest.Title
	}
	if skillRequest.Image != nil {
		found.Image = convertFileToPath(skillRequest.Image)
	}

	newSkill, err := s.repository.Update(found)

	return newSkill, err
}

func (s *service) Delete(ID int) (models.Skill, error) {
	skill, _ := s.repository.FindByID(ID)
	deletedSkill, err := s.repository.Delete(skill)

	return deletedSkill, err
}
