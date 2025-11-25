package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Payment struct {
	PaymentID       uuid.UUID      `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	PaymentDate      time.Time         `json:"payment_date"`
	PaymentAmmount      float32   `json:"payment_ammount"`
	PaymentMethod      string     `json:"payment_method"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type PaymentRepository interface {
	Create(c context.Context, payment Payment) error
	Fetch(c context.Context) ([]Payment, error)
	FetchById(c context.Context, id int) (Payment, error)
	Update(c context.Context, updatedPayment Payment) error
	Delete(c context.Context, id int) error
}
