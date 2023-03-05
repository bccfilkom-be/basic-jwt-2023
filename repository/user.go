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
		FindAllByIdUser(idUser uint) (*[]domain.User, error)
		Delete(id uint) error
		Update(user *domain.User) error
	}
)

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func(ur *userRepository) Create(user *domain.User) error {
	if err := ur.db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func(ur *userRepository) FindById(idUser uint) (*domain.User, error) {
	var user domain.User
	if err := ur.db.First(&user, idUser).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func(ur *userRepository) FindAllByIdUser(idUser uint) (*[]domain.User, error) {
	var users []domain.User
	if err := ur.db.Find(&users, idUser).Error; err != nil {
		return nil, err
	}
	return &users, nil
}

func(ur *userRepository) Delete(id uint) error {
	var user domain.User
	if err := ur.db.Delete(&user, id).Error; err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) Update(user *domain.User) error {
	var updateUser domain.User
	if err := ur.db.Model(&updateUser).Updates(user).Error; err != nil {
		return err
	}
	return nil
}

// Todo: Implement Interface User
