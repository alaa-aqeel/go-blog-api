package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID      `gorm:"type:uuid;primaryKey;" json:"id"`
	Name      string         `gorm:"type:varchar(100);not null"`
	Username  string         `gorm:"type:varchar(100);uniqueIndex;not null"`
	Password  string         `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.ID = uuid.New()

	return nil
}
