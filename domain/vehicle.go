package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Vehicle struct {
	VehicleID       uuid.UUID      `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	LicensePlate      string         `json:"plate"`
	Maker      string   `json:"maker"`
	Model      string     `json:"model"`
	Color      string     `json:"color"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type VehicleRepository interface {
	Create(c context.Context, vehicle Vehicle) error
	Fetch(c context.Context) ([]Vehicle, error)
	FetchById(c context.Context, id int) (Vehicle, error)
	Update(c context.Context, updatedVehicle Vehicle) error
	Delete(c context.Context, id int) error
}
