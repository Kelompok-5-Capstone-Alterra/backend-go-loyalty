package faqService

import (
	"backend-go-loyalty/internal/dto"
	"backend-go-loyalty/internal/entity"
	faqRepository "backend-go-loyalty/internal/repository/faq"
	"context"
	"time"
)

type IFaqService interface {
	GetAllFAQ(ctx context.Context, keyword string) (dto.FAQResponses, error)
	GetFAQByID(ctx context.Context, id uint64) (dto.FAQResponse, error)
	CreateFAQ(ctx context.Context, req dto.FAQRequest) error
	UpdateFAQ(ctx context.Context, req dto.FAQUpdateRequest, id uint64) error
	DeleteFAQ(ctx context.Context, id uint64) error
}

type faqService struct {
	fr faqRepository.IFaqRepository
}

func NewFAQService(fr faqRepository.IFaqRepository) faqService {
	return faqService{
		fr: fr,
	}
}

func (fs faqService) GetAllFAQ(ctx context.Context, keyword string) (dto.FAQResponses, error) {
	data, err := fs.fr.GetAllFAQByKeyword(ctx, keyword)
	if err != nil {
		return dto.FAQResponses{}, err
	}
	var faqs dto.FAQResponses
	for _, val := range data {
		faq := dto.FAQResponse{
			ID:        val.ID,
			Question:  val.Question,
			Answer:    val.Answer,
			CreatedAt: val.CreatedAt,
			UpdatedAt: val.UpdatedAt,
			DeletedAt: val.DeletedAt,
		}
		faqs = append(faqs, faq)
	}
	return faqs, nil
}
func (fs faqService) GetFAQByID(ctx context.Context, id uint64) (dto.FAQResponse, error) {
	data, err := fs.fr.GetFAQByID(ctx, id)
	if err != nil {
		return dto.FAQResponse{}, err
	}
	faq := dto.FAQResponse{
		ID:        data.ID,
		Question:  data.Question,
		Answer:    data.Answer,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
		DeletedAt: data.DeletedAt,
	}
	return faq, nil
}
func (fs faqService) CreateFAQ(ctx context.Context, req dto.FAQRequest) error {
	faq := entity.FAQ{
		Question:  req.Question,
		Answer:    req.Answer,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := fs.fr.CreateFAQ(ctx, faq)
	return err
}
func (fs faqService) UpdateFAQ(ctx context.Context, req dto.FAQUpdateRequest, id uint64) error {
	faq := entity.FAQ{
		Question:  req.Question,
		Answer:    req.Answer,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := fs.fr.UpdateFAQ(ctx, faq, id)
	return err
}
func (fs faqService) DeleteFAQ(ctx context.Context, id uint64) error {
	err := fs.fr.DeleteFAQ(ctx, id)
	return err
}
