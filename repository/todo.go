package repository

import (
	"github.com/bccfilkom-be/basic-jwt-2023/domain"
	"gorm.io/gorm"
)

type (
	todoRepository struct {
		db *gorm.DB
	}
	TodoRepository interface {
		Create(todo *domain.Todo) error
		FindById(id uint) (*domain.Todo, error)
		FindAllByIdUser(idUser uint) ([]*domain.Todo, error)
		Delete(id uint) error
		Update(todo *domain.Todo) error
	}
)

func NewTodoRepository(db *gorm.DB) TodoRepository {
	return &todoRepository{db: db}
}

// Todo: Implement Interface Todo
