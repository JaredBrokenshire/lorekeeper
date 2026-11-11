package seeders

import (
	"bytes"
	"fmt"
	"github.com/labstack/gommon/random"
	"lorekeeper/db/models"
	"lorekeeper/pkg/logging"
	"os"
	"path/filepath"
)

const classLogoSeedDir = "db/seeders/assets/classes/logos"

func (s *Seeder) SetClassLogos() {
	if s.FileStore == nil {
		logging.Logger.LogWarn().Msg("seeder: skipping SetClassLogos — no file store configured")
		return
	}

	var classes []models.Class
	if err := s.DB.Find(&classes).Error; err != nil {
		logging.Logger.LogWarn().Msg(fmt.Sprintf("Seeder: SetClassLogos could not load classes: %v", err.Error()))
		return
	}

	// Resolve the seed asset directory, walking up from cwd if needed.
	assetDir := resolveDir(classLogoSeedDir)
	if assetDir == "" {
		logging.Logger.LogWarn().Msg(fmt.Sprintf("Seeder: error finding asset dir %q - skipping", classLogoSeedDir))
		return
	}

	for _, class := range classes {
		// Skip if this class already has a logo.
		if class.LogoID != 0 {
			continue
		}

		filename := fmt.Sprintf("%v.png", class.ID)
		srcPath := filepath.Join(assetDir, filename)

		data, err := os.ReadFile(srcPath)
		if err != nil {
			// No seed image for this class — not an error, just skip.
			logging.Logger.LogInfo().Msg(fmt.Sprintf("Seeder: no seed logo for class %q (%v) - skipping", class.Name, srcPath))
			continue
		}

		storagePath := "/classes/logos"
		newFileName := fmt.Sprintf("%v.png", random.String(16))
		if err := s.FileStore.Save(storagePath, bytes.NewReader(data), newFileName); err != nil {
			logging.Logger.LogWarn().Msg(fmt.Sprintf("Seeder: error saving logo for class %q: %v", class.Name, err.Error()))
			continue
		}

		file := models.File{
			Name:     newFileName,
			Location: fmt.Sprintf("%v/%v", storagePath, newFileName),
		}
		if err := s.DB.Create(&file).Error; err != nil {
			logging.Logger.LogWarn().Msg(fmt.Sprintf("Seeder: error creating file record for class %q: %v", class.Name, err.Error()))
			continue
		}

		if err := s.DB.Model(&class).Update("logo_id", file.ID).Error; err != nil {
			logging.Logger.LogWarn().Msg(fmt.Sprintf("Seeder: error updating logo_file_id for class %q: %v", class.Name, err.Error()))
		}
	}
}
