package service

import (
	"backend-go-loyalty/internal/repository/reward"
	"backend-go-loyalty/internal/dto"

	"gorm.io/gorm"
	"context"
)

type RewardService interface {
	CreateReward(ctx context.Context, name string, description string, requiredPoint int) error
	FindAll(ctx context.Context) (*dto.RewardResponse, error)
	FindRewardByID(ctx context.Context, rewardID int) (dto dto.RewardResponse, error)
	UpdateReward(ctx context.Context) error
	DeleteReward(ctx context.Context, rewardID int) error
}

type rewardServiceImpl struct {
	rw repository.IRewardRepository
	db *gorm.DB
	dto dto.RewardResponse
}

func ProvideRewardService(rw repository.IRewardRepository, db *gorm.DB, dto dto.RewardResponse) *rewardServiceImpl {
	return &rewardServiceImpl{
		db: db,
		rw: rw,
		dto: dto,
	}
}

func (rs rewardServiceImpl) FindAll(ctx context.Context) (*dto.RewardResponse, error) {
	rewards, err := rs.rw.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	rewardResponse := dto.RewardResponse{}
	for _, reward := range rewards {
		var item dto.RewardResponse
		item.RewardID = reward.RewardID
		item.Name = reward.Name
		item.Description = reward.Description
		item.RequiredPoint = reward.RequiredPoint
		rewardResponse = append(rewardResponse, &item)
	}
	return rewardResponse, nil
}

func (rs rewardServiceImpl) FindRewardByID(ctx context.Context, rewardID int) (dto dto.RewardResponse, error) {
	rwrd, err := rs.rw.FindRewardByID(ctx, rewardID)
	if err != nil {
		return dto.RewardResponse{}, err
	}
	if len(rwrd) < 1 {
		return nil, errors.New("no reward found")
	}

	rewardResponse := dto.RewardResponse{}
	for _, reward := range rwrd {
		var item dto.RewardResponse
		item.RewardID = reward.RewardID
		item.Name = reward.Name
		item.Description = reward.Description
		item.RequiredPoint = reward.RequiredPoint
	}
	return rewardResponse, nil
}

func (rs rewardServiceImpl) CreateReward(ctx context.Context, name string, description string, requiredPoint int) error {
	err := rs.rw.CreateReward(ctx, reward.Name, reward.Description, reward.RequiredPoint)
	if err != nil {
		return 0, err
	}
	return nil
}

func (rs rewardServiceImpl) UpdateReward(ctx context.Context) error {
	err := rs.rw.UpdateReward(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (rs rewardServiceImpl) DeleteReward(ctx context.Context, rewardID int) error {
	err := rs.rw.DeleteReward(ctx, rewardID)
	if err != nil {
		return err
	}
	return nil
}