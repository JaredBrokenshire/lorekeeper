package responses

import "lorekeeper/db/models"

type ClassResponse struct {
	*SlimClassResponse

	Subclasses []SlimSubclassResponse `json:"subclasses"`
}

type SlimClassResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`

	Logo *FileResponse `json:"logo"`
}

func NewSlimClassResponse(class *models.Class) *SlimClassResponse {
	res := &SlimClassResponse{
		ID:   class.ID,
		Name: class.Name,
	}

	if class.Logo != nil {
		res.Logo = NewFileResponse(class.Logo)
	}

	return res
}

func NewClassResponses(classes []*models.Class) []SlimClassResponse {
	var res []SlimClassResponse

	for _, class := range classes {
		res = append(res, *NewSlimClassResponse(class))
	}

	return res
}
