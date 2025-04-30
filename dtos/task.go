package dtos

type CreateTaskRequest struct {
	Title       string `json:"title" validate:"required,min=3,max=100"`
	Description string `json:"description" validate:"required,min=3,max=500"`
	ParentID    *uint  `json:"parent_id" validate:"omitempty"`
	AssignToID  *uint  `json:"assign_to_id" validate:"required"`
	ReportToID  *uint  `json:"report_to_id" validate:"omitempty"`
	Priority    string `json:"priority" validate:"omitempty,oneof=Low Medium High"`
}

type CreateManyTasksRequest struct {
	Tasks []CreateTaskRequest `json:"tasks" validate:"required"`
}

type UpdateTaskRequest struct {
	Title       string `json:"title" validate:"omitempty,min=3,max=100"`
	Description string `json:"description" validate:"omitempty,min=3,max=500"`
}

type UpdateTaskPriorityRequest struct {
	Priority string `json:"priority" validate:"required,oneof=Low Medium High"`
}

type UpdateTaskStatusRequest struct {
	Status string `json:"status" validate:"required,oneof=Pending InProgress OnHold Completed Cancelled"`
}

type UpdateTaskAssignToRequest struct {
	AssignToID *uint `json:"assign_to_id" validate:"required"`
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
