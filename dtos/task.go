package dtos

type CreateTaskRequest struct {
	Title       string `json:"title" validate:"required,min=3,max=100"`
	Description string `json:"description" validate:"omitempty,min=3,max=500"`
	ParentID    *uint  `json:"parent_id" validate:"omitempty"`
}

type UpdateTaskRequest struct {
	Title       string `json:"title" validate:"omitempty,min=3,max=100"`
	Description string `json:"description" validate:"omitempty,min=3,max=500"`
}

type FilterTaskRequest struct {
	SearchQuery string `json:"search_query" validate:"omitempty"`
	Status      string `json:"status" validate:"omitempty,oneof=Pending InProgress OnHold Completed Cancelled"`
	ParentID    *uint  `json:"parent_id" validate:"omitempty"`
}

type GetAiSuggestion struct {
	Title       string `json:"title" validate:"required,min=3,max=100"`
	Description string `json:"description" validate:"required,min=3,max=500"`
	Context     string `json:"context" validate:"omitempty,min=3,max=500"`
}
