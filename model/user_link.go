package model

import (
	"time"
	"gorm.io/gorm"
	"github.com/chytilp/golinks/database"
)

type UserLink struct {
	UserID  	uint 	`gorm:"primaryKey"`
	User        User 	`gorm:"foreignKey:UserID;references:ID"`
	LinkID      uint 	`gorm:"primaryKey"`
	Link        Link 	`gorm:"foreignKey:LinkID;references:ID"`
	Stars       uint8
	Note        string 	`gorm:"size:256"`
	NotePrivate bool
	Owner		bool	
	CreatedAt time.Time
  	UpdatedAt time.Time
  	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (userlink *UserLink) Save() (*UserLink, error) {
	err := database.Database.Create(&userlink).Error
	if err != nil {
		return &UserLink{}, err
	}
	return userlink, nil
}
// pozor pred automigrate spustit
// db.SetupJoinTable(&User{}, "Links", &UserLink{})
// db.AutoMigrate(
//     &User{},
//     &Link{},
// )
