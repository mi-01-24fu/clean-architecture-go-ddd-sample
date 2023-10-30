package repository

import (
	"clean-architecture-go-ddd-sample/domain/model"
	"clean-architecture-go-ddd-sample/infrastructure/repository/models"
	"context"
	"database/sql"
	"fmt"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"time"
)

type userRepository struct {
	db          *sql.DB
	userFactory model.UserFactorier
}

func NewUserRepository(db *sql.DB, userFactory model.UserFactorier) *userRepository {
	return &userRepository{
		db:          db,
		userFactory: userFactory,
	}
}

func (u userRepository) Exists(ctx context.Context, email model.MailAddress) (bool, error) {
	exists, err := models.Users(
		models.UserWhere.Mailaddress.EQ(email.String()),
	).Exists(ctx, u.db)
	if err != nil {
		return exists, fmt.Errorf("error occurred at existence confirmation process: %v", err)
	}
	if exists {
		return exists, fmt.Errorf("registered this email address already: %v", email.String())
	}
	return exists, nil
}

func (u userRepository) Create(ctx context.Context, user model.User) error {
	input := models.User{
		Username:     user.UserName.String(),
		Mailaddress:  user.MailAddress.String(),
		PasswordHash: user.Password.String(),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	err := input.Insert(ctx, u.db, boil.Infer())
	if err != nil {
		return fmt.Errorf("sign up newly registration error: %v", err)
	}
	return nil
}

func (u userRepository) Select(ctx context.Context, email model.MailAddress) (*model.User, error) {
	user, err := models.Users(
		qm.Where("mailaddress=?", email.String()),
		qm.Limit(1),
	).One(ctx, u.db)
	if err != nil {
		return nil, fmt.Errorf("sign up data acquisition error: %v", err)
	}

	newUser, err := u.userFactory.Create(user.ID, user.Username, user.Mailaddress)
	if err != nil {
		return nil, err
	}
	return newUser, nil
}
