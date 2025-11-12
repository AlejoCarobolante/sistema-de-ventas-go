package usecase

import (
	"context"

	"gorm-template/bootstrap"
	"gorm-template/domain"
)

type Entity1UseCase struct{}

func (eu *Entity1UseCase) Create(c context.Context, entity1 domain.Entity1) error {
	db := bootstrap.DB
	err := db.Create(&entity1)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (eu *Entity1UseCase) Fetch(c context.Context) ([]domain.Entity1, error) {
	db := bootstrap.DB
	entity := []domain.Entity1{}
	err := db.Find(&entity)
	if err.Error != nil {
		return nil, err.Error
	}
	return entity, nil
}

func (eu *Entity1UseCase) FetchById(c context.Context, id int) (domain.Entity1, error) {
	db := bootstrap.DB
	entity1 := domain.Entity1{}
	err := db.Where("id = ?", id).First(&entity1)
	if err.Error != nil {
		return domain.Entity1{}, err.Error
	}
	return entity1, nil
}

func (eu *Entity1UseCase) Update(c context.Context, updatedEntity1 domain.Entity1) error {
	db := bootstrap.DB
	if err := db.Model(&updatedEntity1).
		Omit("deleted_at", "created_at").
		Updates(updatedEntity1).Error; err != nil {
		return err
	}
	return nil
}

func (eu *Entity1UseCase) Delete(c context.Context, id int) error {
	db := bootstrap.DB
	err := db.Where("id = ?", id).Delete(&domain.Entity1{})
	if err.Error != nil {
		return err.Error
	}
	return nil
}
