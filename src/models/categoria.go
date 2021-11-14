package models

import (
	"errors"
)

type Categoria struct {
	ID        uint32 `json:"id"`
	Categoria string `json:"categoria"`
}

//Validators Valida a estrutura Brands
func (c *Categoria) Validators() error {
	if c.Categoria == "" {
		return errors.New("categoria é requerida!")
	}
	if len(c.Categoria) > 45 {
		return errors.New("categoria não pode ter mais de 45 caracteres!")
	}
	return nil
}
