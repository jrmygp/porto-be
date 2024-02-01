package repositories

import (
	"porto-be/models"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]models.Project, error) {
	var projects []models.Project

	err := r.db.Preload("Stacks").Find(&projects).Error

	return projects, err
}

func (r *repository) FindByID(ID int) (models.Project, error) {
	var project models.Project

	err := r.db.Preload("Stacks").Find(&project, ID).Error

	return project, err
}

func (r *repository) Create(project models.Project) (models.Project, error) {
	err := r.db.Create(&project).Error

	// Add data to pivot table
	for _, stackID := range project.Stack_id {
		projectStack := new(models.ProjectStack)
		projectStack.ProjectID = project.ID
		projectStack.SkillID = stackID
		r.db.Create(&projectStack)
	}

	return project, err
}

func (r *repository) Update(project models.Project) (models.Project, error) {
	err := r.db.Save(&project).Error

	if len(project.Stack_id) > 0 {
		for _, stackID := range project.Stack_id {
			projectStack := new(models.ProjectStack)
			projectStack.ProjectID = project.ID
			projectStack.SkillID = stackID
			r.db.Create(&projectStack)
		}
	}

	return project, err
}

func (r *repository) Delete(project models.Project) (models.Project, error) {
	// Delete associated rows in project_stacks table
	projectStack := new(models.ProjectStack)
	r.db.Where("project_id = ?", project.ID).Delete(&projectStack)

	err := r.db.Delete(project).Error

	return project, err
}
