package services

import (
	"fmt"

	"github.com/tommylay1902/crudbasic/api/data"
	"github.com/tommylay1902/crudbasic/internal/dtos"
	"github.com/tommylay1902/crudbasic/internal/error/customerrors"
	"github.com/tommylay1902/crudbasic/internal/models"
)

type TodoService struct {
	TodoDAO data.TodoDAO
}

func NewTodoService(todoDAO data.TodoDAO) *TodoService {
	return &TodoService{TodoDAO: todoDAO}
}

func (tds *TodoService) CreateTodo(todo *models.Todo) error {
	err := tds.TodoDAO.CreateTodo(todo)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func (tds *TodoService) GetTodoById(todoID int) (*models.Todo, error) {

	todo, err := tds.TodoDAO.FindByID(todoID)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (tds *TodoService) GetAllTodos() ([]models.Todo, error) {
	todo, err := tds.TodoDAO.GetAllTodos()
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (tds *TodoService) DeleteTodo(id int) error {
	todo, err := tds.TodoDAO.FindByID(id)
	if err != nil {
		return err
	}
	if err := tds.TodoDAO.DeleteTodo(todo); err != nil {
		return err
	}
	return nil
}

func (tds *TodoService) UpdateTodo(id int, todo *dtos.TodoDTO) error {
	update, err := tds.TodoDAO.FindByID(id)
	if err != nil {
		return err
	}
	hasUpdate := false

	if todo.Todo != nil && *todo.Todo != update.Todo {
		hasUpdate = true
		update.Todo = *todo.Todo
	}
	if todo.Completed != nil && *todo.Completed != update.Completed {
		hasUpdate = true
		update.Completed = *todo.Completed
	}
	if hasUpdate {
		if err := tds.TodoDAO.UpdateTodo(update); err != nil {
			return err
		}
	} else {
		return &customerrors.BadRequestError{
			Message: "no update detected!",
			Code:    400,
		}
	}

	return nil
}
