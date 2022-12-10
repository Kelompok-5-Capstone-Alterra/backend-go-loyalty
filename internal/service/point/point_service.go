package pointService

import (
	"backend-go-loyalty/internal/dto"
	pointRepository "backend-go-loyalty/internal/repository/point"
	"context"

	"github.com/google/uuid"
)

type IPointService interface {
	GetAllPoints(ctx context.Context) (dto.UserCoinResponses, error)
	GetPoint(ctx context.Context, id uuid.UUID) (dto.UserCoinResponse, error)
}

type pointService struct {
	pr pointRepository.IPointRepository
}

func NewPointService(pr pointRepository.IPointRepository) pointService {
	return pointService{
		pr: pr,
	}
}

func (ps pointService) GetAllPoints(ctx context.Context) (dto.UserCoinResponses, error) {
	data, err := ps.pr.GetAllPoints(ctx)
	if err != nil {
		return nil, err
	}

	var res dto.UserCoinResponses
	for _, val := range data {
		point := dto.UserCoinResponse{
			ID:     val.ID,
			Amount: val.Amount,
		}

		res = append(res, point)
	}
	if res == nil {
		res = dto.UserCoinResponses{}
	}
	return res, err
}
func (ps pointService) GetPoint(ctx context.Context, id uuid.UUID) (dto.UserCoinResponse, error) {
	data, err := ps.pr.GetPoint(ctx, id)
	if err != nil {
		return dto.UserCoinResponse{}, err
	}

	point := dto.UserCoinResponse{
		ID:     data.ID,
		Amount: data.Amount,
	}
	return point, nil
}
