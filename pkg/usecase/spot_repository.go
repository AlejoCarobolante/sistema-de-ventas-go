package usecase

import (
	"context"
	"gorm-template/bootstrap"
	"gorm-template/domain"
)

type SpotUseCase struct{}

func (eu *SpotUseCase) Create(c context.Context, spot domain.Spot) error {
	db := bootstrap.DB
	err := db.Create(&spot)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (eu *SpotUseCase) Fetch(c context.Context) ([]domain.Spot, error) {
	db := bootstrap.DB
	entity := []domain.Spot{}
	err := db.Find(&entity)
	if err.Error != nil {
		return nil, err.Error
	}
	return entity, nil
}

func (eu *SpotUseCase) FetchById(c context.Context, id int) (domain.Spot, error) {
	db := bootstrap.DB
	pedido := domain.Spot{}
	err := db.Where("id = ?", id).First(&pedido)
	if err.Error != nil {
		return domain.Spot{}, err.Error
	}
	return pedido, nil
}

func (eu *SpotUseCase) Update(c context.Context, updatedspot domain.Spot) error {
	db := bootstrap.DB
	if err := db.Model(&updatedspot).
		Omit("deleted_at", "created_at").
		Updates(updatedspot).Error; err != nil {
		return err
	}
	return nil
}

func (eu *SpotUseCase) Delete(c context.Context, id int) error {
	db := bootstrap.DB
	err := db.Where("id = ?", id).Delete(&domain.Spot{})
	if err.Error != nil {
		return err.Error
	}
	return nil
}
