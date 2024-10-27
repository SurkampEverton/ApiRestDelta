package models

import "github.com/google/uuid"

type Usuarios struct {
	Usid uuid.UUID `gorm:"primaryKey" json:"id"`
}
