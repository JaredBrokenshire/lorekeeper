package models

type Character struct {
	ID     uint `gorm:"primary_key;auto_increment"`
	UserID uint

	Name  string
	Level int

	RaceID uint
	Race   *Race

	Classes    []*Class    `gorm:"many2many:character_classes;"`
	Subclasses []*Subclass `gorm:"many2many:character_subclasses;"`
}
