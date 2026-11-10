package models

import "time"

type Campaign struct {
	ID   uint `gorm:"primary_key;auto_increment"`
	Name string

	DungeonMasterID uint
	DungeonMaster   *User `gorm:"foreignkey:DungeonMasterID"`

	Characters []*Character

	CreatedAt time.Time
	UpdatedAt time.Time
}
