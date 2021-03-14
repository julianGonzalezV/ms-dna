package repository

import (
	"ms-dna/pkg/dna/domain/entity"
)

// DnaRepository interface establece la firma del repo que gestionará el almacenamiento de la data
type DnaRepository interface {
	// SaveDna almacena registros Dna evaluados
	SaveDna(entity *entity.Dna) error
	// GetDnas retorna todos los dnas almacenados
	GetDnas() ([]*entity.Dna, error)
}
