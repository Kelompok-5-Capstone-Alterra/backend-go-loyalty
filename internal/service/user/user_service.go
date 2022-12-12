package userService

import (
	"backend-go-loyalty/internal/dto"
	"backend-go-loyalty/internal/entity"
	userRepository "backend-go-loyalty/internal/repository/user"
	"backend-go-loyalty/pkg/utils"
	"context"
	"errors"

	"github.com/google/uuid"
)

type UserServiceInterface interface {
	UpdatePassword(ctx context.Context, req dto.UpdatePasswordRequest, id uuid.UUID) error
	UpdateUserData(ctx context.Context, req dto.UserUpdate, id uuid.UUID) (dto.SignInResponse, error)
	DeleteUserData(ctx context.Context, id uuid.UUID) error
	GetUsers(ctx context.Context, query string) (dto.UserResponses, error)
	GetUserByID(ctx context.Context, id uuid.UUID) (dto.UserResponse, error)
}
type userService struct {
	ur userRepository.UserRepositoryInterface
}

func NewUserService(ur userRepository.UserRepositoryInterface) userService {
	return userService{
		ur: ur,
	}
}

func (us userService) GetUsers(ctx context.Context, query string) (dto.UserResponses, error) {
	data, err := us.ur.GetUsers(ctx, query)
	if err != nil {
		return nil, err
	}

	var users dto.UserResponses
	for _, val := range data {
		user := dto.UserResponse{
			ID:           val.ID,
			Name:         val.Name,
			Email:        val.Email,
			MobileNumber: val.MobileNumber,
			CreatedAt:    val.CreatedAt,
			UpdatedAt:    val.UpdatedAt,
			Role: dto.RoleResponse{
				ID:        val.Role.ID,
				Name:      val.Role.Name,
				CreatedAt: val.Role.CreatedAt,
				UpdatedAt: val.Role.UpdatedAt,
			},
			UserCoin: dto.UserCoinResponse{
				ID:     val.UserCoin.ID,
				Amount: val.UserCoin.Amount,
			},
			Credit: dto.CreditResponse{
				ID:     val.Credit.ID,
				Amount: val.Credit.Amount,
			},
		}
		users = append(users, user)
	}
	if users == nil {
		users = dto.UserResponses{}
	}
	return users, nil
}
func (us userService) GetUserByID(ctx context.Context, id uuid.UUID) (dto.UserResponse, error) {
	data, err := us.ur.GetUserByID(ctx, id)
	if err != nil {
		return dto.UserResponse{}, err
	}
	user := dto.UserResponse{
		ID:           data.ID,
		Name:         data.Name,
		Email:        data.Email,
		MobileNumber: data.MobileNumber,
		CreatedAt:    data.CreatedAt,
		UpdatedAt:    data.UpdatedAt,
		Role: dto.RoleResponse{
			ID:        data.Role.ID,
			Name:      data.Role.Name,
			CreatedAt: data.Role.CreatedAt,
			UpdatedAt: data.Role.UpdatedAt,
		},
		UserCoin: dto.UserCoinResponse{
			ID:     data.UserCoin.ID,
			Amount: data.UserCoin.Amount,
		},
		Credit: dto.CreditResponse{
			ID:     data.Credit.ID,
			Amount: data.Credit.Amount,
		},
	}
	return user, nil
}

func (us userService) DeleteUserData(ctx context.Context, id uuid.UUID) error {
	err := us.ur.DeleteUserData(ctx, id)
	return err
}

func (us userService) UpdatePassword(ctx context.Context, req dto.UpdatePasswordRequest, id uuid.UUID) error {
	if req.NewPassword == req.OldPassword {
		return errors.New("new password must be different from old password")
	}
	password := utils.HashPassword(req.OldPassword)
	err := us.ur.MatchPassword(ctx, password, id)
	if err != nil {
		return err
	}
	newPassword := utils.HashPassword(req.NewPassword)
	user := entity.User{
		ID:       id,
		Password: newPassword,
	}
	_, err = us.ur.UpdateUserData(ctx, user)
	if err != nil {
		return err
	}
	return nil
}
func (us userService) UpdateUserData(ctx context.Context, req dto.UserUpdate, id uuid.UUID) (dto.SignInResponse, error) {
	user := entity.User{
		ID:           id,
		Name:         req.Name,
		MobileNumber: req.MobileNumber,
	}
	res, err := us.ur.UpdateUserData(ctx, user)
	if err != nil {
		return dto.SignInResponse{}, err
	}
	data := dto.JWTData{
		Name:         res.Name,
		Email:        res.Email,
		MobileNumber: res.MobileNumber,
		Role: dto.RoleResponse{
			ID:        res.Role.ID,
			Name:      res.Role.Name,
			CreatedAt: res.Role.CreatedAt,
			UpdatedAt: res.Role.UpdatedAt,
		},
		UserCoin: dto.UserCoinResponse{
			ID:     res.UserCoin.ID,
			Amount: res.UserCoin.Amount,
		},
		Credit: dto.CreditResponse{
			ID:     res.Credit.ID,
			Amount: res.Credit.Amount,
		},
	}
	t, rt := utils.CreateLoginToken(id, data)
	tokens := dto.SignInResponse{
		AccessToken:  t,
		RefreshToken: rt,
	}
	return tokens, nil
}
