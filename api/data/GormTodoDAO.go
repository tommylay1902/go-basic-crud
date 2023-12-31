package data

import (
	"errors"

	"github.com/tommylay1902/crudbasic/internal/error/customerrors"
	"github.com/tommylay1902/crudbasic/internal/models"
	"gorm.io/gorm"
)

type GormTodoDAO struct {
	db *gorm.DB
}

func NewGormTodoDAO(db *gorm.DB) *GormTodoDAO {
	return &GormTodoDAO{db: db}
}

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
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &customerrors.ResourceNotFound{
				Message: "Todo not found",
				Code:    404,
			}
		}
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

func (dao *GormTodoDAO) UpdateTodo(todo *models.Todo) error {
	err := dao.db.Save(&todo).Error
	if err != nil {
		return err
	}
	return nil
}
