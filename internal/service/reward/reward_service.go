package rewardService

import (
	"backend-go-loyalty/internal/dto"
	"backend-go-loyalty/internal/entity"
	rewardRepository "backend-go-loyalty/internal/repository/reward"
	"time"

	"context"
)

type IRewardService interface {
	CreateReward(ctx context.Context, req dto.RewardRequest) error
	FindAllReward(ctx context.Context) (dto.RewardsResponse, error)
	FindRewardByID(ctx context.Context, rewardID uint64) (dto.RewardResponse, error)
	UpdateReward(ctx context.Context, req dto.RewardRequest, id uint64) error
	DeleteReward(ctx context.Context, rewardID uint64) error
}

type rewardServiceImpl struct {
	rr rewardRepository.IRewardRepository
}

func NewRewardService(rr rewardRepository.IRewardRepository) rewardServiceImpl {
	return rewardServiceImpl{
		rr: rr,
	}
}

func (rs rewardServiceImpl) FindAllReward(ctx context.Context) (dto.RewardsResponse, error) {
	rewards, err := rs.rr.FindAllReward(ctx)
	if err != nil {
		return nil, err
	}

	var rewardResponses dto.RewardsResponse
	for _, reward := range *rewards {
		var item dto.RewardResponse
		item.ID = reward.ID
		item.Name = reward.Name
		item.Description = reward.Description
		item.RequiredPoint = reward.RequiredPoint
		item.ValidUntil = reward.ValidUntil
		item.CreatedAt = reward.CreatedAt
		item.UpdatedAt = reward.UpdatedAt
		item.DeletedAt = reward.DeletedAt
		item.Category = dto.CategoryResponse{
			ID:        reward.CategoryID,
			Name:      reward.Category.Name,
			CreatedAt: reward.Category.CreatedAt,
			UpdatedAt: reward.Category.UpdatedAt,
		}
		rewardResponses = append(rewardResponses, item)
	}
	if rewardResponses == nil {
		rewardResponses = dto.RewardsResponse{}
	}
	return rewardResponses, nil
}

func (rs rewardServiceImpl) FindRewardByID(ctx context.Context, rewardID uint64) (dto.RewardResponse, error) {
	reward, err := rs.rr.FindRewardByID(ctx, rewardID)
	if err != nil {
		return dto.RewardResponse{}, err
	}

	rewardResponse := dto.RewardResponse{
		ID:            reward.ID,
		Name:          reward.Name,
		Description:   reward.Description,
		RequiredPoint: reward.RequiredPoint,
		ValidUntil:    reward.ValidUntil,
		CreatedAt:     reward.CreatedAt,
		UpdatedAt:     reward.UpdatedAt,
		DeletedAt:     reward.DeletedAt,
		Category: dto.CategoryResponse{
			ID:        reward.CategoryID,
			Name:      reward.Category.Name,
			CreatedAt: reward.Category.CreatedAt,
			UpdatedAt: reward.Category.UpdatedAt,
		},
	}
	return rewardResponse, nil
}

func (rs rewardServiceImpl) CreateReward(ctx context.Context, req dto.RewardRequest) error {
	valid, err := time.Parse(time.RFC3339, req.ValidUntil)
	if err != nil {
		return err
	}
	reward := entity.Reward{
		Name:          req.Name,
		Description:   req.Description,
		RequiredPoint: req.RequiredPoint,
		ValidUntil:    valid,
		CategoryID:    req.CategoryID,
	}

	err = rs.rr.CreateReward(ctx, &reward)
	if err != nil {
		return err
	}
	return nil
}

func (rs rewardServiceImpl) UpdateReward(ctx context.Context, req dto.RewardRequest, id uint64) error {
	reward := entity.Reward{
		Name:          req.Name,
		Description:   req.Description,
		RequiredPoint: req.RequiredPoint,
		CategoryID:    req.CategoryID,
		UpdatedAt:     time.Now(),
	}

	if req.ValidUntil != "" {
		valid, err := time.Parse(time.RFC3339, req.ValidUntil)
		if err != nil {
			return err
		}
		reward.ValidUntil = valid
	}

	err := rs.rr.UpdateReward(ctx, reward, id)
	return err
}

func (rs rewardServiceImpl) DeleteReward(ctx context.Context, rewardID uint64) error {
	err := rs.rr.DeleteReward(ctx, rewardID)
	if err != nil {
		return err
	}
	return nil
}
