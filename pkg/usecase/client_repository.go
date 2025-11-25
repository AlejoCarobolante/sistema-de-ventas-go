package usecase

import (
	"context"
	"gorm-template/bootstrap"
	"gorm-template/domain"
)

type ClientUseCase struct{}

func (cu *ClientUseCase) Create(c context.Context, client domain.Client) error {
	db := bootstrap.DB
	err := db.Create(&client)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (cu *ClientUseCase) Fetch(c context.Context) ([]domain.Client, error) {
	db := bootstrap.DB
	clients := []domain.Client{}
	err := db.Find(&clients)
	if err.Error != nil {
		return nil, err.Error
	}
	return clients, nil
}

func (cu *ClientUseCase) FetchByID(c context.Context, id string) (domain.Client, error) {
	db := bootstrap.DB
	client := domain.Client{}
	err := db.Where("id=?", id).First(&client)
	if err != nil {
		return domain.Client{}, err.Error
	}
	return client, nil
}

func (cu *ClientUseCase) Update(c context.Context, updateClient domain.Client) error {
	db := bootstrap.DB
	if err := db.Model(&updateClient).
		Omit("delete_at", "create_at").
		Updates(updateClient).Error; err != nil {
		return err
	}
	return nil
}

func (cu *ClientUseCase) Delete(c context.Context, id string) error {
	db := bootstrap.DB
	err := db.Where("id=?", id).Delete(&domain.Client{})
	if err != nil {
		return err.Error
	}
	return nil
}
