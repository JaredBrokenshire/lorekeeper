package models

type Subclass struct {
	ID   uint `gorm:"primary_key;auto_increment"`
	Name string

	ClassID uint
	Class   *Class
}
