package project

import (
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type CreateProjectRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
}

func ValidateCreateProject(input CreateProjectRequest) error {
	return validate.Struct(input)
}
