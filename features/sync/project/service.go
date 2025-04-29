package project

import (
	"github.com/mostafaBachir/task-service/database"
	"github.com/mostafaBachir/task-service/models"
)
func CreateProjectService(input CreateProjectRequest) (*models.Project, error) {
	project := models.Project{
		Titre:       input.Name,
		Description: input.Description,
	}

	db := database.DB
	if err := db.Create(&project).Error; err != nil {
		return nil, err
	}

	return &project, nil
}
