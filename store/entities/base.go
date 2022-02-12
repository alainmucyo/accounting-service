package entities

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"time"
)

type Base struct {
	ID        string     `sql:"type:uuid;primary_key;default:uuid_generate_v4()"`
	CreatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}

// BeforeCreate FYI: It should return an error, it causes errors when not returning anything
func (base *Base) BeforeCreate(tx *gorm.DB) error {
	id := uuid.NewV4()
	tx.Statement.SetColumn("ID", id.String())
	return nil
}
