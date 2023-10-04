package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tommylay1902/crudbasic/handlers"
)

// SetupRoutes setup router api
func SetupTodoRoutes(app *fiber.App, todoHandler *handlers.TodoHandler) {
	// Middleware
	api := app.Group("/api/v1/todos")

	api.Post("", todoHandler.CreateTodo)
	api.Get("", todoHandler.GetAllTodos)
	api.Get("/:id", todoHandler.GetTodoById)
	api.Delete("/:id", todoHandler.DeleteTodoById)
}
