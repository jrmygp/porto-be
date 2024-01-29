package services

import (
	"mime/multipart"
	"path/filepath"
	"porto-be/models"
	repositories "porto-be/repositories/project"
	requests "porto-be/requests/project"
)

type service struct {
	repository repositories.Repository
}

func NewService(repository repositories.Repository) *service {
	return &service{repository}
}

// Private function
func convertFileToPath(file *multipart.FileHeader) string {
	baseDirectory := "public/project/"

	filePath := filepath.Join(baseDirectory, file.Filename)

	return filePath
}

func (s *service) FindAll() ([]models.Project, error) {
	projects, err := s.repository.FindAll()
	return projects, err
}

func (s *service) FindByID(ID int) (models.Project, error) {
	project, err := s.repository.FindByID(ID)
	return project, err
}

func (s *service) Create(projectForm requests.CreateProjectRequest) (models.Project, error) {
	// convert yang dari bentukan awalnya sebuah form jadiin ke bentuk model
	project := models.Project{
		Title:       projectForm.Title,
		Description: projectForm.Description,
		Url:         convertFileToPath(projectForm.Url),
	}

	newProject, err := s.repository.Create(project)
	return newProject, err
}

func (s *service) Update(ID int, projectForm requests.CreateProjectRequest) (models.Project, error) {
	p, _ := s.repository.FindByID(ID)

	p.Title = projectForm.Title
	p.Description = projectForm.Description
	p.Url = convertFileToPath(projectForm.Url)

	newProject, err := s.repository.Update(p)
	return newProject, err
}

func (s *service) Delete(ID int) (models.Project, error) {
	book, _ := s.repository.FindByID(ID)
	newBook, err := s.repository.Delete(book)

	return newBook, err
}
