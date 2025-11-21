package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type VehicleType struct {
	VehicleTypeID       uuid.UUID      `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	Name      string         `json:"vehicleType_name"`
	MinLength      float32   `json:"min_length"`
	MaxLength      float32   `json:"max_length"`
	MinWidth      float32   `json:"min_width"`
	MaxWidth      float32   `json:"max_width"`
	MinWeight      float32   `json:"min_weight"`
	MaxWeight      float32   `json:"max_weight"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type VehicleTypeRepository interface {
	Create(c context.Context, vehicleType VehicleType) error
	Fetch(c context.Context) ([]VehicleType, error)
	FetchById(c context.Context, id int) (VehicleType, error)
	Update(c context.Context, updatedVehicleType VehicleType) error
	Delete(c context.Context, id int) error
}
