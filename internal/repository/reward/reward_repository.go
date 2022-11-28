package repository

import (
	"backend-go-loyalty/internal/entity"
	"log"

	"gorm.io/gorm"
)

type IRewardRepository interface {
	CreateReward(reward *entity.Reward) (*entity.Reward, error)
	FindAll(r []*entity.Reward) ([]*entity.Reward, error)
	FindRewardByID(id int) (*entity.Reward, error)
	UpdateReward(r *entity.Reward, id int) (*entity.Reward, error)
	DeleteReward(id int) error
}
type rewardRepository struct {
	Db *gorm.DB
}

func NewRewardRepository(db *gorm.DB) IRewardRepository {
	return &rewardRepository{db}
}

func (rw *rewardRepository) FindAll(r []*entity.Reward) ([]*entity.Reward, error) {
	err := rw.Db.Find(&r).Error
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (rw *rewardRepository) FindRewardByID(id int) (*entity.Reward, error) {
	var reward entity.Reward
	err := rw.Db.Where("id = ?", id).First(&reward).Error
	if err != nil {
		return nil, err
	}
	return &reward, err
}

func (rw *rewardRepository) CreateReward(reward *entity.Reward) (*entity.Reward, error) {
	err := rw.Db.Where("name =? AND descriptions =? AND required_points =?", reward.Name, reward.Description, reward.RequiredPoint).First(&reward).Error

	if err != nil {
		log.Fatalf("error when inserting data: %s", err)
	}

	err = rw.Db.Create(&reward).Error
	if err != nil {
		return nil, err
	}
	return reward, err
}

func (rw *rewardRepository) UpdateReward(r *entity.Reward, id int) (*entity.Reward, error) {
	var reward entity.Reward
	var err = rw.Db.Find(&reward, id).Error
	reward.Name = r.Name
	reward.Description = r.Description
	reward.RequiredPoint = r.RequiredPoint

	err = rw.Db.Save(&reward).Error

	if err != nil {
		return nil, err
	}

	return r, err
}

func (rw *rewardRepository) DeleteReward(id int) error {
	var reward entity.Reward
	err := rw.Db.Delete(&reward, id).Error

	if err != nil {
		return err
	}

	return nil

}
