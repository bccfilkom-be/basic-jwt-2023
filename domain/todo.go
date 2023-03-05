package domain

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Activity string `json:"activity" `
	IdUser   uint   `json:"id_user"`
	User     User   `gorm:"foreignkey:IdUser;association_foreignkey:ID"`
}
