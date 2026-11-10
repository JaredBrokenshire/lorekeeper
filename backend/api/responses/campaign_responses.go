package responses

import (
	"lorekeeper/db/models"
	"time"
)

type CampaignResponse struct {
	*SlimCampaignResponse
}

type SlimCampaignResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`

	DM *SlimUserResponse `json:"dm"`

	Characters []SlimCharacterResponse `json:"characters"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewSlimCampaignResponse(campaign *models.Campaign) *SlimCampaignResponse {
	res := &SlimCampaignResponse{
		ID:        campaign.ID,
		Name:      campaign.Name,
		CreatedAt: campaign.CreatedAt,
		UpdatedAt: campaign.UpdatedAt,
	}

	if campaign.DungeonMaster != nil {
		res.DM = NewSlimUserResponse(campaign.DungeonMaster)
	}

	if len(campaign.Characters) > 0 {
		res.Characters = NewCharacterResponses(campaign.Characters)
	}

	return res
}

func NewCampaignResponses(campaigns []*models.Campaign) []SlimCampaignResponse {
	var res []SlimCampaignResponse

	for _, campaign := range campaigns {
		res = append(res, *NewSlimCampaignResponse(campaign))
	}

	return res
}

type PaginatedCampaignResponse struct {
	Data []SlimCampaignResponse `json:"data"`
	Meta ResponseMeta           `json:"meta"`
}

func NewPaginatedCampaignResponse(campaigns []*models.Campaign, count, page, pageSize int) PaginatedCampaignResponse {
	return PaginatedCampaignResponse{
		Data: NewCampaignResponses(campaigns),
		Meta: ResponseMeta{
			TotalCount: count,
			Page:       page,
			PageSize:   pageSize,
		},
	}
}
