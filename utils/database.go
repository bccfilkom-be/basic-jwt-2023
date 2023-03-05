package utils

import (
	"fmt"
	"os"

	"github.com/bccfilkom-be/basic-jwt-2023/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DriverDatabase struct {
	User     string
	Password string
	Host     string
	Port     string
	DbName   string
}

func ReadEnvDatabase() (DriverDatabase, error) {
	return DriverDatabase{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		DbName:   os.Getenv("DB_NAME"),
	}, nil
}

func MakeConnection(data DriverDatabase) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", data.User, data.Password, data.Host, data.Port, data.DbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if err := db.AutoMigrate(&domain.User{}, &domain.Todo{}); err != nil {
		return nil, err
	}
	return db, nil
}
