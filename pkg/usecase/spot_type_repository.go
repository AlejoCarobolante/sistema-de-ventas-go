package usecase

import (
	"context"
	"gorm-template/bootstrap"
	"gorm-template/domain"
)

type SpotTypeUseCase struct{}

func (eu *SpotTypeUseCase) Create(c context.Context, spotType domain.SpotType) error {
	db := bootstrap.DB
	err := db.Create(&spotType)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (eu *SpotTypeUseCase) Fetch(c context.Context) ([]domain.SpotType, error) {
	db := bootstrap.DB
	entity := []domain.SpotType{}
	err := db.Find(&entity)
	if err.Error != nil {
		return nil, err.Error
	}
	return entity, nil
}

func (eu *SpotTypeUseCase) FetchById(c context.Context, id int) (domain.SpotType, error) {
	db := bootstrap.DB
	pedido := domain.SpotType{}
	err := db.Where("id = ?", id).First(&pedido)
	if err.Error != nil {
		return domain.SpotType{}, err.Error
	}
	return pedido, nil
}

func (eu *SpotTypeUseCase) Update(c context.Context, updatedspotType domain.SpotType) error {
	db := bootstrap.DB
	if err := db.Model(&updatedspotType).
		Omit("deleted_at", "created_at").
		Updates(updatedspotType).Error; err != nil {
		return err
	}
	return nil
}

func (eu *SpotTypeUseCase) Delete(c context.Context, id int) error {
	db := bootstrap.DB
	err := db.Where("id = ?", id).Delete(&domain.SpotType{})
	if err.Error != nil {
		return err.Error
	}
	return nil
}
