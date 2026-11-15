package responses

import "lorekeeper/db/models"

type CampaignFileResponse struct {
	ID uint `json:"id"`

	File   *FileResponse           `json:"file"`
	Chunks []DocumentChunkResponse `json:"chunks"`
}

func NewCampaignFileResponse(campaignFile *models.CampaignFile) *CampaignFileResponse {
	res := &CampaignFileResponse{
		ID: campaignFile.ID,
	}

	if campaignFile.File != nil {
		res.File = NewFileResponse(campaignFile.File)
	}

	if len(campaignFile.Chunks) > 0 {
		res.Chunks = NewDocumentChunkResponses(campaignFile.Chunks)
	}

	return res
}

func NewCampaignFileResponses(campaignFiles []*models.CampaignFile) []CampaignFileResponse {
	var res []CampaignFileResponse

	for _, campaignFile := range campaignFiles {
		res = append(res, *NewCampaignFileResponse(campaignFile))
	}

	return res
}
