package requests

// RegisterUserRequest is the structure of register user request form
type RegisterUserRequest struct {
	PhoneNumber        string    `json:"phone_number" validate:"uniqueInDB=users:phone_number,required_without=UserName Password ConfirmPassword,len=11"`
	UserName           string    `json:"username" validate:"required_without=PhoneNumber"`
	Password           string    `json:"password" validate:"required_without=PhoneNumber"`
	ConfirmPassword    string    `json:"confirm_password" validate:"required_without=PhoneNumber"`
	FirstName          string    `json:"first_name"`
	LastName           string    `json:"last_name"`
	Email              string    `json:"email"`
}