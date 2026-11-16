package services

import (
	"errors"
	"fmt"
	"strings"

	"lorekeeper/db/models"
	"lorekeeper/db/repositories"
	"lorekeeper/pkg/search"
)

var (
	ErrSearchRepository = errors.New("error retrieving campaign document chunks")
	ErrSearchNoResults  = errors.New("no results found for search query")
)

type SearchResult struct {
	Chunk *models.DocumentChunk
	Score float64
}

type SearchServiceInterface interface {
	Search(campaignID interface{}, userID uint, query string) ([]*SearchResult, error)
}

type SearchService struct {
	campaignRepo *repositories.CampaignRepository
	chunkRepo    *repositories.DocumentChunkRepository
}

func NewSearchService(campaignRepo *repositories.CampaignRepository, chunkRepo *repositories.DocumentChunkRepository) *SearchService {
	return &SearchService{
		campaignRepo: campaignRepo,
		chunkRepo:    chunkRepo,
	}
}

func (s *SearchService) Search(campaignID interface{}, userID uint, query string) ([]*SearchResult, error) {
	query = strings.TrimSpace(query)
	accessible, err := s.campaignRepo.ExistsForUser(campaignID, userID)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrSearchRepository, err)
	}
	if !accessible {
		return nil, ErrCampaignNotFound
	}

	chunks, err := s.chunkRepo.GetByCampaignID(campaignID)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrSearchRepository, err)
	}
	if len(chunks) == 0 {
		return nil, ErrSearchNoResults
	}

	documents := make([]search.Document, 0, len(chunks))
	chunksByID := make(map[uint]*models.DocumentChunk, len(chunks))
	for _, chunk := range chunks {
		if chunk == nil {
			continue
		}
		documents = append(documents, search.Document{
			ID:   chunk.ID,
			Text: chunk.Content,
		})
		chunksByID[chunk.ID] = chunk
	}

	ranked := search.RankBM25(documents, query)
	results := make([]*SearchResult, 0, len(ranked))
	for _, result := range ranked {
		if chunk, ok := chunksByID[result.ID]; ok {
			results = append(results, &SearchResult{
				Chunk: chunk,
				Score: result.Score,
			})
		}
	}

	return results, nil
}
