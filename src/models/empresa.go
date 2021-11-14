package models

import (
	"errors"
)

type Empresa struct {
	ID       uint32 `json:"id"`
	Empresa  string `json:"empresa"`
	WhatsApp string `json:"whatsapp"`
}

//Validators Valida a estrutura Brands
func (e *Empresa) Validators() error {
	if e.Empresa == "" {
		return errors.New("empresa é requerido!")
	}
	if len(e.Empresa) > 100 {
		return errors.New("empresa não pode ter mais de 100 caracteres!")
	}
	if len(e.WhatsApp) > 100 {
		return errors.New("WhatsApp não pode ter mais de 15 caracteres!")
	}

	if e.WhatsApp == "" {
		return errors.New("WhatsApp é requerido!")
	}
	return nil
}
