package authRepository

import (
	"backend-go-loyalty/internal/entity"
	"backend-go-loyalty/internal/model"
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AuthRepository interface {
	// GetUsers(ctx context.Context) (entity.Users, error)
	GetUserByID(ctx context.Context, id uuid.UUID) (entity.User, error)
	GetUserByEmail(ctx context.Context, email string) (entity.User, error)
	Login(ctx context.Context, email string, password string) (entity.User, error)
	SignUp(ctx context.Context, req entity.User) error
	InsertOTP(ctx context.Context, otp string, email string) error
	ValidateOTP(ctx context.Context, otp string, email string) error
	InsertForgotPassword(ctx context.Context, req entity.ForgotPassword) error
	ValidateForgotPassword(ctx context.Context, email string, token string, password string) (entity.User, error)
}

type authRepository struct {
	DB *gorm.DB
}

func NewAuthRepository(db *gorm.DB) authRepository {
	return authRepository{
		DB: db,
	}
}

func (ar authRepository) ValidateForgotPassword(ctx context.Context, email string, token string, password string) (entity.User, error) {
	fp := entity.ForgotPassword{}
	err := ar.DB.Where("email = ? AND token = ?", email, token).First(&fp).Error
	if err != nil {
		return entity.User{}, err
	}
	var forgotpassword entity.ForgotPassword
	if fp.ExpiredAt.Before(time.Now()) {
		err = ar.DB.Where("email = ? AND token = ?", email, token).Delete(&forgotpassword).Error
		if err != nil {
			return entity.User{}, err
		}
		return entity.User{}, errors.New("token expired")
	}
	user := entity.User{
		Password: password,
	}
	err = ar.DB.Model(&model.User{}).Where("email = ?", email).Updates(user).Error
	if err != nil {
		return entity.User{}, err
	}
	err = ar.DB.Where("email = ? AND token = ?", email, token).Delete(&forgotpassword).Error
	if err != nil {
		return entity.User{}, err
	}
	user = entity.User{}
	err = ar.DB.Where("email = ?", email).First(&user).Error
	return user, err
}

func (ar authRepository) GetUserByID(ctx context.Context, id uuid.UUID) (entity.User, error) {
	user := entity.User{}
	err := ar.DB.Model(&model.User{}).Preload("Role").First(&user, id).Error
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}
func (ar authRepository) GetUserByEmail(ctx context.Context, email string) (entity.User, error) {
	user := entity.User{}
	err := ar.DB.Model(&model.User{}).Preload("Role").Where("email = ?", email).First(&user).Error
	// if err != nil {
	// 	if err.Error() == "record not found" {
	// 		err = errors.New("user with such email haven't signed up")
	// 	}
	// 	return entity.User{}, err
	// }
	return user, err
}

func (ar authRepository) Login(ctx context.Context, email string, password string) (entity.User, error) {
	user := entity.User{}
	err := ar.DB.Model(&model.User{}).Preload("Role").Where("email = ?", email).First(&user).Error
	if err != nil {
		return entity.User{}, err
	}
	if user.Password != password {
		return entity.User{}, errors.New("wrong password")
	}
	if user.IsActive == false {
		return entity.User{}, errors.New("user not activated")
	}
	return user, nil
}

func (ar authRepository) SignUp(ctx context.Context, req entity.User) error {
	user := entity.User{}
	coin := entity.UserCoin{
		Amount: 0,
	}
	credit := entity.Credit{
		Amount: 0,
	}
	err := ar.DB.Model(&model.User{}).Preload("Role").Where("email = ?", req.Email).First(&user).Error
	if err != nil {
		if err.Error() == "record not found" {
			result := ar.DB.Create(&req)
			if result.Error != nil {
				return result.Error
			}
			err := ar.DB.Create(&coin).Error
			if err != nil {
				return err
			}
			err = ar.DB.Create(&credit).Error
			if err != nil {
				return err
			}
		}
	}
	if !user.IsActive {
		err = ar.DB.Model(&model.User{}).Where("email = ?", req.Email).Updates(&req).Error
		return err
	}
	return errors.New("email already used")

}

func (ar authRepository) InsertOTP(ctx context.Context, otp string, email string) error {
	ar.DB.Unscoped().Where("email = ?", email).Delete(&model.OTP{})
	req := entity.OTP{
		OTPCode:   otp,
		Email:     email,
		CreatedAt: time.Now(),
	}
	err := ar.DB.Create(&req).Error
	if err != nil {
		return err
	}
	return nil
}

func (ar authRepository) ValidateOTP(ctx context.Context, otp string, email string) error {
	err := ar.DB.Where("otp_code = ? AND email = ?", otp, email).First(&entity.OTP{}).Error
	if err != nil {
		var errMessage string
		if err.Error() == "record not found" {
			errMessage = fmt.Sprint("invalid otp code")
		} else {
			errMessage = fmt.Sprint(err.Error())
		}
		return errors.New(errMessage)
	}
	err = ar.DB.Model(&model.User{}).Where("email = ?", email).Update("is_active", true).Error
	if err != nil {
		return err
	}
	ar.DB.Unscoped().Where("email = ? AND otp_code = ?", email, otp).Delete(&model.OTP{})
	return nil
}

func (ar authRepository) InsertForgotPassword(ctx context.Context, req entity.ForgotPassword) error {
	err := ar.DB.Create(&req).Error
	return err
}
