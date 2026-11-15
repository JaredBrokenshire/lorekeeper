package mocks

import (
	"lorekeeper/db/models"
	"lorekeeper/pkg/document"
)

type MockDocumentService struct {
	ExtractCalls    []string
	ExtractResponse []document.Page
	ExtractError    error

	ChunkCalls []struct {
		CampaignFileID uint
		Pages          []document.Page
	}
	ChunkResponse []*models.DocumentChunk
	ChunkError    error
}

func NewMockDocumentService() *MockDocumentService {
	m := &MockDocumentService{}
	m.Reset()
	return m
}

func (m *MockDocumentService) Extract(path string) ([]document.Page, error) {
	m.ExtractCalls = append(m.ExtractCalls, path)
	return m.ExtractResponse, m.ExtractError
}

func (m *MockDocumentService) Chunk(campaignFileID uint, pages []document.Page) ([]*models.DocumentChunk, error) {
	m.ChunkCalls = append(m.ChunkCalls, struct {
		CampaignFileID uint
		Pages          []document.Page
	}{campaignFileID, pages})
	return m.ChunkResponse, m.ChunkError
}

func (m *MockDocumentService) Reset() {
	m.ExtractCalls = []string{}
	m.ExtractResponse = []document.Page{}
	m.ExtractError = nil

	m.ChunkCalls = []struct {
		CampaignFileID uint
		Pages          []document.Page
	}{}
	m.ChunkResponse = []*models.DocumentChunk{}
	m.ChunkError = nil
}
