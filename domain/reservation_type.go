package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ReservationType struct {
	ReservationTypeID       uuid.UUID      `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	RTName      string         `json:"rt_name"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type ReservationTypeRepository interface {
	Create(c context.Context, reservationType ReservationType) error
	Fetch(c context.Context) ([]ReservationType, error)
	FetchById(c context.Context, id int) (ReservationType, error)
	Update(c context.Context, updatedReservationType ReservationType) error
	Delete(c context.Context, id int) error
}
