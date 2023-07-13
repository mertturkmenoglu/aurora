package auth

import (
	"aurora/services/aws/models"
	"github.com/google/uuid"
	"strings"
	"unicode"
)

func doesUserExist(email string) (bool, error) {
	var auth models.Auth
	authResult, err := auth.GetByEmail(email)

	if err != nil || (models.Auth{}) == *authResult {
		return false, err
	}

	return true, nil
}

func customPasswordCheck(password string) bool {
	hasUpper := false
	hasLower := false

	for _, char := range password {
		if unicode.IsUpper(char) {
			hasUpper = true
		}
		if unicode.IsLower(char) {
			hasLower = true
		}
		if hasLower && hasUpper {
			break
		}
	}

	return hasUpper && hasLower
}

func generateRandomShortId() string {
	sid := uuid.NewString()
	return strings.Split(sid, "-")[0]
}
