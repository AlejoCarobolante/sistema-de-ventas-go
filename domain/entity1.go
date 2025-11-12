package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Entity1 struct {
	ID        uuid.UUID      `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	Name      string         `json:"name"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type Entity1Repository interface {
	Create(c context.Context, entity1 Entity1) error
	Fetch(c context.Context) ([]Entity1, error)
	FetchById(c context.Context, id int) (Entity1, error)
	Update(c context.Context, updatedEntity1 Entity1) error
	Delete(c context.Context, id int) error
}
