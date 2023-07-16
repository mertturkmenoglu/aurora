package utils

import (
	"aurora/db"
	"aurora/db/models"
	"github.com/google/uuid"
	"strings"
	"unicode"
)

func DoesUserExist(email string) (bool, error) {
	var auth models.Auth
	res := db.Client.First(&auth, "email = ?", email)

	if res.Error != nil {
		if strings.Contains(res.Error.Error(), "record not found") {
			return false, nil
		}
		return false, res.Error
	}

	return true, nil
}

func CustomPasswordCheck(password string) bool {
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

func GenerateRandomShortId() string {
	sid := uuid.NewString()
	return strings.Split(sid, "-")[0]
}
