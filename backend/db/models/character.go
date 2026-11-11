package models

type Character struct {
	ID     uint `gorm:"primary_key;auto_increment"`
	UserID uint
	User   *User

	Name  string
	Level int

	RaceID uint
	Race   *Race

	Classes    []*Class    `gorm:"many2many:character_classes;"`
	Subclasses []*Subclass `gorm:"many2many:character_subclasses;"`

	CampaignID uint
	Campaign   *Campaign

	ProfilePictureID uint
	ProfilePicture   *File `gorm:"foreignkey:ProfilePictureID"`
}
