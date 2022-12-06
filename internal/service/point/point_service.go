package pointService

import (
	"backend-go-loyalty/internal/dto"
	pointRepository "backend-go-loyalty/internal/repository/point"
	"context"

	"github.com/google/uuid"
)

type IPointService interface {
	GetAllPoints(ctx context.Context) (dto.PointResponses, error)
	GetPoint(ctx context.Context, id uuid.UUID) (dto.PointResponse, error)
}

type pointService struct {
	pr pointRepository.IPointRepository
}

func NewPointService(pr pointRepository.IPointRepository) pointService {
	return pointService{
		pr: pr,
	}
}

func (ps pointService) GetAllPoints(ctx context.Context) (dto.PointResponses, error) {
	data, err := ps.pr.GetAllPoints(ctx)
	if err != nil {
		return nil, err
	}

	var res dto.PointResponses
	for _, val := range data {
		point := dto.PointResponse{
			ID:     val.ID,
			Amount: val.Amount,
			User: dto.UserResponse{
				ID:           val.User.ID,
				Name:         val.User.Name,
				Email:        val.User.Email,
				MobileNumber: val.User.MobileNumber,
				CreatedAt:    val.User.CreatedAt,
				UpdatedAt:    val.User.UpdatedAt,
				Role: dto.RoleResponse{
					ID:        val.User.Role.ID,
					Name:      val.User.Role.Name,
					CreatedAt: val.User.Role.CreatedAt,
					UpdatedAt: val.User.Role.UpdatedAt,
				},
			},
		}

		res = append(res, point)
	}
	return res, err
}
func (ps pointService) GetPoint(ctx context.Context, id uuid.UUID) (dto.PointResponse, error) {
	data, err := ps.pr.GetPoint(ctx, id)
	if err != nil {
		return dto.PointResponse{}, err
	}

	point := dto.PointResponse{
		ID:     data.ID,
		Amount: data.Amount,
		User: dto.UserResponse{
			ID:           data.User.ID,
			Name:         data.User.Name,
			Email:        data.User.Email,
			MobileNumber: data.User.MobileNumber,
			CreatedAt:    data.User.CreatedAt,
			UpdatedAt:    data.User.UpdatedAt,
			Role: dto.RoleResponse{
				ID:        data.User.Role.ID,
				Name:      data.User.Role.Name,
				CreatedAt: data.User.Role.CreatedAt,
				UpdatedAt: data.User.Role.UpdatedAt,
			},
		},
	}
	return point, nil
}
