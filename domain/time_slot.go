package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TimeSlot struct {
	TimeSlotID       uuid.UUID      `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	DayOfWeek      string         `json:"week_day"`
	StartTime      time.Time   `json:"start_time"`
	EndTime      time.Time   `json:"end_time"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type TimeSlotRepository interface {
	Create(c context.Context, timeSlot TimeSlot) error
	Fetch(c context.Context) ([]TimeSlot, error)
	FetchById(c context.Context, id int) (TimeSlot, error)
	Update(c context.Context, updatedTimeSlot TimeSlot) error
	Delete(c context.Context, id int) error
}
