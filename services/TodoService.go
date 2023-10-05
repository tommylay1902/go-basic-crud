package services

import (
	"fmt"

	"github.com/tommylay1902/crudbasic/data"
	"github.com/tommylay1902/crudbasic/models"
)

type TodoService struct {
	// You can add any dependencies or database connections here
	TodoDAO data.TodoDAO
}

func NewTodoService(todoDAO data.TodoDAO) *TodoService {
	return &TodoService{TodoDAO: todoDAO}
}

// CreateUser creates a new user.
func (tds *TodoService) CreateTodo(todo *models.Todo) error {
	err := tds.TodoDAO.CreateTodo(todo)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

// GetUserByID retrieves a user by ID.
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
