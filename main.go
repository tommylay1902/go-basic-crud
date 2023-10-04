package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/tommylay1902/crudbasic/data"
	"github.com/tommylay1902/crudbasic/handlers"
	"github.com/tommylay1902/crudbasic/models"
	"github.com/tommylay1902/crudbasic/router"
	"github.com/tommylay1902/crudbasic/services"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	portString := os.Getenv("PORT")

	if portString == "" {
		log.Fatal("Port is not specified")
	}

	dsn := "host=localhost user=postgres password=mysecretpassword dbname=postgres port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database")
	}

	defer func() {
		dbInstance, _ := db.DB()
		_ = dbInstance.Close()
	}()

	db.AutoMigrate(&models.Todo{})

	app := fiber.New()

	todoDAO := data.NewGormTodoDAO(db)
	todoService := services.NewTodoService(todoDAO)
	todoHandler := handlers.InitializeTodoHandler(todoService)

	router.SetupTodoRoutes(app, todoHandler)

	app.Listen(":" + portString)

}
