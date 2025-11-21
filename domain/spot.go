package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Spot struct {
	SpotID       uuid.UUID      `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	Code      string         `json:"code"`
	Level      string   `json:"level"`
	IsAvailable      bool     `json:"available"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type SpotRepository interface {
	Create(c context.Context, spot Spot) error
	Fetch(c context.Context) ([]Spot, error)
	FetchById(c context.Context, id int) (Spot, error)
	Update(c context.Context, updatedSpot Spot) error
	Delete(c context.Context, id int) error
}
