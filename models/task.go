package models

import (
	"time"
)

type Task struct {
	ID                uint      `gorm:"primaryKey"`
	Titre             string    `gorm:"size:255;not null"`
	Description       string    `gorm:"type:text"`
	Status            string    `gorm:"size:1;not null;default:'A'"`

	ProjectID         uint      `gorm:"not null"`
	Project           Project   `gorm:"foreignKey:ProjectID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	DateAjout         time.Time `gorm:"autoCreateTime"`
	DateModification  time.Time `gorm:"autoUpdateTime"`
	DateDebut         *time.Time
	DateFin           *time.Time
	DateLimite          *time.Time

	InstigateurID     *uint // Juste l'ID de l'initiateur
	ParticipantsIDs   []uint `gorm:"-" json:"participants_ids"` // Liste d'IDs d'utilisateurs participants, NON stockée en DB

	ImagePath         *string
	FilePath          *string
}
