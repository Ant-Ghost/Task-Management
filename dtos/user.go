package dtos

type RegisterUserRequest struct {
	Username   string `json:"username" validate:"required,min=3,max=20"`
	Email      string `json:"email" validate:"required,email"`
	Password   string `json:"password" validate:"required,min=6,max=20"`
	SlackID    string `json:"slack_id" validate:"omitempty,min=6,max=20"`
	Experience int    `json:"experience" validate:"omitempty,min=0,max=100"`
}

type UpdateUserRequest struct {
	Username   string `json:"username" validate:"omitempty,min=3,max=20"`
	Email      string `json:"email" validate:"omitempty,email"`
	SlackID    string `json:"slack_id" validate:"omitempty,min=6,max=20"`
	Experience *int   `json:"experience" validate:"omitempty,min=0,max=100"`
}

type LoginUserRequest struct {
	Username string `json:"username" validate:"required,min=3,max=20"`
	Password string `json:"password" validate:"required,min=6,max=20"`
}
