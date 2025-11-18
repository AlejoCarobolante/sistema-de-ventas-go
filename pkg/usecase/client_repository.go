package usecase

import (
	"context"
	"gorm-template/bootstrap"
	"gorm-template/domain"
)

type ClientUseCase struct{}

func (eu *ClientUseCase) Create(c context.Context, client domain.Client) error {
	db := bootstrap.DB
	err := db.Create(&client)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (eu *ClientUseCase) Fetch(c context.Context) ([]domain.Client, error) {
	db := bootstrap.DB
	entity := []domain.Client{}
	err := db.Find(&entity)
	if err.Error != nil {
		return nil, err.Error
	}
	return entity, nil
}

func (eu *ClientUseCase) FetchById(c context.Context, id int) (domain.Client, error) {
	db := bootstrap.DB
	pedido := domain.Client{}
	err := db.Where("id = ?", id).First(&pedido)
	if err.Error != nil {
		return domain.Client{}, err.Error
	}
	return pedido, nil
}

func (eu *ClientUseCase) Update(c context.Context, updatedclient domain.Client) error {
	db := bootstrap.DB
	if err := db.Model(&updatedclient).
		Omit("deleted_at", "created_at").
		Updates(updatedclient).Error; err != nil {
		return err
	}
	return nil
}

func (eu *ClientUseCase) Delete(c context.Context, id int) error {
	db := bootstrap.DB
	err := db.Where("id = ?", id).Delete(&domain.Client{})
	if err.Error != nil {
		return err.Error
	}
	return nil
}
