package userService

import (
	"backend-go-loyalty/internal/dto"
	"backend-go-loyalty/internal/entity"
	userRepository "backend-go-loyalty/internal/repository/user"
	"backend-go-loyalty/pkg/utils"
	"context"
)

type UserServiceInterface interface {
	UpdatePassword(ctx context.Context, req dto.UpdatePasswordRequest, id uint64) error
	UpdateUserData(ctx context.Context, req dto.UserUpdate, id uint64) (dto.SignInResponse, error)
	DeleteUserData(ctx context.Context, id uint64) error
}
type userService struct {
	ur userRepository.UserRepositoryInterface
}

func NewUserService(ur userRepository.UserRepositoryInterface) userService {
	return userService{
		ur: ur,
	}
}

func (us userService) DeleteUserData(ctx context.Context, id uint64) error {
	err := us.ur.DeleteUserData(ctx, id)
	return err
}

func (us userService) UpdatePassword(ctx context.Context, req dto.UpdatePasswordRequest, id uint64) error {
	password := utils.HashPassword(req.OldPassword)
	err := us.ur.MatchPassword(ctx, password, id)
	if err != nil {
		return nil
	}
	req.NewPassword = utils.HashPassword(req.NewPassword)
	user := entity.User{
		ID:       id,
		Password: req.NewPassword,
	}
	_, err = us.ur.UpdateUserData(ctx, user)
	if err != nil {
		return nil
	}
	return nil
}
func (us userService) UpdateUserData(ctx context.Context, req dto.UserUpdate, id uint64) (dto.SignInResponse, error) {
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
	}
	t, rt := utils.CreateLoginToken(id, data)
	tokens := dto.SignInResponse{
		AccessToken:  t,
		RefreshToken: rt,
	}
	return tokens, nil
}
