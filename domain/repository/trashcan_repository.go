package repository

import (
	"context"
	domain "gomibakokun_backend/domain/trashcan"
)

type TrashcanRepository interface {
	CreateTrashcan(ctx context.Context, trashcan *domain.Trashcan) error
	GetAllTrashcan(ctx context.Context) ([]*domain.Trashcan, error)
	DeleteTrashcan(ctx context.Context, id string) error
}
