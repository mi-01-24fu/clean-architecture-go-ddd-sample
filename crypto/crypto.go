package crypto

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type Encrypt interface {
	PasswordEncrypt(password string) (string, error)
}

type encrypt struct{}

func NewEncrypt() *encrypt {
	return &encrypt{}
}

func (e encrypt) PasswordEncrypt(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("error has occurred during password encryption received_error: %v", err)
	}
	return string(hash), nil
}
