package models

type Class struct {
	ID   uint `gorm:"primary_key;auto_increment"`
	Name string

	Subclasses []*Subclass

	LogoID uint
	Logo   *File `gorm:"foreignkey:LogoID"`
}
