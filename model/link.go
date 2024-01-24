package model

import (
	"gorm.io/gorm"
	"github.com/chytilp/golinks/database"
)

type Link struct {
	gorm.Model
	Name     string   `gorm:"size:255;not null;unique" json:"name"`
	Address  string   `gorm:"size:255;not null;" json:"address"`
	CategoryID uint   `gorm:"not null;" json:"categoryId"`
	Category Category
	Roles    []Role   `gorm:"many2many:role_links;"`
	Users    []UserLink
}

func (link *Link) Save() (*Link, error) {
	err := database.Database.Create(&link).Error
	if err != nil {
		return &Link{}, err
	}
	return link, nil
}
