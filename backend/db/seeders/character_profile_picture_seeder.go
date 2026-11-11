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

const characterProfilePictureSeedDir = "db/seeders/assets/characters/profile_pictures"

func (s *Seeder) SetCharacterProfilePictures() {
	if s.FileStore == nil {
		logging.Logger.LogWarn().Msg("seeder: skipping SetCharacterProfilePictures — no file store configured")
		return
	}

	var characters []models.Character
	if err := s.DB.Find(&characters).Error; err != nil {
		logging.Logger.LogWarn().Msg(fmt.Sprintf("Seeder: SetCharacterProfilePictures could not load characters: %v", err.Error()))
		return
	}

	// Resolve the seed asset directory, walking up from cwd if needed.
	assetDir := resolveDir(characterProfilePictureSeedDir)
	if assetDir == "" {
		logging.Logger.LogWarn().Msg(fmt.Sprintf("Seeder: error finding asset dir %q - skipping", characterProfilePictureSeedDir))
		return
	}

	for _, character := range characters {
		// Skip if this class already has a logo.
		if character.ProfilePictureID != 0 {
			continue
		}

		filename := fmt.Sprintf("%v.png", character.ID)
		srcPath := filepath.Join(assetDir, filename)

		data, err := os.ReadFile(srcPath)
		if err != nil {
			// No seed image for this class — not an error, just skip.
			logging.Logger.LogInfo().Msg(fmt.Sprintf("Seeder: no seed profile picture for character %q (%v) - skipping", character.Name, srcPath))
			continue
		}

		storagePath := "/characters/profile_pictures"
		newFileName := fmt.Sprintf("%v.png", random.String(16))
		if err := s.FileStore.Save(storagePath, bytes.NewReader(data), newFileName); err != nil {
			logging.Logger.LogWarn().Msg(fmt.Sprintf("Seeder: error saving profile picture for character %q: %v", character.Name, err.Error()))
			continue
		}

		file := models.File{
			Name:     newFileName,
			Location: fmt.Sprintf("%v/%v", storagePath, newFileName),
		}
		if err := s.DB.Create(&file).Error; err != nil {
			logging.Logger.LogWarn().Msg(fmt.Sprintf("Seeder: error creating file record for character %q: %v", character.Name, err.Error()))
			continue
		}

		if err := s.DB.Model(&character).Update("profile_picture_id", file.ID).Error; err != nil {
			logging.Logger.LogWarn().Msg(fmt.Sprintf("Seeder: error updating profile_picture_id for character %q: %v", character.Name, err.Error()))
		}
	}
}
