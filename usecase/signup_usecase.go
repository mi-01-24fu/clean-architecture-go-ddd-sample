package usecase

import (
	"clean-architecture-go-ddd-sample/crypto"
	"clean-architecture-go-ddd-sample/domain/model"
	"context"
	"fmt"
)

type SignupUsecase interface {
	Signup(ctx context.Context, req SignupRequest) (*model.User, error)
}

type signupUsecase struct {
	userRepository model.UserRepository
	encrypt        crypto.Encrypt
}

func NewSignupUsecase(userRepository model.UserRepository, encrypt crypto.Encrypt) *signupUsecase {
	return &signupUsecase{userRepository: userRepository, encrypt: encrypt}
}

type SignupRequest struct {
	UserName    string
	MailAddress string
	Password    string
}

func (s signupUsecase) Signup(ctx context.Context, req SignupRequest) (*model.User, error) {
	userName, err := model.NewUserName(req.UserName)
	if err != nil {
		return nil, err
	}

	mailAddress, err := model.NewMailAddress(req.MailAddress)
	if err != nil {
		return nil, err
	}

	exists, err := s.userRepository.GetByEmail(ctx, *mailAddress)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, fmt.Errorf("exists email already: %v", *mailAddress)
	}

	password, err := model.NewPassword(req.Password)
	if err != nil {
		return nil, err
	}

	hash, err := s.encrypt.PasswordEncrypt(password.String())
	if err != nil {
		return nil, err
	}
	unregisteredUser := model.NewSignupUser(*userName, *mailAddress, *(model.NewSignupPassword(hash)))

	registeredUser, err := s.userRepository.Create(ctx, *unregisteredUser)
	if err != nil {
		return nil, err
	}

	return registeredUser, nil
}
