package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ReservationStatus struct {
	ReservationStatusID       uuid.UUID      `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	RSName      string         `json:"rs_name"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type ReservationStatusRepository interface {
	Create(c context.Context, payment ReservationStatus) error
	Fetch(c context.Context) ([]ReservationStatus, error)
	FetchById(c context.Context, id int) (ReservationStatus, error)
	Update(c context.Context, updatedReservationStatus ReservationStatus) error
	Delete(c context.Context, id int) error
}
