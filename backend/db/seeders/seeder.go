package seeders

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"lorekeeper/pkg/document"
	"lorekeeper/pkg/file_storage"
	"os"
	"path/filepath"
)

type Seeder struct {
	DB        *gorm.DB
	FileStore file_storage.Store
	Extractor document.ExtractorInterface
}

func NewSeeder(db *gorm.DB, fileStore file_storage.Store) *Seeder {
	return &Seeder{
		DB:        db,
		FileStore: fileStore,
		Extractor: document.NewExtractor(),
	}
}

func (s *Seeder) Seed() {
	fmt.Println("Seeding roles")
	s.SetRoles()
	fmt.Println("Seeded roles")

	fmt.Println("Seeding users")
	s.SetUsers()
	fmt.Println("Seeded users")

	fmt.Println("Seeding races")
	s.SetRaces()
	fmt.Println("Seeded races")

	fmt.Println("Seeding classes")
	s.SetClasses()
	fmt.Println("Seeded classes")

	fmt.Println("Seeding class logos")
	s.SetClassLogos()
	fmt.Println("Seeded class logos")

	fmt.Println("Seeding subclasses")
	s.SetSubclasses()
	fmt.Println("Seeded subclasses")

	fmt.Println("Seeding campaigns")
	s.SetCampaigns()
	fmt.Println("Seeded campaigns")

	fmt.Println("Seeding campaign session notes")
	s.SetCampaignSessionNotes()
	fmt.Println("Seeded campaign session notes")

	fmt.Println("Seeding characters")
	s.SetCharacters()
	fmt.Println("Seeded characters")

	fmt.Println("Seeded character profile pictures")
	s.SetCharacterProfilePictures()
	fmt.Println("Seeded character profile pictures")
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
