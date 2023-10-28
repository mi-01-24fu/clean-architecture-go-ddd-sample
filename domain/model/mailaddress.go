package model

import (
	"errors"
	"regexp"
)

type MailAddress struct {
	value string
}

const regularMail = "^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"

func NewMailAddress(value string) (*MailAddress, error) {
	regexPattern := regexp.MustCompile(regularMail)

	if value == "" {
		return nil, errors.New("not input email address")
	}

	if !regexPattern.MatchString(value) {
		return nil, errors.New("email address format is incorrect")
	}

	return &MailAddress{value: value}, nil
}
