package seeders

import (
	"github.com/jinzhu/gorm"
)

type Seeder struct {
	DB *gorm.DB
}

func NewSeeder(db *gorm.DB) *Seeder {
	return &Seeder{
		DB: db,
	}
}

func (s *Seeder) Seed() {
	s.SetRoles()
	s.SetUsers()
	s.SetRaces()
	s.SetClasses()
	s.SetSubclasses()
	s.SetCampaigns()
	s.SetCharacters()
}
