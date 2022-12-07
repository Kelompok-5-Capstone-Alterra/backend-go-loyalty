package userRepository

import (
	"backend-go-loyalty/internal/entity"
	"backend-go-loyalty/internal/model"
	"context"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	UpdateUserData(ctx context.Context, req entity.User) (entity.User, error)
	MatchPassword(ctx context.Context, password string, id uuid.UUID) error
	DeleteUserData(ctx context.Context, id uuid.UUID) error
	GetUsers(ctx context.Context, name string) (entity.Users, error)
	GetUserByID(ctx context.Context, id uuid.UUID) (entity.User, error)
}

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(dbConn *gorm.DB) userRepository {
	return userRepository{
		DB: dbConn,
	}
}

func (ur userRepository) GetUsers(ctx context.Context, name string) (entity.Users, error) {
	users := entity.Users{}
	keyword := "%" + name + "%"
	err := ur.DB.Model(&model.User{}).Preload("Role").Where("name LIKE ?", keyword).Find(&users).Error
	return users, err
}
func (ur userRepository) GetUserByID(ctx context.Context, id uuid.UUID) (entity.User, error) {
	user := entity.User{}
	err := ur.DB.Model(&model.User{}).Preload("Role").First(&user, id).Error
	return user, err
}

func (ur userRepository) DeleteUserData(ctx context.Context, id uuid.UUID) error {
	err := ur.DB.Delete(&entity.User{}, id).Error
	return err
}

func (ur userRepository) UpdateUserData(ctx context.Context, req entity.User) (entity.User, error) {
	result := ur.DB.Model(&model.User{}).Where("id = ?", req.ID).Updates(req)
	if result.Error != nil {
		return entity.User{}, result.Error
	}
	res := entity.User{}
	err := ur.DB.Model(&model.User{}).Preload("Role").First(&res, req.ID).Error
	if err != nil {
		return entity.User{}, err
	}
	return res, nil
}

func (ur userRepository) MatchPassword(ctx context.Context, password string, id uuid.UUID) error {
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
