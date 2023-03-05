package repository

import (
	"github.com/bccfilkom-be/basic-jwt-2023/domain"
	"gorm.io/gorm"
)

type (
	userRepository struct {
		db *gorm.DB
	}
	UserRepository interface {
		Create(user *domain.User) error
		FindById(id uint) (*domain.User, error)
		FindAllByIdUser(idUser uint) ([]*domain.User, error)
		Delete(id uint) error
		Update(user *domain.User) error
	}
)

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

// Todo: Implement Interface User
