package repository

import (
	"context"

	"gorm.io/gorm"
)

type DefaultRepositoryInterface interface {
	Welcome(ctx context.Context) string
	WithTx(tx *gorm.DB) DefaultRepositoryInterface
}
