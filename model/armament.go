package model

import "gorm.io/gorm"

type Armament struct {
	gorm.Model
	SpaceshipID uint   `json:"-"`
	Title       string `json:"title"`
	Qty         string `json:"qty"`
}
