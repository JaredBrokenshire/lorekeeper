package seeders

import (
	"bytes"
	"fmt"
	"github.com/ledongthuc/pdf"
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
	if err := s.DB.Preload("Files").Preload("Files.File").Find(&campaigns).Error; err != nil {
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

			file := &models.File{
				Name:     newFileName,
				Location: fmt.Sprintf("%v/%v", storagePath, newFileName),
				AuthorID: authorID,
			}
			campaignFile := &models.CampaignFile{
				CampaignID: campaign.ID,
				File:       file,
			}
			err = s.DB.Create(campaignFile).Error
			if err != nil {
				logging.Logger.LogWarn().Msg(fmt.Sprintf("Seeder: error creating session note %q for campaign %v: %v", entry.Name(), campaign.ID, err))
				continue
			}

			osFile, reader, err := pdf.Open(filepath.Join("assets/files", file.Location))
			if err != nil {
				logging.Logger.LogWarn().Msg(fmt.Sprintf("Seeder: error opening file for campaign %v: %v", campaign.ID, err))
			}
			defer osFile.Close()

			pages, err := s.Extractor.Extract(reader)
			if err != nil {
				logging.Logger.LogWarn().Msg(fmt.Sprintf("Seeder: error extracting page content from session notes for campaign %v: %v", campaign.ID, err))
				continue
			}

			chunks, err := s.Extractor.Chunk(pages)
			if err != nil {
				logging.Logger.LogWarn().Msg(fmt.Sprintf("Seeder: error chunking page content from session notes for campaign %v: %v", campaign.ID, err))
				continue
			}

			for _, chunk := range chunks {
				var existing models.DocumentChunk
				s.DB.
					Where("campaign_file_id = ?", campaignFile.ID).
					Where("page_number = ?", chunk.PageNumber).
					Where("chunk_index = ?", chunk.Index).
					Where("content = ?", chunk.Content).
					First(&existing)

				if existing.ID != 0 {
					continue
				}

				documentChunk := &models.DocumentChunk{
					CampaignFileID: campaignFile.ID,
					PageNumber:     chunk.PageNumber,
					Index:          chunk.Index,
					Content:        chunk.Content,
				}
				err := s.DB.Create(&documentChunk).Error
				if err != nil {
					logging.Logger.LogWarn().Msg(fmt.Sprintf("Error creating document chunk with content `%v` for campaign `%v` in seeder", chunk.Content, campaign.ID))
					continue
				}
			}
		}
	}
}
