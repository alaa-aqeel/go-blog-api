package database

import (
	"context"

	"gorm.io/gorm"
)

// dbKey the key used to store the database in the context.
type DbKey struct{}

// DB returns the Gorm instance stored in the given context. Returns the given fallback
// if no Gorm DB could be found in the context.
func GetDb(ctx context.Context, default_db *gorm.DB) *gorm.DB {
	db := ctx.Value(DbKey{})
	if db == nil {
		return default_db
	}
	return db.(*gorm.DB)
}
