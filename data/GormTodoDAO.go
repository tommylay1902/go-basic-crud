package data

import (
	"errors"

	"github.com/tommylay1902/crudbasic/models"
	"gorm.io/gorm"
)

type GormTodoDAO struct {
	db *gorm.DB
}

func NewGormTodoDAO(db *gorm.DB) *GormTodoDAO {
	return &GormTodoDAO{db: db}
}

// CreateTodo(todo *models.Todo) error
// 	GetAllTodos() ([]*models.Todo, error)
// 	FindByID(id int) (*models.Todo, error)

func (dao *GormTodoDAO) GetAllTodos() ([]models.Todo, error) {
	var todo []models.Todo
	dao.db.Find(&todo)
	return todo, nil
}

func (dao *GormTodoDAO) CreateTodo(todo *models.Todo) error {
	dao.db.Create(todo)
	return nil
}

func (dao *GormTodoDAO) FindByID(id int) (*models.Todo, error) {
	todo := new(models.Todo)
	dao.db.First(todo, id)
	return todo, nil
}

func (dao *GormTodoDAO) DeleteTodo(todo *models.Todo) error {
	if err := dao.db.Delete(&todo); err != nil {
		return errors.New("yup")
	} else {
		return nil
	}
}
