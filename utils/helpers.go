package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type Helper struct {
}

func (h *Helper) RemoveDashesFromString(uuidStr string) string {
	formatted := ""
	for _, char := range uuidStr {
		if char != '-' {
			formatted += string(char)
		}
	}
	return formatted
}

func (h *Helper) EncryptPassword(password string) (hPass string) {
	pass := []byte(password)
	result, _ := bcrypt.GenerateFromPassword(pass, bcrypt.DefaultCost)

	hassedPass := string(result)
	return hassedPass
}

func (h *Helper) ComparePassowrds(hashPassowrd string, passowrd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassowrd), []byte(passowrd))
	if err != nil {
		fmt.Println(err)
		return false
	} else {
		return true
	}
}
