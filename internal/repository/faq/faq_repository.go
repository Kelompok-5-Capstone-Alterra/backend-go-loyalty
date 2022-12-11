package faqRepository

import (
	"backend-go-loyalty/internal/entity"
	"backend-go-loyalty/internal/model"
	"context"

	"gorm.io/gorm"
)

type IFaqRepository interface {
	GetAllFAQByKeyword(ctx context.Context, keyword string) (entity.FAQs, error)
	GetFAQByID(ctx context.Context, id uint64) (entity.FAQ, error)
	CreateFAQ(ctx context.Context, req entity.FAQ) error
	UpdateFAQ(ctx context.Context, req entity.FAQ, id uint64) error
	DeleteFAQ(ctx context.Context, id uint64) error
}

type faqRepository struct {
	db *gorm.DB
}

func NewFAQRepository(db *gorm.DB) faqRepository {
	return faqRepository{
		db: db,
	}
}

func (fr faqRepository) GetAllFAQByKeyword(ctx context.Context, keyword string) (entity.FAQs, error) {
	keyword = "%" + keyword + "%"
	faqs := entity.FAQs{}
	err := fr.db.Where("question LIKE ?", keyword).Find(&faqs).Error
	return faqs, err
}
func (fr faqRepository) GetFAQByID(ctx context.Context, id uint64) (entity.FAQ, error) {
	faq := entity.FAQ{}
	err := fr.db.First(&faq, id).Error
	return faq, err
}
func (fr faqRepository) CreateFAQ(ctx context.Context, req entity.FAQ) error {
	err := fr.db.Create(&req).Error
	return err
}
func (fr faqRepository) UpdateFAQ(ctx context.Context, req entity.FAQ, id uint64) error {
	err := fr.db.Model(&model.FAQ{}).Where("id = ?", id).Updates(req).Error
	return err
}
func (fr faqRepository) DeleteFAQ(ctx context.Context, id uint64) error {
	faq := entity.FAQ{}
	err := fr.db.Delete(&faq, id).Error
	return err
}
