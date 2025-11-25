package usecase

import (
	"context"
	"gorm-template/bootstrap"
	"gorm-template/domain"
)

type CompatibilityRuleUseCase struct{}

func (cru *CompatibilityRuleUseCase) Create(c context.Context, compatibilityRule domain.CompatibilityRule) error {
	db := bootstrap.DB
	err := db.Create(&compatibilityRule)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (cru *CompatibilityRuleUseCase) Fetch(c context.Context) ([]domain.CompatibilityRule, error) {
	db := bootstrap.DB
	entity := []domain.CompatibilityRule{}
	err := db.Find(&entity)
	if err.Error != nil {
		return nil, err.Error
	}
	return entity, nil
}

func (cru *CompatibilityRuleUseCase) FetchById(c context.Context, id int) (domain.CompatibilityRule, error) {
	db := bootstrap.DB
	pedido := domain.CompatibilityRule{}
	err := db.Where("id = ?", id).First(&pedido)
	if err.Error != nil {
		return domain.CompatibilityRule{}, err.Error
	}
	return pedido, nil
}

func (cru *CompatibilityRuleUseCase) Update(c context.Context, updatedcompatibilityRule domain.CompatibilityRule) error {
	db := bootstrap.DB
	if err := db.Model(&updatedcompatibilityRule).
		Omit("deleted_at", "created_at").
		Updates(updatedcompatibilityRule).Error; err != nil {
		return err
	}
	return nil
}

func (cru *CompatibilityRuleUseCase) Delete(c context.Context, id int) error {
	db := bootstrap.DB
	err := db.Where("id = ?", id).Delete(&domain.CompatibilityRule{})
	if err.Error != nil {
		return err.Error
	}
	return nil
}
