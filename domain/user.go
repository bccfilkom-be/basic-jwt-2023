package domain

import "gorm.io/gorm"

type (
	User struct {
		gorm.Model
		Username string `json:"username"`
		Password string `json:"password"`
		Role     string `json:"role"`
		Name     string `json:"name"`
		Todos    []Todo `json:"todos" gorm:"foreignkey:IdUser"`
	}
	UserRegister struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Role     string `json:"role"`
		Name     string `json:"name"`
	}
)
