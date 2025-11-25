package usecase

import (
	"context"
	"gorm-template/bootstrap"
	"gorm-template/domain"
)

type PenaltyUseCase struct{}

func (eu *PenaltyUseCase) Create(c context.Context, parking domain.Penalty) error {
	db := bootstrap.DB
	err := db.Create(&parking)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (eu *PenaltyUseCase) Fetch(c context.Context) ([]domain.Penalty, error) {
	db := bootstrap.DB
	entity := []domain.Penalty{}
	err := db.Find(&entity)
	if err.Error != nil {
		return nil, err.Error
	}
	return entity, nil
}

func (eu *PenaltyUseCase) FetchById(c context.Context, id int) (domain.Penalty, error) {
	db := bootstrap.DB
	pedido := domain.Penalty{}
	err := db.Where("id = ?", id).First(&pedido)
	if err.Error != nil {
		return domain.Penalty{}, err.Error
	}
	return pedido, nil
}

func (eu *PenaltyUseCase) Update(c context.Context, updatedparking domain.Penalty) error {
	db := bootstrap.DB
	if err := db.Model(&updatedparking).
		Omit("deleted_at", "created_at").
		Updates(updatedparking).Error; err != nil {
		return err
	}
	return nil
}

func (eu *PenaltyUseCase) Delete(c context.Context, id int) error {
	db := bootstrap.DB
	err := db.Where("id = ?", id).Delete(&domain.Penalty{})
	if err.Error != nil {
		return err.Error
	}
	return nil
}
