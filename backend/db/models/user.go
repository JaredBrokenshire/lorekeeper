package models

import "time"

type User struct {
	ID       uint   `gorm:"primary_key;auto_increment"`
	Username string `gorm:"unique_index"`
	Email    string `gorm:"unique_index"`
	Password string

	Roles []*Role `gorm:"many2many:user_roles"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
