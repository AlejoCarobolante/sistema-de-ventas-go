package usecase

import (
	"context"
	"gorm-template/bootstrap"
	"gorm-template/domain"
)

type ReservationStatusUseCase struct{}

func (eu *ReservationStatusUseCase) Create(c context.Context, reservationStatus domain.ReservationStatus) error {
	db := bootstrap.DB
	err := db.Create(&reservationStatus)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (eu *ReservationStatusUseCase) Fetch(c context.Context) ([]domain.ReservationStatus, error) {
	db := bootstrap.DB
	entity := []domain.ReservationStatus{}
	err := db.Find(&entity)
	if err.Error != nil {
		return nil, err.Error
	}
	return entity, nil
}

func (eu *ReservationStatusUseCase) FetchById(c context.Context, id int) (domain.ReservationStatus, error) {
	db := bootstrap.DB
	pedido := domain.ReservationStatus{}
	err := db.Where("id = ?", id).First(&pedido)
	if err.Error != nil {
		return domain.ReservationStatus{}, err.Error
	}
	return pedido, nil
}

func (eu *ReservationStatusUseCase) Update(c context.Context, updatedreservationStatus domain.ReservationStatus) error {
	db := bootstrap.DB
	if err := db.Model(&updatedreservationStatus).
		Omit("deleted_at", "created_at").
		Updates(updatedreservationStatus).Error; err != nil {
		return err
	}
	return nil
}

func (eu *ReservationStatusUseCase) Delete(c context.Context, id int) error {
	db := bootstrap.DB
	err := db.Where("id = ?", id).Delete(&domain.ReservationStatus{})
	if err.Error != nil {
		return err.Error
	}
	return nil
}
