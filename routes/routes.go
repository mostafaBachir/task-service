package routes

import (
	projectsync "github.com/mostafaBachir/task-service/features/sync/project"
	//tasksync "github.com/mostafaBachir/task-service/features/sync/task"
	projectws "github.com/mostafaBachir/task-service/features/async/project"
	//taskws "github.com/mostafaBachir/task-service/features/async/task"
	"github.com/mostafaBachir/task-service/security"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", security.JWTMiddleware)

	userProjects := api.Group("/:user_id/projects")
	userProjects.Post("/", projectsync.CreateProject)
	app.Use("/ws/projects", projectws.WebSocketAuthMiddleware)
	app.Get("/ws/projects", websocket.New(projectws.WebSocketConnect))
	// 🔹 Projets d'un utilisateur
	//userProjects := api.Group("/:user_id/projects")
	//userProjects.Get("/", project.ListProjects)
	//userProjects.Post("/", project.CreateProject)
	//userProjects.Get("/:project_id", project.GetProject)
	//userProjects.Put("/:project_id", project.UpdateProject)
	//userProjects.Delete("/:project_id", project.DeleteProject)

	// 🔹 Tâches d'un projet
	//projectTasks := userProjects.Group("/:project_id/tasks")
	//projectTasks.Get("/", task.ListTasks)
	//projectTasks.Post("/", task.CreateTask)
	//projectTasks.Get("/:task_id", task.GetTask)
	//projectTasks.Put("/:task_id", task.UpdateTask)
	//projectTasks.Delete("/:task_id", task.DeleteTask)

	// 🔹 (optionnel) changer statut d'une tâche
	//projectTasks.Patch("/:task_id/status", task.UpdateTaskStatus)

	// 🔹 (optionnel) ajouter ou retirer un participant
	//projectTasks.Post("/:task_id/participants", task.AddParticipant)
	//projectTasks.Delete("/:task_id/participants/:participant_id", task.RemoveParticipant)

	// 🔹 (optionnel) upload fichier
	//projectTasks.Post("/:task_id/upload", task.UploadFile)
}
