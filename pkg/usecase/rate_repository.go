package usecase

import (
	"context"
	"gorm-template/bootstrap"
	"gorm-template/domain"
)

type RateUseCase struct{}

func (eu *RateUseCase) Create(c context.Context, rate domain.Rate) error {
	db := bootstrap.DB
	err := db.Create(&rate)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (eu *RateUseCase) Fetch(c context.Context) ([]domain.Rate, error) {
	db := bootstrap.DB
	entity := []domain.Rate{}
	err := db.Find(&entity)
	if err.Error != nil {
		return nil, err.Error
	}
	return entity, nil
}

func (eu *RateUseCase) FetchById(c context.Context, id int) (domain.Rate, error) {
	db := bootstrap.DB
	pedido := domain.Rate{}
	err := db.Where("id = ?", id).First(&pedido)
	if err.Error != nil {
		return domain.Rate{}, err.Error
	}
	return pedido, nil
}

func (eu *RateUseCase) Update(c context.Context, updatedrate domain.Rate) error {
	db := bootstrap.DB
	if err := db.Model(&updatedrate).
		Omit("deleted_at", "created_at").
		Updates(updatedrate).Error; err != nil {
		return err
	}
	return nil
}

func (eu *RateUseCase) Delete(c context.Context, id int) error {
	db := bootstrap.DB
	err := db.Where("id = ?", id).Delete(&domain.Rate{})
	if err.Error != nil {
		return err.Error
	}
	return nil
}
