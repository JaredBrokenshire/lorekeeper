package services

import (
	"errors"
	"github.com/ledongthuc/pdf"
	"lorekeeper/db/models"
	"lorekeeper/db/repositories"
	"lorekeeper/pkg/document"
	"path/filepath"
	"strings"
)

type DocumentServiceInterface interface {
	Extract(path string) ([]document.Page, error)
	Chunk(campaignFileID uint, pages []document.Page) ([]*models.DocumentChunk, error)
}

var (
	ErrDocumentEmptyPath = errors.New("document path is empty")
	ErrDocumentOpen      = errors.New("error opening document")
	ErrDocumentExtract   = errors.New("error extracting page")
	ErrDocumentChunk     = errors.New("error chunking document")
	ErrDocumentSaveChunk = errors.New("error saving document chunk")
	ErrDocumentPageRead  = errors.New("error reading document page")
)

type DocumentService struct {
	extractor document.ExtractorInterface

	chunkRepo *repositories.DocumentChunkRepository
}

func NewDocumentService(chunkRepo *repositories.DocumentChunkRepository) *DocumentService {
	return &DocumentService{
		extractor: document.NewExtractor(),
		chunkRepo: chunkRepo,
	}
}

func (s *DocumentService) Extract(path string) ([]document.Page, error) {
	if strings.TrimSpace(path) == "" {
		return nil, ErrDocumentEmptyPath
	}

	file, reader, err := pdf.Open(filepath.Join("../../assets", path))
	if err != nil {
		return nil, ErrDocumentOpen
	}
	defer file.Close()

	pages, err := s.extractor.Extract(reader)
	if err != nil {
		return nil, ErrDocumentExtract
	}

	return pages, nil
}

func (s *DocumentService) Chunk(campaignFileID uint, pages []document.Page) ([]*models.DocumentChunk, error) {
	chunks, err := s.extractor.Chunk(pages)
	if err != nil {
		return nil, ErrDocumentChunk
	}

	var res []*models.DocumentChunk
	for _, chunk := range chunks {
		c := &models.DocumentChunk{
			PageNumber:     chunk.PageNumber,
			Index:          chunk.Index,
			Content:        chunk.Content,
			CampaignFileID: campaignFileID,
		}
		err := s.chunkRepo.Create(&c)
		if err != nil {
			return nil, ErrDocumentSaveChunk
		}

		res = append(res, c)
	}

	return res, nil
}
