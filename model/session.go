package model

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	SessionID uuid.UUID `gorm:"primaryKey" json:"session_id" cookie:"session_id"`
	Expires   time.Time `json:"-"`
	UserRef   uint      `json:"-"`
	CreatedAt int64     `gorm:"autoCreateTime" json:"-" `
}
