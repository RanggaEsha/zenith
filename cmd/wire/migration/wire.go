//go:build wireinject

package migration

import (
	"github.com/arifai/zenith/internal/model/migration"
	"github.com/google/uuid"
	"github.com/google/wire"
	"gorm.io/gorm"
)

func ProvideMigration(db *gorm.DB, id uuid.UUID) *migration.Migration {
	wire.Build(migration.New)
	return &migration.Migration{}
}