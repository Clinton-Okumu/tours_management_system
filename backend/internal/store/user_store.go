package store

import (
	"backend/internal/models"
	"context"
	"errors"

	"gorm.io/gorm"
)

var ErrUserNotFound = errors.New("user not found")

type UserStore interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	GetUserByUsername(ctx context.Context, username string) (*models.User, error)
	GetUserToken(scope, tokenPlaintext string) (*models.User, error)
}

type userStore struct {
	db *gorm.DB
}

func NewUserStore(db *gorm.DB) UserStore {
	return &userStore{db}
}

func (us *userStore) CreateUser(ctx context.Context, user *models.User) error {
	return us.db.WithContext(ctx).Create(user).Error
}

func (us *userStore) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	err := us.db.WithContext(ctx).
		Where("email = ?", email).
		First(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrUserNotFound
	}
	return &user, err
}

func (us *userStore) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	var user models.User
	err := us.db.WithContext(ctx).
		Where("username = ?", username).
		First(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrUserNotFound
	}
	return &user, err
}

func (us *userStore) GetUserToken(scope, tokenPlaintext string) (*models.User, error) {
	var user models.User
	err := us.db.WithContext(context.Background()).
		Where("token = ?", tokenPlaintext).
		First(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrUserNotFound
	}
	return &user, err
}
