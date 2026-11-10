package responses

import "lorekeeper/db/models"

type RaceResponse struct {
	*SlimRaceResponse
}

type SlimRaceResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func NewSlimRaceResponse(race *models.Race) *SlimRaceResponse {
	return &SlimRaceResponse{
		ID:   race.ID,
		Name: race.Name,
	}
}
