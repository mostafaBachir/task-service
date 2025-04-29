package project

import (
	"github.com/gofiber/fiber/v2"
)

func CreateProject(c *fiber.Ctx) error {
	var payload CreateProjectRequest
	if err := c.BodyParser(&payload); err != nil {
		return fiber.ErrBadRequest
	}

	if err := ValidateCreateProject(payload); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	project, err := CreateProjectService(payload)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to create project")
	}

	return c.Status(fiber.StatusCreated).JSON(project)
}
