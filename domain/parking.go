package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Parking struct {
	ParkingID       uuid.UUID      `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	Name      string         `json:"parking_name"`
	Adress      string   `json:"parking_adress"`
	Capacity      int64     `json:"capacity"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type ParkingRepository interface {
	Create(c context.Context, parking Parking) error
	Fetch(c context.Context) ([]Parking, error)
	FetchById(c context.Context, id int) (Parking, error)
	Update(c context.Context, updatedParking Parking) error
	Delete(c context.Context, id int) error
}
