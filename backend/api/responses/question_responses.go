package responses

import "lorekeeper/services"

type QuestionResponse struct {
	Answer  string                   `json:"answer"`
	Sources []QuestionSourceResponse `json:"sources"`
}

type QuestionSourceResponse struct {
	FileName   string  `json:"file_name"`
	PageNumber int     `json:"page_number"`
	Score      float64 `json:"score"`
}

func NewQuestionResponse(answer *services.QuestionAnswer) *QuestionResponse {
	if answer == nil {
		return &QuestionResponse{
			Answer: "There was no relevant information found.",
		}
	}

	res := &QuestionResponse{
		Answer: answer.Answer,
	}

	if len(answer.Sources) > 0 {
		res.Sources = NewQuestionSourceResponses(answer.Sources)
	}

	return res
}

func NewQuestionSourceResponse(source *services.SearchResult) QuestionSourceResponse {
	return QuestionSourceResponse{
		FileName:   source.Chunk.CampaignFile.File.Name,
		PageNumber: source.Chunk.PageNumber,
		Score:      source.Score,
	}
}

func NewQuestionSourceResponses(sources []*services.SearchResult) []QuestionSourceResponse {
	var res []QuestionSourceResponse

	for _, source := range sources {
		res = append(res, NewQuestionSourceResponse(source))
	}

	return res
}
