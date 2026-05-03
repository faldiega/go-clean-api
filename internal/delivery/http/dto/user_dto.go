package dto

type CreateUserRequest struct {
	Name  string `json:"name" validate:"required,max=50"`
	Email string `json:"email" validate:"required"`
}

type UpdateUserRequest struct {
	Name     string `json:"name" validate:"required,max=50"`
	Email    string `json:"email" validate:"required"`
	IsActive bool   `json:"is_active"`
}

type UserResponse struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Email       string  `json:"email"`
	IsActive    bool    `json:"is_active"`
	CreatedDate string  `json:"created_date"`
	UpdatedDate *string `json:"updated_date"`
}
