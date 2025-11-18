package usecase

import (
	"context"

	"gorm-template/bootstrap"
	"gorm-template/domain"
)

type ReservationUseCase struct{}

func (eu *ReservationUseCase) Create(c context.Context, reservation domain.Reservation) error {
	db := bootstrap.DB
	err := db.Create(&reservation)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (eu *ReservationUseCase) Fetch(c context.Context) ([]domain.Reservation, error) {
	db := bootstrap.DB
	entity := []domain.Reservation{}
	err := db.Find(&entity)
	if err.Error != nil {
		return nil, err.Error
	}
	return entity, nil
}

func (eu *ReservationUseCase) FetchById(c context.Context, id int) (domain.Reservation, error) {
	db := bootstrap.DB
	pedido := domain.Reservation{}
	err := db.Where("id = ?", id).First(&pedido)
	if err.Error != nil {
		return domain.Reservation{}, err.Error
	}
	return pedido, nil
}

func (eu *ReservationUseCase) Update(c context.Context, updatedReservation domain.Reservation) error {
	db := bootstrap.DB
	if err := db.Model(&updatedReservation).
		Omit("deleted_at", "created_at").
		Updates(updatedReservation).Error; err != nil {
		return err
	}
	return nil
}

func (eu *ReservationUseCase) Delete(c context.Context, id int) error {
	db := bootstrap.DB
	err := db.Where("id = ?", id).Delete(&domain.Reservation{})
	if err.Error != nil {
		return err.Error
	}
	return nil
}
