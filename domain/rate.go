package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Rate struct {
	RateID       uuid.UUID      `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	PricePerHour      float32         `json:"price_hour"`
	PricePerMinute      float32   `json:"price_minute"`
	OverstayRatePerMinute      int8     `json:"overstay_rate_minute"`
	RateDescription      string     `json:"rate_description"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type RateRepository interface {
	Create(c context.Context, rate Rate) error
	Fetch(c context.Context) ([]Rate, error)
	FetchById(c context.Context, id int) (Rate, error)
	Update(c context.Context, updatedRate Rate) error
	Delete(c context.Context, id int) error
}
