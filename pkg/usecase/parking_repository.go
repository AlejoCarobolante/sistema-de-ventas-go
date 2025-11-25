package usecase

import (
	"context"
	"gorm-template/bootstrap"
	"gorm-template/domain"
)

type ParkingUseCase struct{}

func (eu *ParkingUseCase) Create(c context.Context, parking domain.Parking) error {
	db := bootstrap.DB
	err := db.Create(&parking)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (eu *ParkingUseCase) Fetch(c context.Context) ([]domain.Parking, error) {
	db := bootstrap.DB
	entity := []domain.Parking{}
	err := db.Find(&entity)
	if err.Error != nil {
		return nil, err.Error
	}
	return entity, nil
}

func (eu *ParkingUseCase) FetchById(c context.Context, id int) (domain.Parking, error) {
	db := bootstrap.DB
	pedido := domain.Parking{}
	err := db.Where("id = ?", id).First(&pedido)
	if err.Error != nil {
		return domain.Parking{}, err.Error
	}
	return pedido, nil
}

func (eu *ParkingUseCase) Update(c context.Context, updatedparking domain.Parking) error {
	db := bootstrap.DB
	if err := db.Model(&updatedparking).
		Omit("deleted_at", "created_at").
		Updates(updatedparking).Error; err != nil {
		return err
	}
	return nil
}

func (eu *ParkingUseCase) Delete(c context.Context, id int) error {
	db := bootstrap.DB
	err := db.Where("id = ?", id).Delete(&domain.Parking{})
	if err.Error != nil {
		return err.Error
	}
	return nil
}
