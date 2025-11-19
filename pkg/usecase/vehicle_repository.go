package usecase

import (
	"context"

	"gorm-template/bootstrap"
	"gorm-template/domain"
)

type VehicleUseCase struct{}

func (eu *VehicleUseCase) Create(c context.Context, vehicle domain.Vehicle) error {
	db := bootstrap.DB
	err := db.Create(&vehicle)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (eu *VehicleUseCase) Fetch(c context.Context) ([]domain.Vehicle, error) {
	db := bootstrap.DB
	entity := []domain.Vehicle{}
	err := db.Find(&entity)
	if err.Error != nil {
		return nil, err.Error
	}
	return entity, nil
}

func (eu *VehicleUseCase) FetchById(c context.Context, id int) (domain.Vehicle, error) {
	db := bootstrap.DB
	pedido := domain.Vehicle{}
	err := db.Where("id = ?", id).First(&pedido)
	if err.Error != nil {
		return domain.Vehicle{}, err.Error
	}
	return pedido, nil
}

func (eu *VehicleUseCase) Update(c context.Context, updatedVehicle domain.Vehicle) error {
	db := bootstrap.DB
	if err := db.Model(&updatedVehicle).
		Omit("deleted_at", "created_at").
		Updates(updatedVehicle).Error; err != nil {
		return err
	}
	return nil
}

func (eu *VehicleUseCase) Delete(c context.Context, id int) error {
	db := bootstrap.DB
	err := db.Where("id = ?", id).Delete(&domain.Vehicle{})
	if err.Error != nil {
		return err.Error
	}
	return nil
}
