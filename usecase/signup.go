package usecase

import (
	"clean-architecture-go-ddd-sample/crypto"
	"clean-architecture-go-ddd-sample/domain/model"
	"clean-architecture-go-ddd-sample/domain/repository"
	"context"
	"fmt"
)

type SignupUsecase interface {
	Signup(ctx context.Context, req SignupRequest) (*model.User, error)
}

type signupUsecase struct {
	userRepository repository.UserRepository
	encrypt        crypto.Encrypt
}

func NewSignupUsecase(userRepository repository.UserRepository, encrypt crypto.Encrypt) *signupUsecase {
	return &signupUsecase{userRepository: userRepository, encrypt: encrypt}
}

type SignupRequest struct {
	UserName    string `json:"user_name"`
	MailAddress string `json:"mail_address"`
	Password    string `json:"password"`
}

type SignupResponse struct {
	ID          int    `json:"id"`
	UserName    string `json:"user_name"`
	MailAddress string `json:"mail_address"`
}

type validationResult struct {
	userName    *model.UserName
	mailAddress *model.MailAddress
	password    *model.Password
}

func (s signupUsecase) Signup(ctx context.Context, req SignupRequest) (*model.User, error) {
	// バリデーションチェック
	input, err := validateInput(req)
	if err != nil {
		return nil, err
	}

	// メールアドレスの重複確認
	exists, err := s.userRepository.Exists(ctx, *input.mailAddress)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, fmt.Errorf("exists email already: %v", *input.mailAddress)
	}

	// パスワードハッシュ化
	hash, err := s.encrypt.PasswordEncrypt(input.password.String())
	if err != nil {
		return nil, err
	}

	// 新規登録処理
	newUser, err := s.registerAndFetchUser(
		ctx,
		model.NewSignupUser(*input.userName, *input.mailAddress, *(model.NewPasswordFromHash(hash))))
	if err != nil {
		return nil, err
	}

	return newUser, nil
}

func validateInput(req SignupRequest) (*validationResult, error) {
	userName, err := model.NewUserName(req.UserName)
	if err != nil {
		return nil, err
	}

	mailAddress, err := model.NewMailAddress(req.MailAddress)
	if err != nil {
		return nil, err
	}

	password, err := model.NewPasswordFromPlaintext(req.Password)
	if err != nil {
		return nil, err
	}

	return &validationResult{
		userName:    userName,
		mailAddress: mailAddress,
		password:    password,
	}, nil
}

func (s signupUsecase) registerAndFetchUser(ctx context.Context, inputUser *model.User) (*model.User, error) {
	if err := s.userRepository.Create(ctx, *inputUser); err != nil {
		return nil, err
	}

	createdUser, err := s.userRepository.Select(ctx, inputUser.MailAddress)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}
