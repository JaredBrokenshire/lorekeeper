package responses

import "lorekeeper/db/models"

type SubclassResponse struct {
	*SlimSubclassResponse
}

type SlimSubclassResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func NewSlimSubclassResponse(subclass *models.Subclass) *SlimSubclassResponse {
	return &SlimSubclassResponse{
		ID:   subclass.ID,
		Name: subclass.Name,
	}
}

func NewSubclassResponses(subclasses []*models.Subclass) []SlimSubclassResponse {
	var res []SlimSubclassResponse

	for _, subclass := range subclasses {
		res = append(res, *NewSlimSubclassResponse(subclass))
	}

	return res
}
