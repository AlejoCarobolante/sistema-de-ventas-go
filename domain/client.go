package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Client struct {
	ID        uuid.UUID
	Name      string
	Email     string
	Balance   int
	Create_at time.Time
	Update_at time.Time
	Delete_at time.Time
}

type ClientRepository interface {
	Create(c context.Context, client Client) error
	Fetch(c context.Context) ([]Client, error)
	FetchByID(c context.Context, id string) (Client, error)
	Update(c context.Context, updateClient Client) error
	Delete(c context.Context, id string) error
}
