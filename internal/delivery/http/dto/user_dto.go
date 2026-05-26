package dto

type CreateUserRequest struct {
	Name  string `json:"name" validate:"required,max=50"`
	Email string `json:"email" validate:"required"`
}

type UpdateUserRequest struct {
	Name        *string `json:"name"         validate:"omitempty,min=3"`
	Email       *string `json:"email"        validate:"omitempty,email"`
	IsActive    *bool   `json:"is_active"    validate:"omitempty"`
	KtpNo       *string `json:"ktp_no"       validate:"omitempty,min=16,max=16"`
	Address     *string `json:"address"      validate:"omitempty,min=3"`
	PhoneNumber *string `json:"phone_number" validate:"omitempty,min=10,max=15"`
}

type UserResponse struct {
	ID          int     `json:"id"`
	Name        *string `json:"name"`
	Email       *string `json:"email"`
	IsActive    *bool   `json:"is_active"`
	KtpNo       *string `json:"ktp_no"`
	Address     *string `json:"address"`
	PhoneNumber *string `json:"phone_number"`
	CreatedDate string  `json:"created_date"`
	UpdatedDate *string `json:"updated_date"`
}
