package models

import "gopkg.in/validator.v2"

type Login struct {
	User string `json:"user" validate:"required"`
	Pass string `json:"pass" validate:"required"`
}

func ValidarDadosLogin(login *Login) error {
	if err := validator.Validate(login); err != nil {
		return err
	}

	return nil
}
