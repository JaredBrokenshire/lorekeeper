package responses

import "lorekeeper/db/models"

type CharacterResponse struct {
	*SlimCharacterResponse
}

type SlimCharacterResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Level int    `json:"level"`

	Race       SlimRaceResponse       `json:"race"`
	Classes    []SlimClassResponse    `json:"classes"`
	Subclasses []SlimSubclassResponse `json:"subclasses"`
}

func NewSlimCharacterResponse(character *models.Character) *SlimCharacterResponse {
	res := &SlimCharacterResponse{
		ID:    character.ID,
		Name:  character.Name,
		Level: character.Level,
	}

	if character.Race.ID != 0 {
		res.Race = *NewSlimRaceResponse(character.Race)
	}

	if len(character.Classes) > 0 {
		res.Classes = NewClassResponses(character.Classes)
	}

	if len(character.Subclasses) > 0 {
		res.Subclasses = NewSubclassResponses(character.Subclasses)
	}

	return res
}

func NewCharacterResponses(characters []*models.Character) []SlimCharacterResponse {
	var res []SlimCharacterResponse

	for _, character := range characters {
		res = append(res, *NewSlimCharacterResponse(character))
	}

	return res
}

type PaginatedCharacterResponse struct {
	Data []SlimCharacterResponse `json:"data"`
	Meta ResponseMeta            `json:"meta"`
}

func NewPaginatedCharacterResponse(characters []*models.Character, count, page, pageSize int) PaginatedCharacterResponse {
	return PaginatedCharacterResponse{
		Data: NewCharacterResponses(characters),
		Meta: ResponseMeta{
			TotalCount: count,
			Page:       page,
			PageSize:   pageSize,
		},
	}
}
