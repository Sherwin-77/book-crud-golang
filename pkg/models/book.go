package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Book struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primary_key"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (b *Book) BeforeCreate(tx *gorm.DB) error {
	var err error
	if b.ID == uuid.Nil {
		b.ID, err = uuid.NewV7()
	}

	return err
}
