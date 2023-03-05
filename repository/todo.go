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
		FindAllByIdTodo(idTodo uint) ([]*domain.Todo, error)
		Delete(id uint) error
		Update(todo *domain.Todo) error
	}
)

func NewTodoRepository(db *gorm.DB) TodoRepository {
	return &todoRepository{db: db}
}

func (tr *todoRepository) Create(todo *domain.Todo) error {
	if err := tr.db.Preload("User").Create(&todo).Error; err != nil {
		return err
	}
	return nil
}

func (tr *todoRepository) FindById(idTodo uint) (*domain.Todo, error) {
	var todo *domain.Todo
	if err := tr.db.First(&todo, idTodo).Error; err != nil {
		return nil, err
	}
	return todo, nil
}

func (tr *todoRepository) FindAllByIdTodo(idTodo uint) ([]*domain.Todo, error) {
	var todos []*domain.Todo
	if err := tr.db.Find(&todos, idTodo).Error; err != nil {
		return nil, err
	}
	return todos, nil
}

func (tr *todoRepository) Delete(id uint) error {
	var todo domain.Todo
	if err := tr.db.Delete(&todo, id).Error; err != nil {
		return err
	}
	return nil
}

func (tr *todoRepository) Update(todo *domain.Todo) error {
	var updateTodo domain.Todo
	if err := tr.db.Model(&updateTodo).Updates(todo).Error; err != nil {
		return err
	}
	return nil
}

// Todo: Implement Interface Todo
