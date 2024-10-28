package models

import (
	"github.com/google/uuid"
	"gopkg.in/validator.v2"
)

type Usuarios struct {
	Usid        uuid.UUID `gorm:"primaryKey" json:"id" validate:"required"`
	Usnome      string    `json:"nome" validate:"required"`
	Usid_pessoa uuid.UUID `json:"id_pessoa"`
	Email       string    `json:"pass" validate:"required"`

	//Ussenha     string    `json:"senha" validate:"required"`
}

func ValidarDadosUsuario(usuario *Usuarios) error {
	if err := validator.Validate(usuario); err != nil {
		return err
	}

	return nil
}
