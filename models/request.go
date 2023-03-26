package models

type CreateUserRequest struct {
	Email string `json:"email" validate:"email,required"`
}
