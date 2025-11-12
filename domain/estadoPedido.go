package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EstadoPedido struct {
	ID        uuid.UUID      `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	NombreEP      string         `json:"nameEP"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type EstadoPedidoRepository interface {
	Create(c context.Context, estadopedido EstadoPedido) error
	Fetch(c context.Context) ([]EstadoPedido, error)
	FetchById(c context.Context, id int) (EstadoPedido, error)
	Update(c context.Context, updatedEstadoPedido EstadoPedido) error
	Delete(c context.Context, id int) error
}
