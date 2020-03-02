package types

import (
	"github.com/ajangi/gAuthService/app/utils/translate"
)
// ErrorsList is the list of errors that validator should return
type ErrorsList map[string]translate.ValidationFieldMessage
