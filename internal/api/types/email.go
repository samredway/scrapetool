package types

type EmailRequest struct {
	Email string `json:"email" validate:"required,email"`
}
