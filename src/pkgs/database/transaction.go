package database

import "gorm.io/gorm"

func WithTransaction(callback func(db *gorm.DB) error) error {
	tx := DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	if err := callback(tx); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
