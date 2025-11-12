package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Pedido struct {
	ID        uuid.UUID      `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	NumeroPedido      int         `json:"number"`
	FechaPedido      time.Time   `json:"date"`
	FechaEntrega    time.Time   `json:"delivery_date"`
	DireccionEntrega string      `json:"delivery_address"`
	TotalPedido	float64     `json:"total"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type PedidoRepository interface {
	Create(c context.Context, pedido Pedido) error
	Fetch(c context.Context) ([]Pedido, error)
	FetchById(c context.Context, id int) (Pedido, error)
	Update(c context.Context, updatedPedido Pedido) error
	Delete(c context.Context, id int) error
}
