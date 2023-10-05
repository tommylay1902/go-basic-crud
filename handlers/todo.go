package handlers

import (
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/tommylay1902/crudbasic/models"
	"github.com/tommylay1902/crudbasic/services"
	"gorm.io/gorm"
)

// UserHandler is a handler for user-related routes.
type TodoHandler struct {
	TodoService *services.TodoService
}

// InitializeUserHandler creates a new UserHandler with the UserService dependency.
func InitializeTodoHandler(todoService *services.TodoService) *TodoHandler {
	return &TodoHandler{TodoService: todoService}
}

// Define route handling functions that use the UserService.
func (tdh *TodoHandler) CreateTodo(c *fiber.Ctx) error {
	// Parse request data, call UserService functions, and generate response
	// Example: uh.UserService.CreateUser(user)
	// Create a variable to hold the parsed request body
	var requestBody models.Todo

	// Parse the request body into the struct
	if err := c.BodyParser(&requestBody); err != nil {
		// Handle parsing errors, possibly by sending an error response
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Call your service or DAO to create the todo using `newTodo`
	tdh.TodoService.CreateTodo(&requestBody)

	// Send a success response
	return c.JSON(fiber.Map{
		"message": "Todo created successfully",
	})
}

func (tdh *TodoHandler) GetTodoById(c *fiber.Ctx) error {

	idParam := c.Params("id")
	// Convert the ID string to an integer
	id, err := strconv.Atoi(idParam)
	if err != nil {
		// Handle invalid or non-integer ID (e.g., return an error response)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}

	todo, serviceErr := tdh.TodoService.GetTodoById(id)

	if serviceErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}

	return c.JSON(todo)

}

func (th *TodoHandler) GetAllTodos(c *fiber.Ctx) error {
	// Retrieve todos from the service or DAO
	todos, err := th.TodoService.GetAllTodos()
	if err != nil {

	}

	// Send the todos as a JSON response
	return c.JSON(todos)
}

func (th *TodoHandler) DeleteTodoById(c *fiber.Ctx) error {
	idParam := c.Params("id")
	// Convert the ID string to an integer
	id, err := strconv.Atoi(idParam)
	if err != nil {
		// Handle invalid or non-integer ID (e.g., return an error response)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Bad Request",
		})

	}

	serviceErr := th.TodoService.DeleteTodo(id)

	if serviceErr != nil {

		if errors.Is(serviceErr, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(
				fiber.Map{
					"error": "Todo was not found",
				})
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(
				fiber.Map{
					"error": "Server error",
				})
		}
	}

	c.Status(fiber.StatusOK)
	return nil
}
