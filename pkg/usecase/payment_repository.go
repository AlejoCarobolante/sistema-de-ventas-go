package usecase

import (
	"context"
	"gorm-template/bootstrap"
	"gorm-template/domain"
)

type PaymentUseCase struct{}

func (eu *PaymentUseCase) Create(c context.Context, payment domain.Payment) error {
	db := bootstrap.DB
	err := db.Create(&payment)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (eu *PaymentUseCase) Fetch(c context.Context) ([]domain.Payment, error) {
	db := bootstrap.DB
	entity := []domain.Payment{}
	err := db.Find(&entity)
	if err.Error != nil {
		return nil, err.Error
	}
	return entity, nil
}

func (eu *PaymentUseCase) FetchById(c context.Context, id int) (domain.Payment, error) {
	db := bootstrap.DB
	pedido := domain.Payment{}
	err := db.Where("id = ?", id).First(&pedido)
	if err.Error != nil {
		return domain.Payment{}, err.Error
	}
	return pedido, nil
}

func (eu *PaymentUseCase) Update(c context.Context, updatedpayment domain.Payment) error {
	db := bootstrap.DB
	if err := db.Model(&updatedpayment).
		Omit("deleted_at", "created_at").
		Updates(updatedpayment).Error; err != nil {
		return err
	}
	return nil
}

func (eu *PaymentUseCase) Delete(c context.Context, id int) error {
	db := bootstrap.DB
	err := db.Where("id = ?", id).Delete(&domain.Payment{})
	if err.Error != nil {
		return err.Error
	}
	return nil
}
