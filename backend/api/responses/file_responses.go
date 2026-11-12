package responses

import (
	"lorekeeper/db/models"
	"time"
)

type FileResponse struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`

	Author *SlimUserResponse `json:"author"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewFileResponse(file *models.File) *FileResponse {
	res := &FileResponse{
		ID:       file.ID,
		Name:     file.Name,
		Location: file.Location,

		CreatedAt: file.CreatedAt,
		UpdatedAt: file.UpdatedAt,
	}

	if file.Author != nil {
		res.Author = NewSlimUserResponse(file.Author)
	}

	return res
}

func NewFileResponses(files []*models.File) []FileResponse {
	var res []FileResponse

	for _, file := range files {
		res = append(res, *NewFileResponse(file))
	}

	return res
}
