package requests

type CampaignSearchRequest struct {
	Question string `json:"question" validate:"required,max=200"`
}
