package responses

import "lorekeeper/services"

type SearchResponse struct {
	ID             uint          `json:"id"`
	Content        string        `json:"content"`
	Score          float64       `json:"score"`
	PageNumber     int           `json:"page_number"`
	ChunkIndex     int           `json:"chunk_index"`
	CampaignFileID uint          `json:"campaign_file_id"`
	File           *FileResponse `json:"file,omitempty"`
}

func NewSearchResponse(result *services.SearchResult) *SearchResponse {
	res := &SearchResponse{
		ID:             result.Chunk.ID,
		Content:        result.Chunk.Content,
		Score:          result.Score,
		PageNumber:     result.Chunk.PageNumber,
		ChunkIndex:     result.Chunk.Index,
		CampaignFileID: result.Chunk.CampaignFileID,
	}

	if result.Chunk.CampaignFile != nil && result.Chunk.CampaignFile.File != nil {
		res.File = NewFileResponse(result.Chunk.CampaignFile.File)
	}

	return res
}

func NewSearchResponses(results []*services.SearchResult) []SearchResponse {
	var res []SearchResponse

	for _, result := range results {
		res = append(res, *NewSearchResponse(result))
	}

	return res
}
