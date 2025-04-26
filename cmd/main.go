package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/mostafaBachir/task-service/config"
	"github.com/mostafaBachir/task-service/database"
	// ajoute ceci aussi si pas encore fait
	"github.com/mostafaBachir/task-service/models"
)

func main() {
	// 1. Charger la configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("❌ Erreur de chargement de la configuration :", err)
	}

	// 2. Connexion à la base de données
	database.ConnectDatabase(cfg.DatabaseURL)
	
	err = database.DB.AutoMigrate(&models.Project{})
	if err != nil {
		log.Fatal("❌ Migration Project échouée :", err)
	}
	// 3. Créer l'application Fiber
	app := fiber.New()

	// 4. Route simple pour tester
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong 🏓")
	})

	// 5. Démarrer le serveur
	log.Println("🚀 Server Fiber démarré sur le port", cfg.Port)
	app.Listen(":" + cfg.Port)
}
