package usecase

import (
	"context"
	"gorm-template/bootstrap"
	"gorm-template/domain"
)

type ReservationTypeUseCase struct{}

func (eu *ReservationTypeUseCase) Create(c context.Context, reservationType domain.ReservationType) error {
	db := bootstrap.DB
	err := db.Create(&reservationType)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (eu *ReservationTypeUseCase) Fetch(c context.Context) ([]domain.ReservationType, error) {
	db := bootstrap.DB
	entity := []domain.ReservationType{}
	err := db.Find(&entity)
	if err.Error != nil {
		return nil, err.Error
	}
	return entity, nil
}

func (eu *ReservationTypeUseCase) FetchById(c context.Context, id int) (domain.ReservationType, error) {
	db := bootstrap.DB
	pedido := domain.ReservationType{}
	err := db.Where("id = ?", id).First(&pedido)
	if err.Error != nil {
		return domain.ReservationType{}, err.Error
	}
	return pedido, nil
}

func (eu *ReservationTypeUseCase) Update(c context.Context, updatedreservationType domain.ReservationType) error {
	db := bootstrap.DB
	if err := db.Model(&updatedreservationType).
		Omit("deleted_at", "created_at").
		Updates(updatedreservationType).Error; err != nil {
		return err
	}
	return nil
}

func (eu *ReservationTypeUseCase) Delete(c context.Context, id int) error {
	db := bootstrap.DB
	err := db.Where("id = ?", id).Delete(&domain.ReservationType{})
	if err.Error != nil {
		return err.Error
	}
	return nil
}
