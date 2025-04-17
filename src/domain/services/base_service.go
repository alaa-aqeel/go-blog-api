package services

import "gorm.io/gorm"

type BaseService[T any] struct {
	Db *gorm.DB
}

func (b *BaseService[T]) Query() *gorm.DB {
	var t T
	return b.Db.Model(&t)
}
