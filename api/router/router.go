package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tommylay1902/crudbasic/api/handlers"
)

func SetupTodoRoutes(app *fiber.App, todoHandler *handlers.TodoHandler) {
	api := app.Group("/api/v1/todos")

	api.Post("", todoHandler.CreateTodo)
	api.Get("", todoHandler.GetAllTodos)
	api.Get("/:id", todoHandler.GetTodoById)
	api.Delete("/:id", todoHandler.DeleteTodoById)
	api.Put("/:id", todoHandler.UpdateTodo)
}
