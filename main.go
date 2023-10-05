package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tommylay1902/crudbasic/api/data"
	"github.com/tommylay1902/crudbasic/api/handlers"
	"github.com/tommylay1902/crudbasic/api/router"
	"github.com/tommylay1902/crudbasic/api/services"
	"github.com/tommylay1902/crudbasic/internal/config/db"
	"github.com/tommylay1902/crudbasic/internal/config/environment"
)

func main() {

	port := environment.SetupEnvironment()
	db := db.SetupDB()

	defer func() {
		dbInstance, _ := db.DB()
		_ = dbInstance.Close()
	}()

	app := fiber.New()

	todoDAO := data.NewGormTodoDAO(db)
	todoService := services.NewTodoService(todoDAO)
	todoHandler := handlers.InitializeTodoHandler(todoService)

	router.SetupTodoRoutes(app, todoHandler)

	app.Listen(":" + port)

}
