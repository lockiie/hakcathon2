package models

import (
	"errors"
)

type Produto struct {
	ID        uint32    `json:"id"`
	Produto   string    `json:"produto"`
	Foto      string    `json:"foto"`
	Descricao string    `json:"descricao"`
	Valor     float32   `json:"valor"`
	Categoria Categoria `json:"categoria"`
	Empresa   Empresa   `json:"empresa"`
}

//Validators Valida a estrutura Brands
func (p *Produto) Validators() error {
	if p.Produto == "" {
		return errors.New("produto é requerido!")
	}
	if len(p.Produto) > 100 {
		return errors.New("produto não pode ter mais de 100 caracteres!")
	}

	// if p.Foto == "" {
	// 	return errors.New("foto é requerido!")
	// }
	// if len(p.Foto) > 45 {
	// 	return errors.New("foto não pode ter mais de 45 caracteres!")
	// }

	if p.Valor == 0 {
		return errors.New("Valor é requerido")
	}

	if p.Categoria.ID == 0 {
		return errors.New("categoria é requerida")
	}
	if p.Empresa.ID == 0 {
		return errors.New("empresa é requerida")
	}
	return nil
}
