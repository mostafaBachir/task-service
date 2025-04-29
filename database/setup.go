package database

import (
	"log"

	"github.com/mostafaBachir/task-service/config"
	"github.com/mostafaBachir/task-service/models"
)

func InitDatabase(cfg *config.Config) {
	ConnectDatabase(cfg.DatabaseURL)

	err := DB.AutoMigrate(
		&models.Project{},
		&models.Task{},
	)
	if err != nil {
		log.Fatal("❌ Échec de la migration automatique :", err)
	}
}
