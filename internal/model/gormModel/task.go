package gormModel

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Task struct {
	ID        uuid.UUID `gorm:"primarykey; type:uuid; default:uuid_generate_v4()"`
	Name      string    `gorm:"type:varchar(100);"`
	Done      bool      `gorm:"default:false"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
