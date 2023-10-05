package data

import "github.com/tommylay1902/crudbasic/internal/models"

type TodoDAO interface {
	CreateTodo(todo *models.Todo) error
	GetAllTodos() ([]models.Todo, error)
	FindByID(id int) (*models.Todo, error)
	DeleteTodo(*models.Todo) error
	UpdateTodo(*models.Todo) error
}
