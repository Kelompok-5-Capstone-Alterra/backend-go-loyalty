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
}

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(dbConn *gorm.DB) userRepository {
	return userRepository{
		DB: dbConn,
	}
}

func (ur userRepository) UpdateUserData(ctx context.Context, req entity.User) (entity.User, error) {
	result := ur.DB.Model(&model.User{}).Where("id = ?", req.ID).Updates(req)
	var errMessage string
	if result.RowsAffected == 0 {
		errMessage = "old password not match"
	}
	if result.Error != nil {
		return entity.User{}, errors.New(errMessage)
	}
	res := entity.User{}
	err := ur.DB.Model(&model.User{}).Preload("Role").First(&res, req.ID).Error
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
