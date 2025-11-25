package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Penalty struct {
	PenaltyID       uuid.UUID      `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	Reason      string         `json:"reason"`
	DelayMinutes      int32   `json:"delay"`
	PenaltyAmmount      int64     `json:"penalty_ammount"`
	IsPaid      bool     `json:"is_paid"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type PenaltyRepository interface {
	Create(c context.Context, penalty Penalty) error
	Fetch(c context.Context) ([]Penalty, error)
	FetchById(c context.Context, id int) (Penalty, error)
	Update(c context.Context, updatedPenalty Penalty) error
	Delete(c context.Context, id int) error
}
