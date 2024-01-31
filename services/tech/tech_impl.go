package services

import (
	"mime/multipart"
	"path/filepath"
	"porto-be/models"
	repositories "porto-be/repositories/tech"
	requests "porto-be/requests/tech"
)

type service struct {
	repository repositories.Repository
}

func NewService(repository repositories.Repository) *service {
	return &service{repository}
}

// Private function
func convertFileToPath(file *multipart.FileHeader) string {
	baseDirectory := "public/tech/"

	filePath := filepath.Join(baseDirectory, file.Filename)

	return filePath
}

func (s *service) FindAll() ([]models.Tech, error) {
	techs, err := s.repository.FindAll()
	return techs, err
}

func (s *service) FindByID(ID int) (models.Tech, error) {
	tech, err := s.repository.FindByID(ID)
	return tech, err
}

func (s *service) Create(techRequest requests.CreateTechRequest) (models.Tech, error) {
	tech := models.Tech{
		Title:      techRequest.Title,
		Image:      convertFileToPath(techRequest.Image),
		Percentage: techRequest.Percentage,
	}

	newTech, err := s.repository.Create(tech)
	return newTech, err
}

func (s *service) Update(ID int, techRequest requests.UpdateTechRequest) (models.Tech, error) {
	t, _ := s.repository.FindByID(ID)

	if techRequest.Title != "" {
		t.Title = techRequest.Title
	}
	if techRequest.Image != nil {
		t.Image = convertFileToPath(techRequest.Image)
	}
	if techRequest.Percentage != 0 {
		t.Percentage = techRequest.Percentage
	}

	newTech, err := s.repository.Update(t)
	return newTech, err
}

func (s *service) Delete(ID int) (models.Tech, error) {
	tech, _ := s.repository.FindByID(ID)
	deletedTech, err := s.repository.Delete(tech)

	return deletedTech, err
}
