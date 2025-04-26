package models

import (
	"time"
)

type Project struct {
	ID                uint      `gorm:"primaryKey"`
	Titre             string    `gorm:"size:255;not null"`
	Description       string    `gorm:"type:text"`
	DateAjout         time.Time `gorm:"autoCreateTime"`
	DateModification  time.Time `gorm:"autoUpdateTime"`
}
