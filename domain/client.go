package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Client struct {
	ClientID        uuid.UUID      `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	ClientName      string         `json:"client_name"`
	CreatedAt time.Time      `json:"created_at"`
	Email      string      `json:"email"`
	CellPhone	string     `json:"phone"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type ClientRepository interface {
	Create(c context.Context, client Client) error
	Fetch(c context.Context) ([]Client, error)
	FetchById(c context.Context, id int) (Client, error)
	Update(c context.Context, updatedClient Client) error
	Delete(c context.Context, id int) error
}
