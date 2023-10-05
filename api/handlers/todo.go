package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/tommylay1902/crudbasic/api/services"
	"github.com/tommylay1902/crudbasic/internal/dtos"
	"github.com/tommylay1902/crudbasic/internal/error/customerrors"
	"github.com/tommylay1902/crudbasic/internal/error/errorhandler"
	"github.com/tommylay1902/crudbasic/internal/models"
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
		return errorhandler.HandleError(
			&customerrors.BadRequestError{
				Message: err.Error(),
				Code:    400,
			}, c)

	}

	err := tdh.TodoService.CreateTodo(&requestBody)

	if err != nil {
		return errorhandler.HandleError(err, c)
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Todo created successfully",
	})
}

func (tdh *TodoHandler) GetTodoById(c *fiber.Ctx) error {

	idParam := c.Params("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		return errorhandler.HandleError(
			&customerrors.BadRequestError{
				Message: err.Error(),
				Code:    400,
			}, c)
	}

	todo, serviceErr := tdh.TodoService.GetTodoById(id)

	if serviceErr != nil {
		return errorhandler.HandleError(serviceErr, c)
	}

	return c.JSON(todo)

}

func (th *TodoHandler) GetAllTodos(c *fiber.Ctx) error {

	todos, err := th.TodoService.GetAllTodos()

	if err != nil {
		return errorhandler.HandleError(err, c)
	}

	return c.JSON(todos)
}

func (th *TodoHandler) DeleteTodoById(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		return errorhandler.HandleError(
			&customerrors.BadRequestError{
				Message: err.Error(),
				Code:    400,
			}, c)
	}

	serviceErr := th.TodoService.DeleteTodo(id)

	if serviceErr != nil {
		return errorhandler.HandleError(serviceErr, c)

	}
	c.Status(fiber.StatusOK)
	return nil
}

func (tdh *TodoHandler) UpdateTodo(c *fiber.Ctx) error {
	var requestBody dtos.TodoDTO
	if err := c.BodyParser(&requestBody); err != nil {
		return errorhandler.HandleError(
			&customerrors.BadRequestError{
				Message: err.Error(),
				Code:    400,
			}, c)

	}

	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return errorhandler.HandleError(
			&customerrors.BadRequestError{
				Message: err.Error(),
				Code:    400,
			}, c)
	}
	serviceErr := tdh.TodoService.UpdateTodo(id, &requestBody)
	if serviceErr != nil {
		return errorhandler.HandleError(serviceErr, c)

	}

	return nil
}
