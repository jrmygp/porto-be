package repositories

import (
	"porto-be/models"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]models.Project, error)
	FindByID(ID int) (models.Project, error)
	Create(project models.Project) (models.Project, error)
	Update(project models.Project) (models.Project, error)
	Delete(project models.Project) (models.Project, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]models.Project, error) {
	var projects []models.Project

	err := r.db.Find(&projects).Error

	return projects, err
}

func (r *repository) FindByID(ID int) (models.Project, error) {
	var project models.Project

	err := r.db.Find(&project, ID).Error

	return project, err
}

func (r *repository) Create(project models.Project) (models.Project, error) {
	err := r.db.Create(&project).Error

	return project, err
}

func (r *repository) Update(project models.Project) (models.Project, error) {
	err := r.db.Save(&project).Error
	return project, err
}

func (r *repository) Delete(project models.Project) (models.Project, error) {
	err := r.db.Delete(project).Error
	return project, err
}
