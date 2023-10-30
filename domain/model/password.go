package model

import (
	"errors"
	"fmt"
	"regexp"
)

type Password struct {
	value string
}

func NewPasswordFromPlaintext(value string) (*Password, error) {
	p := &Password{value: value}

	if err := p.isValid(); err != nil {
		return nil, err
	}

	return p, nil
}

func NewPasswordFromHash(value string) *Password {
	return &Password{value: value}
}

func (p *Password) isValid() error {
	if !p.checkLength() {
		return errors.New("password should be between 6 and 15 characters")
	}
	if !p.includeLowercase() {
		return errors.New("password lower case not contains")
	}
	if !p.includeUppercase() {
		return errors.New("password upper case not contains")
	}
	if !p.includeNumeric() {
		return errors.New("password number not contains")
	}
	if !p.includeSymbol() {
		return errors.New("password special symbol not contains")
	}
	return nil
}

// 文字数チェック
func (p *Password) checkLength() bool {
	fmt.Println(p.value)
	return len([]rune(p.value)) >= 6 && len([]rune(p.value)) <= 15
}

// 小文字が含まれるかどうか
func (p *Password) includeLowercase() bool {
	return p.checkRegexp("[a-z]")
}

// 大文字が含まれるかどうか
func (p *Password) includeUppercase() bool {
	return p.checkRegexp("[A-Z]")
}

// 数値が含まれるかどうか
func (p *Password) includeNumeric() bool {
	return p.checkRegexp("[0-9]")
}

// 特殊記号が含まれるかどうか
func (p *Password) includeSymbol() bool {
	availableChar := p.checkRegexp(`^[0-9a-zA-Z\-^$*.@%]+$`)
	checkIsSymbol := p.checkRegexp(`[\-^$*.@%]`)

	return availableChar && checkIsSymbol
}

// 正規表現共通関数
func (p *Password) checkRegexp(reg string) bool {
	return regexp.MustCompile(reg).MatchString(p.value)
}

func (p *Password) String() string {
	return p.value
}
