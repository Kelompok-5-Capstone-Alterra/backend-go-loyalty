package rewardRepository

import (
	"backend-go-loyalty/internal/entity"
	"backend-go-loyalty/internal/model"
	"context"

	"gorm.io/gorm"
)

type IRewardRepository interface {
	CreateReward(ctx context.Context, reward *entity.Reward) error
	FindAllReward(ctx context.Context) (*entity.Rewards, error)
	FindRewardByID(ctx context.Context, id uint64) (*entity.Reward, error)
	UpdateReward(ctx context.Context, r *entity.Reward, id uint64) error
	DeleteReward(ctx context.Context, id uint64) error
}

type rewardRepository struct {
	DB *gorm.DB
}

func NewRewardRepository(db *gorm.DB) IRewardRepository {
	return &rewardRepository{db}
}

func (rr *rewardRepository) FindAllReward(ctx context.Context) (*entity.Rewards, error) {
	var rewards entity.Rewards
	err := rr.DB.Find(&rewards).Error
	if err != nil {
		return nil, err
	}
	return &rewards, nil
}

func (rr *rewardRepository) FindRewardByID(ctx context.Context, id uint64) (*entity.Reward, error) {
	var reward entity.Reward
	err := rr.DB.First(&reward, id).Error
	if err != nil {
		return nil, err
	}
	return &reward, err
}

func (rr *rewardRepository) CreateReward(ctx context.Context, reward *entity.Reward) error {
	err := rr.DB.Create(&reward).Error
	if err != nil {
		return err
	}
	return nil
}

func (rr *rewardRepository) UpdateReward(ctx context.Context, r *entity.Reward, id uint64) error {
	err := rr.DB.Model(&model.Reward{}).Where("id = ?", id).Updates(r).Error

	if err != nil {
		return err
	}
	return nil
}

func (rr *rewardRepository) DeleteReward(ctx context.Context, id uint64) error {
	var reward entity.Reward
	err := rr.DB.Delete(&reward, id).Error

	if err != nil {
		return err
	}
	return nil
}
