package data

import (
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
	err := dao.db.Find(&todo).Error
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (dao *GormTodoDAO) CreateTodo(todo *models.Todo) error {
	err := dao.db.Create(todo).Error
	if err != nil {
		return err
	}
	return nil
}

func (dao *GormTodoDAO) FindByID(id int) (*models.Todo, error) {
	todo := new(models.Todo)
	err := dao.db.First(todo, id).Error
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (dao *GormTodoDAO) DeleteTodo(todo *models.Todo) error {
	err := dao.db.Delete(&todo).Error
	if err != nil {
		return err
	} else {
		return nil
	}
}
