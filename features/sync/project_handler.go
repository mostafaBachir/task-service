package sync

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mostafaBachir/task-service/database"
	"github.com/mostafaBachir/task-service/models"
)

// Créer un nouveau projet
func CreateProject(c *fiber.Ctx) error {
	project := new(models.Project)

	// Parse le JSON du corps de la requête
	if err := c.BodyParser(project); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Requête invalide",
		})
	}

	// Enregistre dans la base de données
	if err := database.DB.Create(&project).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Erreur lors de la création du projet",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(project)
}
