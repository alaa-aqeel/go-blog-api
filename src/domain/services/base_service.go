package services

import "gorm.io/gorm"

type BaseService struct {
	Db *gorm.DB
}
