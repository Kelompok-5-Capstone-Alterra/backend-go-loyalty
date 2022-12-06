package redeemRepository

import (
	"backend-go-loyalty/internal/entity"
	"backend-go-loyalty/internal/model"
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IRedeemRepository interface {
	CreateRedeem(ctx context.Context, redeem *entity.Redeem) error
	GetAllRedeems(ctx context.Context) (*entity.Redeems, error)
	GetAllRedeemByUserID(ctx context.Context, userID uuid.UUID) (*entity.Redeems, error)
	GetRedeemByID(ctx context.Context, id uint64) (*entity.Redeem, error)
	UpdateRedeem(ctx context.Context, d *entity.Redeem, id uint64) error
	DeleteRedeem(ctx context.Context, id uint64) error
}

type redeemRepository struct {
	DB *gorm.DB
}

func NewRedeemRepository(db *gorm.DB) IRedeemRepository {
	return &redeemRepository{db}
}

func (dr *redeemRepository) GetAllRedeems(ctx context.Context) (*entity.Redeems, error) {
	var redeems entity.Redeems
	err := dr.DB.Model(&model.Redeem{}).Preload("Reward").Preload("User").Find(&redeems).Error
	if err != nil {
		return nil, err
	}
	return &redeems, nil
}

func (dr *redeemRepository) GetAllRedeemByUserID(ctx context.Context, userID uuid.UUID) (*entity.Redeems, error) {
	var redeems entity.Redeems
	err := dr.DB.Model(&model.Redeem{}).Preload("Reward").Preload("User").Where("user_id = ?", userID).Find(&redeems).Error
	if err != nil {
		return nil, err
	}
	return &redeems, nil
}

func (dr *redeemRepository) GetRedeemByID(ctx context.Context, id uint64) (*entity.Redeem, error) {
	var redeem entity.Redeem
	err := dr.DB.Model(&model.Redeem{}).Preload("Reward").Preload("User").First(&redeem, id).Error
	if err != nil {
		return nil, err
	}
	return &redeem, err
}

func (dr *redeemRepository) CreateRedeem(ctx context.Context, redeem *entity.Redeem) error {
	err := dr.DB.Create(&redeem).Error
	if err != nil {
		return err
	}
	return nil
}

func (dr *redeemRepository) UpdateRedeem(ctx context.Context, d *entity.Redeem, id uint64) error {
	err := dr.DB.Model(&model.Redeem{}).Where("id = ?", id).Updates(d).Error
	if err != nil {
		return err
	}
	return nil
}

func (dr *redeemRepository) DeleteRedeem(ctx context.Context, id uint64) error {
	var redeem entity.Redeem
	err := dr.DB.Delete(&redeem, id).Error

	if err != nil {
		return err
	}
	return nil
}
