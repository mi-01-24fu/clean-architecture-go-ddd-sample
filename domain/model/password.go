package model

import (
	"errors"
	"regexp"
)

type Password struct {
	value string
}

func NewSignupPassword(value string) *Password {
	return &Password{value: value}
}

func NewPassword(value string) (*Password, error) {
	p := &Password{value: value}

	err := p.isValid()
	if err != nil {
		return nil, err
	}

	return &Password{value: value}, nil
}

func (p *Password) isValid() error {
	if !includeLowercase(p.value) {
		return errors.New("password lower case not contains")
	}
	if !includeUppercase(p.value) {
		return errors.New("password upper case not contains")
	}
	if !includeNumeric(p.value) {
		return errors.New("password number not contains")
	}
	if !includeSymbol(p.value) {
		return errors.New("password special symbol not contains")
	}
	return nil
}

// 小文字が含まれるかどうか
func includeLowercase(value string) bool {
	return checkRegexp("[a-z]", value)
}

// 大文字が含まれるかどうか
func includeUppercase(value string) bool {
	return checkRegexp("[A-Z]", value)
}

// 数値が含まれるかどうか
func includeNumeric(value string) bool {
	return checkRegexp("[0-9]", value)
}

// 特殊記号が含まれるかどうか
func includeSymbol(value string) bool {
	availableChar := checkRegexp(`^[0-9a-zA-Z\-^$*.@]+$`, value)
	checkIsSymbol := checkRegexp(`[\-^$*.@]`, value)

	return availableChar && checkIsSymbol
}

// 正規表現共通関数
func checkRegexp(reg, value string) bool {
	r := regexp.MustCompile(reg).Match([]byte(value))
	return r
}

func (p *Password) String() string {
	return p.value
}
