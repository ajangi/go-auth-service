package translate

// ValidationField is the structure of a single validation field
type ValidationField struct {
	Fa string `json:"fa"`
	En string `json:"en"`
}

// ValidationFieldMessage is the structure of a single validation field message
type ValidationFieldMessage struct {
	Fa string `json:"fa"`
	En string `json:"en"`
}

// ValidationFields is the list of fields and translations!
type ValidationFields map[string]ValidationField

// Fields is the list of all fields and their translations
var Fields = ValidationFields{
	"PhoneNumber": ValidationField{Fa: "موبایل", En: "Phone Number"},
	"UserName": ValidationField{Fa: "نام کاربری", En: "Username"},
	"Password": ValidationField{Fa: "رمز عبور", En: "Password"},
	"ConfirmPassword": ValidationField{Fa: "تکرار رمز عبور", En: "Confirm Password"},
}
