package seeders

import (
	"bytes"
	"fmt"
	"github.com/jinzhu/gorm"
	"lorekeeper/db/models"
	"lorekeeper/pkg/logging"
	"os"
	"path/filepath"
	"strings"

	"github.com/labstack/gommon/random"
)

const campaignSessionNotesSeedDir = "db/seeders/assets/campaigns"

var campaignSessionNoteAuthors = map[string]uint{
	"Kael_Drakeshield":  1,
	"Faelan_Haversham":  2,
	"Niko_Nightstriker": 3,
	"Idris":             4,
	"Karma":             5,
}

func (s *Seeder) SetCampaignSessionNotes() {
	if s.FileStore == nil {
		logging.Logger.LogWarn().Msg(
			"seeder: skipping SetCampaignSessionNotes — no file store configured",
		)
		return
	}

	var campaigns []models.Campaign
	if err := s.DB.Preload("Files").Find(&campaigns).Error; err != nil {
		logging.Logger.LogWarn().Msg(fmt.Sprintf("Seeder: SetCampaignSessionNotes could not load campaigns: %v", err))
		return
	}

	assetDir := resolveDir(campaignSessionNotesSeedDir)
	if assetDir == "" {
		logging.Logger.LogWarn().Msg(fmt.Sprintf("Seeder: error finding asset dir %q - skipping", campaignSessionNotesSeedDir))
		return
	}

	for _, campaign := range campaigns {
		if len(campaign.Files) > 0 {
			continue
		}

		campaignDir := filepath.Join(assetDir, fmt.Sprintf("%v", campaign.ID), "session-notes")

		entries, err := os.ReadDir(campaignDir)
		if err != nil {
			logging.Logger.LogInfo().Msg(fmt.Sprintf("Seeder: no session notes found for campaign %v - skipping", campaign.ID))
			continue
		}

		storagePath := fmt.Sprintf("/campaigns/%v/session-notes", campaign.ID)

		for _, entry := range entries {
			if entry.IsDir() {
				continue
			}

			if filepath.Ext(entry.Name()) != ".pdf" {
				continue
			}

			srcPath := filepath.Join(campaignDir, entry.Name())

			data, err := os.ReadFile(srcPath)
			if err != nil {
				logging.Logger.LogWarn().Msg(fmt.Sprintf("Seeder: error reading session note %q: %v", srcPath, err))
				continue
			}

			newFileName := fmt.Sprintf("%v.pdf", random.String(16))

			err = s.FileStore.Save(storagePath, bytes.NewReader(data), newFileName)
			if err != nil {
				logging.Logger.LogWarn().Msg(fmt.Sprintf("Seeder: error saving session note %q for campaign %v: %v", entry.Name(), campaign.ID, err))
				continue
			}

			authorID := uint(0)
			for key, value := range campaignSessionNoteAuthors {
				if strings.Contains(entry.Name(), key) {
					authorID = value
					break
				}
			}
			if authorID == 0 {
				logging.Logger.LogInfo().Msg(fmt.Sprintf("Seeder: no author configured for session note %q - skipping", entry.Name()))
				continue
			}

			file := models.File{
				Name:     newFileName,
				Location: fmt.Sprintf("%v/%v", storagePath, newFileName),
				AuthorID: authorID,
			}

			// Create the file and associate it with the campaign.
			err = s.DB.Transaction(func(tx *gorm.DB) error {
				if err := tx.Create(&file).Error; err != nil {
					return err
				}

				return tx.Model(&campaign).Association("Files").Append(&file).Error
			})

			if err != nil {
				logging.Logger.LogWarn().Msg(fmt.Sprintf("Seeder: error creating session note %q for campaign %v: %v", entry.Name(), campaign.ID, err))
				continue
			}
		}
	}
}
