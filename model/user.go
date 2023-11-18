package model

import (
	"github.com/google/uuid"
)

type User struct {
	ID        uint      `gorm:"primaryKey" json:"-"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	Sessions  []Session `gorm:"foreignKey:UserRef; constraint:OnUpdate:CASCADE, OnDelete:CASCADE;" json:"-"`
	CreatedAt int64     `gorm:"autoCreateTime" json:"-" `
	UpdatedAt int64     `gorm:"autoUpdateTime:milli" json:"-"`
	SessionID uuid.UUID `json:"-" gorm:"-"`
}
