package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Reservation struct {
	ReservationID        uuid.UUID      `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	Start      time.Time         `json:"start"`
	End      time.Time   `json:"end"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type ReservationRepository interface {
	Create(c context.Context, reservation Reservation) error
	Fetch(c context.Context) ([]Reservation, error)
	FetchById(c context.Context, id int) (Reservation, error)
	Update(c context.Context, updatedReservation Reservation) error
	Delete(c context.Context, id int) error
}
