package authService

import (
	"backend-go-loyalty/internal/dto"
	"backend-go-loyalty/internal/entity"
	authRepository "backend-go-loyalty/internal/repository/auth"
	"backend-go-loyalty/pkg/config"
	"backend-go-loyalty/pkg/utils"
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
)

type AuthService interface {
	Login(ctx context.Context, req dto.SignInRequest) (dto.SignInResponse, error)
	SignUp(ctx context.Context, req dto.SignUpRequest) error
	ValidateOTP(ctx context.Context, req dto.ValidateOTP) error
	RequestNewOTP(ctx context.Context, email string) error
	RegenerateToken(ctx context.Context, rt string) (dto.SignInResponse, error)
	ForgotPasswordToken(ctx context.Context, req dto.ForgotPasswordTokenRequest) error
	ValidateForgotPasswordToken(ctx context.Context, email string, token string, req dto.NewPassword) (dto.SignInResponse, error)
}

type authService struct {
	ar authRepository.AuthRepository
}

func NewAuthService(ar authRepository.AuthRepository) authService {
	return authService{
		ar: ar,
	}
}

func (as authService) ValidateForgotPasswordToken(ctx context.Context, email string, token string, req dto.NewPassword) (dto.SignInResponse, error) {
	password := utils.HashPassword(req.Password)
	user, err := as.ar.ValidateForgotPassword(ctx, email, token, password)
	if err != nil {
		return dto.SignInResponse{}, err
	}
	data := dto.JWTData{
		Name:         user.Name,
		Email:        user.Email,
		MobileNumber: user.MobileNumber,
		Role: dto.RoleResponse{
			ID:        user.Role.ID,
			Name:      user.Role.Name,
			CreatedAt: user.Role.CreatedAt,
			UpdatedAt: user.Role.UpdatedAt,
		},
		UserCoin: dto.UserCoinResponse{
			ID:     user.UserCoin.ID,
			Amount: user.UserCoin.Amount,
		},
		Credit: dto.CreditResponse{
			ID:     user.Credit.ID,
			Amount: user.Credit.Amount,
		},
	}
	accessToken, refreshToken := utils.CreateLoginToken(user.ID, data)
	res := dto.SignInResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	return res, nil
}

func (as authService) ForgotPasswordToken(ctx context.Context, req dto.ForgotPasswordTokenRequest) error {
	_, err := as.ar.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return err
	}
	token := utils.HashPassword(utils.GenerateOTP())
	token = token[:6]
	forgotpassword := entity.ForgotPassword{
		Email:     req.Email,
		Token:     token,
		ExpiredAt: time.Now().Add(5 * time.Minute),
	}
	err = as.ar.InsertForgotPassword(ctx, forgotpassword)
	utils.ForgotPasswordToEmail(forgotpassword)
	return err
}

func (as authService) Login(ctx context.Context, req dto.SignInRequest) (dto.SignInResponse, error) {
	hashPass := utils.HashPassword(req.Password)
	user, err := as.ar.Login(ctx, req.Email, hashPass)
	if err != nil {
		return dto.SignInResponse{}, err
	}
	data := dto.JWTData{
		Name:         user.Name,
		Email:        user.Email,
		MobileNumber: user.MobileNumber,
		Role: dto.RoleResponse{
			ID:        user.Role.ID,
			Name:      user.Role.Name,
			CreatedAt: user.Role.CreatedAt,
			UpdatedAt: user.Role.UpdatedAt,
		},
		UserCoin: dto.UserCoinResponse{
			ID:     user.UserCoin.ID,
			Amount: user.UserCoin.Amount,
		},
		Credit: dto.CreditResponse{
			ID:     user.Credit.ID,
			Amount: user.Credit.Amount,
		},
	}
	accessToken, refreshToken := utils.CreateLoginToken(user.ID, data)
	res := dto.SignInResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	return res, nil
}

func (as authService) SignUp(ctx context.Context, req dto.SignUpRequest) error {
	hashPass := utils.HashPassword(req.Password)
	id := uuid.New()
	user := entity.User{
		ID:           id,
		Name:         req.Name,
		Email:        req.Email,
		Password:     hashPass,
		MobileNumber: req.MobileNumber,
		IsActive:     false,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		RoleID:       2,
	}
	err := as.ar.SignUp(ctx, user)
	if err != nil {
		return err
	}
	otp := utils.GenerateOTP()
	err = as.ar.InsertOTP(ctx, otp, user.Email)
	if err != nil {
		return err
	}
	err = utils.SendOTPToEmail(otp, user.Email)
	if err != nil {
		return err
	}
	return nil
}

func (as authService) RequestNewOTP(ctx context.Context, email string) error {
	user, err := as.ar.GetUserByEmail(ctx, email)
	if err != nil {
		return err
	}
	if user.IsActive {
		return errors.New("cannot send new otp to activated user")
	}
	otp := utils.GenerateOTP()
	err = as.ar.InsertOTP(ctx, otp, email)
	if err != nil {
		return err
	}
	err = utils.SendOTPToEmail(otp, email)
	if err != nil {
		return err
	}
	return nil
}

func (as authService) ValidateOTP(ctx context.Context, req dto.ValidateOTP) error {
	err := as.ar.ValidateOTP(ctx, req.OTP, req.Email)
	if err != nil {
		return err
	}
	return nil
}

func (as authService) RegenerateToken(ctx context.Context, rt string) (dto.SignInResponse, error) {
	tokenEnv := config.GetTokenEnv()
	id, created_at, err := utils.GetDataFromRefreshToken(rt)
	interval := time.Now().Sub(created_at)
	if interval.Hours() > float64(tokenEnv.RefreshTokenTTLHour) {
		return dto.SignInResponse{}, errors.New("refresh token invalid")
	}
	if err != nil {
		return dto.SignInResponse{}, err
	}
	user, err := as.ar.GetUserByID(ctx, id)
	data := dto.JWTData{
		Name:         user.Name,
		Email:        user.Email,
		MobileNumber: user.MobileNumber,
		Role: dto.RoleResponse{
			ID:        user.Role.ID,
			Name:      user.Role.Name,
			CreatedAt: user.Role.CreatedAt,
			UpdatedAt: user.Role.UpdatedAt,
		},
		UserCoin: dto.UserCoinResponse{
			ID:     user.UserCoin.ID,
			Amount: user.UserCoin.Amount,
		},
		Credit: dto.CreditResponse{
			ID:     user.Credit.ID,
			Amount: user.Credit.Amount,
		},
	}
	accessToken, refreshToken := utils.CreateLoginToken(user.ID, data)
	res := dto.SignInResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	return res, nil
}
