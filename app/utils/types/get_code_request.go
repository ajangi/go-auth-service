package types

// GetCodeRequest is the object of getting code request object
type GetCodeRequest struct {
	PhoneNumber string `json:"phone_number" form:"phone_number" query:"phone_number" validate:"required,email"`
}