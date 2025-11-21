package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SpotType struct {
	SpotTypeID       uuid.UUID      `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	Name      string         `json:"name"`
	MinArea      float32   `json:"min_area"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type SpotTypeRepository interface {
	Create(c context.Context, spotType SpotType) error
	Fetch(c context.Context) ([]SpotType, error)
	FetchById(c context.Context, id int) (SpotType, error)
	Update(c context.Context, updatedSpotType SpotType) error
	Delete(c context.Context, id int) error
}
