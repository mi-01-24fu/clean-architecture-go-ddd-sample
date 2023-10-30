package repository

import (
	"clean-architecture-go-ddd-sample/domain/model"
	"context"
)

type UserRepository interface {
	Exists(ctx context.Context, email model.MailAddress) (bool, error)
	Create(ctx context.Context, user model.User) error
	Select(ctx context.Context, email model.MailAddress) (*model.User, error)
}
