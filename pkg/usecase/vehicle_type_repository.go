package usecase

import (
	"context"
	"gorm-template/bootstrap"
	"gorm-template/domain"
)

type VehicleTypeUseCase struct{}

func (eu *VehicleTypeUseCase) Create(c context.Context, vehicleType domain.VehicleType) error {
	db := bootstrap.DB
	err := db.Create(&vehicleType)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (eu *VehicleTypeUseCase) Fetch(c context.Context) ([]domain.VehicleType, error) {
	db := bootstrap.DB
	entity := []domain.VehicleType{}
	err := db.Find(&entity)
	if err.Error != nil {
		return nil, err.Error
	}
	return entity, nil
}

func (eu *VehicleTypeUseCase) FetchById(c context.Context, id int) (domain.VehicleType, error) {
	db := bootstrap.DB
	pedido := domain.VehicleType{}
	err := db.Where("id = ?", id).First(&pedido)
	if err.Error != nil {
		return domain.VehicleType{}, err.Error
	}
	return pedido, nil
}

func (eu *VehicleTypeUseCase) Update(c context.Context, updatedvehicleType domain.VehicleType) error {
	db := bootstrap.DB
	if err := db.Model(&updatedvehicleType).
		Omit("deleted_at", "created_at").
		Updates(updatedvehicleType).Error; err != nil {
		return err
	}
	return nil
}

func (eu *VehicleTypeUseCase) Delete(c context.Context, id int) error {
	db := bootstrap.DB
	err := db.Where("id = ?", id).Delete(&domain.VehicleType{})
	if err.Error != nil {
		return err.Error
	}
	return nil
}
