package model

import (
	"gorm.io/gorm"
	"github.com/chytilp/golinks/database"
)

type Role struct {
	gorm.Model
	Name  string `gorm:"size:50;not null;unique" json:"name"`
	Users []User `gorm:"many2many:user_roles;"`
	Links []Link `gorm:"many2many:role_links;"`
}

func (role *Role) Save() (*Role, error) {
	err := database.Database.Create(&role).Error
	if err != nil {
		return &Role{}, err
	}
	return role, nil
}
