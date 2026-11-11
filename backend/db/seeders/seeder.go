package seeders

import (
	"github.com/jinzhu/gorm"
	"lorekeeper/pkg/file_storage"
	"os"
	"path/filepath"
)

type Seeder struct {
	DB        *gorm.DB
	FileStore file_storage.Store
}

func NewSeeder(db *gorm.DB, fileStore file_storage.Store) *Seeder {
	return &Seeder{
		DB:        db,
		FileStore: fileStore,
	}
}

func (s *Seeder) Seed() {
	s.SetRoles()
	s.SetUsers()
	s.SetRaces()
	s.SetClasses()
	s.SetClassLogos()
	s.SetSubclasses()
	s.SetCampaigns()
	s.SetCharacters()
}

// resolveDir returns the first path (trying cwd, then ../cwd, then ../../cwd)
// where the directory exists, or empty string if none found.
func resolveDir(path string) string {
	prefixes := []string{"", "../", "../../"}
	for _, prefix := range prefixes {
		candidate := filepath.Join(prefix, path)
		if _, err := os.Stat(candidate); err == nil {
			return candidate
		}
	}
	return ""
}
