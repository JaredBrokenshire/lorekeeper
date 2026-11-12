package models

type Race struct {
	ID   uint `gorm:"primary_key;auto_increment"`
	Name string
}
