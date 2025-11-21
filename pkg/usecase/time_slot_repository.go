package usecase

import (
	"context"
	"gorm-template/bootstrap"
	"gorm-template/domain"
)

type TimeSlotUseCase struct{}

func (eu *TimeSlotUseCase) Create(c context.Context, timeSlot domain.TimeSlot) error {
	db := bootstrap.DB
	err := db.Create(&timeSlot)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (eu *TimeSlotUseCase) Fetch(c context.Context) ([]domain.TimeSlot, error) {
	db := bootstrap.DB
	entity := []domain.TimeSlot{}
	err := db.Find(&entity)
	if err.Error != nil {
		return nil, err.Error
	}
	return entity, nil
}

func (eu *TimeSlotUseCase) FetchById(c context.Context, id int) (domain.TimeSlot, error) {
	db := bootstrap.DB
	pedido := domain.TimeSlot{}
	err := db.Where("id = ?", id).First(&pedido)
	if err.Error != nil {
		return domain.TimeSlot{}, err.Error
	}
	return pedido, nil
}

func (eu *TimeSlotUseCase) Update(c context.Context, updatedtimeSlot domain.TimeSlot) error {
	db := bootstrap.DB
	if err := db.Model(&updatedtimeSlot).
		Omit("deleted_at", "created_at").
		Updates(updatedtimeSlot).Error; err != nil {
		return err
	}
	return nil
}

func (eu *TimeSlotUseCase) Delete(c context.Context, id int) error {
	db := bootstrap.DB
	err := db.Where("id = ?", id).Delete(&domain.TimeSlot{})
	if err.Error != nil {
		return err.Error
	}
	return nil
}
