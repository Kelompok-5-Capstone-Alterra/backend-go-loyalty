package userRepository

import (
	"backend-go-loyalty/internal/entity"
	"backend-go-loyalty/internal/model"
	"context"
	"errors"

	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	UpdateUserData(ctx context.Context, req entity.User) (entity.User, error)
	MatchPassword(ctx context.Context, password string, id uint64) error
	DeleteUserData(ctx context.Context, id uint64) error
}

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(dbConn *gorm.DB) userRepository {
	return userRepository{
		DB: dbConn,
	}
}

func (ur userRepository) DeleteUserData(ctx context.Context, id uint64) error {
	err := ur.DB.Delete(&entity.User{}, id).Error
	return err
}

func (ur userRepository) UpdateUserData(ctx context.Context, req entity.User) (entity.User, error) {
	err := ur.DB.Model(&model.User{}).Where("id = ?", req.ID).Updates(req).Error
	if err != nil {
		return entity.User{}, err
	}
	res := entity.User{}
	err = ur.DB.Model(&model.User{}).Preload("Role").First(&res, req.ID).Error
	if err != nil {
		return entity.User{}, err
	}
	return res, nil
}

func (ur userRepository) MatchPassword(ctx context.Context, password string, id uint64) error {
	user := entity.User{}
	err := ur.DB.First(&user, id).Error
	if err != nil {
		return err
	}
	if user.Password != password {
		return errors.New("password not match")
	}
	return nil
}
