package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/mostafaBachir/task-service/config"
	"github.com/mostafaBachir/task-service/database"
	"github.com/mostafaBachir/task-service/routes"

)

func main() {
	// 1. Charger la configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("❌ Erreur de chargement de la configuration :", err)
	}

	// 2. Initialiser la base de données
	database.InitDatabase(cfg)

	// 3. Créer l'application Fiber
	app := fiber.New()

	routes.SetupRoutes(app)

	// 4. Route simple pour tester
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong 🏓")
	})

	// 5. Démarrer le serveur
	log.Println("🚀 Server Fiber démarré sur le port", cfg.Port)
	app.Listen(":" + cfg.Port)
}
