package repository

import (
	"clean-architecture-go-ddd-sample/domain/model"
	"context"
)

type userRepository struct{}

func NewUserRepository() *userRepository {
	return &userRepository{}
}

func (u userRepository) GetByEmail(ctx context.Context, email model.MailAddress) (*model.User, error) {

}
