package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CompatibilityRule struct {
	CompatibilityRuleID       uuid.UUID      `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	IsCompatible      bool         `json:"is_compatible"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type CompatibilityRuleRepository interface {
	Create(c context.Context, compatibilityRule CompatibilityRule) error
	Fetch(c context.Context) ([]CompatibilityRule, error)
	FetchById(c context.Context, id int) (CompatibilityRule, error)
	Update(c context.Context, updatedCompatibilityRule CompatibilityRule) error
	Delete(c context.Context, id int) error
}
