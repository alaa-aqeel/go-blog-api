package database

import "gorm.io/gorm"

// Transaction executes a transaction. If the given function returns an error, the transaction
// is rolled back. Otherwise it is automatically committed before `Transaction()` returns.
//
// The Gorm DB associated with this session is injected into the context as a value so `session.DB()`
// can be used to retrieve it.
func WithTransaction(callback func(*gorm.DB) error) error {
	tx := DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	err := callback(tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
