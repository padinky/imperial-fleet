package model

import "gorm.io/gorm"

type Spaceship struct {
	gorm.Model
	Name      string     `json:"name"`
	Class     string     `json:"class"`
	Crew      int16      `json:"crew"`
	Image     string     `json:"image"`
	Status    string     `json:"status"`
	Armaments []Armament `gorm:"foreignKey:SpaceshipID; constraint:OnUpdate:CASCADE, OnDelete:CASCADE;" json:"armaments"`
}
