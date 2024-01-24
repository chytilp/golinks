package model

import (
	"gorm.io/gorm"
	"github.com/chytilp/golinks/database"
)

type Category struct {
	gorm.Model
	Name     string `gorm:"size:50;not null;unique" json:"name"`
	ParentID *uint	`json:"parentId"`
}

func (category *Category) Save() (*Category, error) {
	err := database.Database.Create(&category).Error
	if err != nil {
		return &Category{}, err
	}
	return category, nil
}
