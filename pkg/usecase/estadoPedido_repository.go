package usecase

import (
	"context"

	"gorm-template/bootstrap"
	"gorm-template/domain"
)

type EstadoPedidoUseCase struct{}

func (eu *EstadoPedidoUseCase) Create(c context.Context, estadopedido domain.EstadoPedido) error {
	db := bootstrap.DB
	err := db.Create(&estadopedido)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (eu *EstadoPedidoUseCase) Fetch(c context.Context) ([]domain.EstadoPedido, error) {
	db := bootstrap.DB
	entity := []domain.EstadoPedido{}
	err := db.Find(&entity)
	if err.Error != nil {
		return nil, err.Error
	}
	return entity, nil
}

func (eu *EstadoPedidoUseCase) FetchById(c context.Context, id int) (domain.EstadoPedido, error) {
	db := bootstrap.DB
	pedido := domain.EstadoPedido{}
	err := db.Where("id = ?", id).First(&pedido)
	if err.Error != nil {
		return domain.EstadoPedido{}, err.Error
	}
	return pedido, nil
}

func (eu *EstadoPedidoUseCase) Update(c context.Context, updatedEstadoPedido domain.EstadoPedido) error {
	db := bootstrap.DB
	if err := db.Model(&updatedEstadoPedido).
		Omit("deleted_at", "created_at").
		Updates(updatedEstadoPedido).Error; err != nil {
		return err
	}
	return nil
}

func (eu *EstadoPedidoUseCase) Delete(c context.Context, id int) error {
	db := bootstrap.DB
	err := db.Where("id = ?", id).Delete(&domain.EstadoPedido{})
	if err.Error != nil {
		return err.Error
	}
	return nil
}
