package handlers

import (
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/tommylay1902/crudbasic/api/services"
	"github.com/tommylay1902/crudbasic/internal/dtos"
	"github.com/tommylay1902/crudbasic/internal/error/errorhandler"
	"github.com/tommylay1902/crudbasic/internal/models"
	"gorm.io/gorm"
)

type TodoHandler struct {
	TodoService *services.TodoService
}

func InitializeTodoHandler(todoService *services.TodoService) *TodoHandler {
	return &TodoHandler{TodoService: todoService}
}

func (tdh *TodoHandler) CreateTodo(c *fiber.Ctx) error {

	var requestBody models.Todo

	if err := c.BodyParser(&requestBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	err := tdh.TodoService.CreateTodo(&requestBody)

	if err != nil {

		if errors.Is(err, gorm.ErrPrimaryKeyRequired) {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{
				"error": "id already exists",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Todo created successfully",
	})
}

func (tdh *TodoHandler) GetTodoById(c *fiber.Ctx) error {

	idParam := c.Params("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Bad Request",
		})
	}

	todo, serviceErr := tdh.TodoService.GetTodoById(id)

	if serviceErr != nil {
		errResponse := errorhandler.HandleError(serviceErr, c)
		return errResponse
	}

	return c.JSON(todo)

}

func (th *TodoHandler) GetAllTodos(c *fiber.Ctx) error {

	todos, err := th.TodoService.GetAllTodos()

	if err != nil {
		errResponse := errorhandler.HandleError(err, c)
		return errResponse
	}

	return c.JSON(todos)
}

func (th *TodoHandler) DeleteTodoById(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "bad request",
		})
	}

	serviceErr := th.TodoService.DeleteTodo(id)

	if serviceErr != nil {
		errorhandler.HandleError(serviceErr, c)
	}
	c.Status(fiber.StatusOK)
	return nil
}

func (tdh *TodoHandler) UpdateTodo(c *fiber.Ctx) error {
	var requestBody dtos.TodoDTO
	if err := c.BodyParser(&requestBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {

	}

	serviceErr := tdh.TodoService.UpdateTodo(id, &requestBody)
	if serviceErr != nil {
		errResponse := errorhandler.HandleError(serviceErr, c)
		return errResponse
	}

	return nil
}
