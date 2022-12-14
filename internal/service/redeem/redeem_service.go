package redeemService

import (
	"backend-go-loyalty/internal/dto"
	"backend-go-loyalty/internal/entity"
	pointRepository "backend-go-loyalty/internal/repository/point"
	redeemRepository "backend-go-loyalty/internal/repository/redeem"
	rewardRepository "backend-go-loyalty/internal/repository/reward"
	userRepository "backend-go-loyalty/internal/repository/user"
	"context"
	"errors"

	"github.com/google/uuid"
)

type IRedeemService interface {
	CreateRedeem(ctx context.Context, req dto.RedeemRequest, userID uuid.UUID) error
	GetAllRedeemByUserID(ctx context.Context, userID uuid.UUID) (dto.RedeemResponses, error)
	GetAllRedeems(ctx context.Context) (dto.RedeemResponses, error)
	GetAllIncludeSoftDeleted(ctx context.Context) (dto.RedeemResponses, error)
	GetRedeemByID(ctx context.Context, redeemID uint64, userID uuid.UUID) (dto.RedeemResponse, error)
	AdminGetRedeemByID(ctx context.Context, redeemID uint64) (dto.RedeemResponse, error)
	UpdateRedeem(ctx context.Context, req dto.RedeemUpdateRequest, id uint64) error
	DeleteRedeem(ctx context.Context, redeemID uint64) error
}

type redeemServiceImpl struct {
	dr redeemRepository.IRedeemRepository
	pr pointRepository.IPointRepository
	rr rewardRepository.IRewardRepository
	ur userRepository.UserRepositoryInterface
}

func NewRedeemService(dr redeemRepository.IRedeemRepository, pr pointRepository.IPointRepository, rr rewardRepository.IRewardRepository, ur userRepository.UserRepositoryInterface) redeemServiceImpl {
	return redeemServiceImpl{
		dr: dr,
		pr: pr,
		rr: rr,
		ur: ur,
	}
}

func (ds redeemServiceImpl) GetAllIncludeSoftDeleted(ctx context.Context) (dto.RedeemResponses, error) {
	redeems, err := ds.dr.GetAllIncludeSoftDelete(ctx)
	if err != nil {
		return nil, err
	}
	var redeemResponses dto.RedeemResponses
	for _, redeem := range *redeems {
		var item dto.RedeemResponse
		item.ID = redeem.ID
		item.PointSpent = redeem.PointSpent
		item.CreatedAt = redeem.CreatedAt
		item.DeletedAt = redeem.DeletedAt
		item.Reward = dto.RewardResponse{
			ID:            redeem.Reward.ID,
			Name:          redeem.Reward.Name,
			Description:   redeem.Reward.Description,
			RequiredPoint: redeem.Reward.RequiredPoint,
			ValidUntil:    redeem.Reward.ValidUntil,
			CreatedAt:     redeem.Reward.CreatedAt,
			UpdatedAt:     redeem.Reward.UpdatedAt,
			DeletedAt:     redeem.Reward.DeletedAt,
		}
		redeemResponses = append(redeemResponses, item)
	}
	if redeemResponses == nil {
		redeemResponses = dto.RedeemResponses{}
	}
	return redeemResponses, nil
}
func (ds redeemServiceImpl) GetAllRedeems(ctx context.Context) (dto.RedeemResponses, error) {
	redeems, err := ds.dr.GetAllRedeems(ctx)
	if err != nil {
		return nil, err
	}
	var redeemResponses dto.RedeemResponses
	for _, redeem := range *redeems {
		var item dto.RedeemResponse
		item.ID = redeem.ID
		item.PointSpent = redeem.PointSpent
		item.CreatedAt = redeem.CreatedAt
		item.DeletedAt = redeem.DeletedAt
		item.Reward = dto.RewardResponse{
			ID:            redeem.Reward.ID,
			Name:          redeem.Reward.Name,
			Description:   redeem.Reward.Description,
			RequiredPoint: redeem.Reward.RequiredPoint,
			ValidUntil:    redeem.Reward.ValidUntil,
			CreatedAt:     redeem.Reward.CreatedAt,
			UpdatedAt:     redeem.Reward.UpdatedAt,
			DeletedAt:     redeem.Reward.DeletedAt,
		}
		redeemResponses = append(redeemResponses, item)
	}
	if redeemResponses == nil {
		redeemResponses = dto.RedeemResponses{}
	}
	return redeemResponses, nil
}

