package usecase

import (
	"context"

	"gorm-template/bootstrap"
	"gorm-template/domain"
)

type PedidoUseCase struct{}

func (eu *PedidoUseCase) Create(c context.Context, pedido domain.Pedido) error {
	db := bootstrap.DB
	err := db.Create(&pedido)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (eu *PedidoUseCase) Fetch(c context.Context) ([]domain.Pedido, error) {
	db := bootstrap.DB
	entity := []domain.Pedido{}
	err := db.Find(&entity)
	if err.Error != nil {
		return nil, err.Error
	}
	return entity, nil
}

func (eu *PedidoUseCase) FetchById(c context.Context, id int) (domain.Pedido, error) {
	db := bootstrap.DB
	pedido := domain.Pedido{}
	err := db.Where("id = ?", id).First(&pedido)
	if err.Error != nil {
		return domain.Pedido{}, err.Error
	}
	return pedido, nil
}

func (eu *PedidoUseCase) Update(c context.Context, updatedPedido domain.Pedido) error {
	db := bootstrap.DB
	if err := db.Model(&updatedPedido).
		Omit("deleted_at", "created_at").
		Updates(updatedPedido).Error; err != nil {
		return err
	}
	return nil
}

func (eu *PedidoUseCase) Delete(c context.Context, id int) error {
	db := bootstrap.DB
	err := db.Where("id = ?", id).Delete(&domain.Pedido{})
	if err.Error != nil {
		return err.Error
	}
	return nil
}