func (ds redeemServiceImpl) GetAllRedeemByUserID(ctx context.Context, userID uuid.UUID) (dto.RedeemResponses, error) {
	redeems, err := ds.dr.GetAllRedeemByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	var redeemResponses dto.RedeemResponses
	for _, redeem := range *redeems {
		var item dto.RedeemResponse
		item.ID = redeem.ID
		item.PointSpent = redeem.PointSpent
		item.CreatedAt = redeem.CreatedAt
		item.DeletedAt = redeem.DeletedAt
		item.Reward = dto.RewardResponse{
			ID:            redeem.Reward.ID,
			Name:          redeem.Reward.Name,
			Description:   redeem.Reward.Description,
			RequiredPoint: redeem.Reward.RequiredPoint,
			ValidUntil:    redeem.Reward.ValidUntil,
			CreatedAt:     redeem.Reward.CreatedAt,
			UpdatedAt:     redeem.Reward.UpdatedAt,
			DeletedAt:     redeem.Reward.DeletedAt,
		}
		redeemResponses = append(redeemResponses, item)
	}
	if redeemResponses == nil {
		redeemResponses = dto.RedeemResponses{}
	}
	return redeemResponses, nil
}

func (ds redeemServiceImpl) GetRedeemByID(ctx context.Context, redeemID uint64, userID uuid.UUID) (dto.RedeemResponse, error) {
	redeem, err := ds.dr.GetRedeemByID(ctx, redeemID, userID)
	if err != nil {
		return dto.RedeemResponse{}, err
	}
	redeemResponse := dto.RedeemResponse{
		ID:         redeem.ID,
		PointSpent: redeem.PointSpent,
		CreatedAt:  redeem.CreatedAt,
		DeletedAt:  redeem.DeletedAt,
		Reward: dto.RewardResponse{
			ID:            redeem.Reward.ID,
			Name:          redeem.Reward.Name,
			Description:   redeem.Reward.Description,
			RequiredPoint: redeem.Reward.RequiredPoint,
			ValidUntil:    redeem.Reward.ValidUntil,
			CreatedAt:     redeem.Reward.CreatedAt,
			UpdatedAt:     redeem.Reward.UpdatedAt,
			DeletedAt:     redeem.Reward.DeletedAt,
		},
	}
	return redeemResponse, nil
}
func (ds redeemServiceImpl) AdminGetRedeemByID(ctx context.Context, redeemID uint64) (dto.RedeemResponse, error) {
	redeem, err := ds.dr.AdminGetRedeemByID(ctx, redeemID)
	if err != nil {
		return dto.RedeemResponse{}, err
	}
	redeemResponse := dto.RedeemResponse{
		ID:         redeem.ID,
		PointSpent: redeem.PointSpent,
		CreatedAt:  redeem.CreatedAt,
		DeletedAt:  redeem.DeletedAt,
		Reward: dto.RewardResponse{
			ID:            redeem.Reward.ID,
			Name:          redeem.Reward.Name,
			Description:   redeem.Reward.Description,
			RequiredPoint: redeem.Reward.RequiredPoint,
			ValidUntil:    redeem.Reward.ValidUntil,
			CreatedAt:     redeem.Reward.CreatedAt,
			UpdatedAt:     redeem.Reward.UpdatedAt,
			DeletedAt:     redeem.Reward.DeletedAt,
		},
	}
	return redeemResponse, nil
}

func (ds redeemServiceImpl) CreateRedeem(ctx context.Context, req dto.RedeemRequest, userID uuid.UUID) error {
	/* TODO:
	- Check point amount.
	- If enough, decrease. if not enough, return error.
	- if enough, create.
	*/
	reward, err := ds.rr.FindRewardByID(ctx, req.RewardID)
	if err != nil {
		return err
	}
	user, err := ds.ur.GetUserByID(ctx, userID)
	if err != nil {
		return err
	}
	sub := int(user.UserCoin.Amount) - int(reward.RequiredPoint)
	if sub < 0 {
		return errors.New("not enough point")
	}

	// proceed to creation of redeem
	redeem := entity.Redeem{
		RewardID:   req.RewardID,
		UserID:     userID,
		PointSpent: reward.RequiredPoint,
	}

	err = ds.dr.CreateRedeem(ctx, &redeem)
	if err != nil {
		return err
	}
	pointUpdate := entity.UserCoin{
		Amount: int64(sub),
	}
	err = ds.pr.UpdatePoint(ctx, user.UserCoin.ID, pointUpdate)
	return err
}

func (ds redeemServiceImpl) UpdateRedeem(ctx context.Context, req dto.RedeemUpdateRequest, id uint64) error {
	redeem := entity.Redeem{
		RewardID:   req.RewardID,
		PointSpent: req.PointSpent,
	}
	err := ds.dr.UpdateRedeem(ctx, &redeem, id)
	if err != nil {
		return err
	}
	return nil
}

func (ds redeemServiceImpl) DeleteRedeem(ctx context.Context, redeemID uint64) error {
	err := ds.dr.DeleteRedeem(ctx, redeemID)
	if err != nil {
		return err
	}
	return nil
}
