package service

import (
	"ms-dna/pkg/dna/domain/entity"
	"ms-dna/pkg/dna/domain/repository"
	"ms-dna/shared/customerror"
)

// DnaServiceInterface estbalece las reglas y operaciones del negocio
type DnaServiceInterface interface {
	ValidateDna(entity *entity.Dna) error
	GetStats() (*entity.Ratio, error)
}

type dnaService struct {
	dnaRepository repository.DnaRepository
}

// New creates a domain service for core logic
func New(repository repository.DnaRepository) DnaServiceInterface {
	return &dnaService{repository}
}

// ValidateDna valida el dna y llama a guardar en la base de datos
func (service *dnaService) ValidateDna(entity *entity.Dna) error {
	/// Se valida que la entrada sea correcta
	if sequenceValidation(entity.Sequence) {
		// Se valida si es mutante
		isM := isAMutantDna(entity.Sequence)
		if isM {
			entity.IsMutant = 1
			return service.dnaRepository.SaveDna(entity)
		} else {
			service.dnaRepository.SaveDna(entity)
			return customerror.ErrNoMutantDna
		}

	} else {
		/// Si no es correcta no se incia el computo sino que se devuelve el error que es un BAD REQUEST
		return customerror.ErrDNASeqValidation
	}

}

// ChangePassword searches a record
func (service *dnaService) GetStats() (*entity.Ratio, error) {
	dnas, error := service.dnaRepository.GetDnas()
	if error != nil {
		return nil, error
	} else {

	}
	return getRatio(dnas), nil
}

/// Se obtienen el ratio para retornar en el consumo de esatdisticas
func getRatio(dnas []*entity.Dna) *entity.Ratio {
	mutantCount := 0
	humanCount := 0
	rat := float64(0.0)
	for _, dna := range dnas {
		if dna.IsMutant > 0 {
			mutantCount++
		} else {
			humanCount++
		}
	}
	if humanCount > 0 {
		rat = float64(float64(mutantCount) / float64(humanCount))
	}

	return &entity.Ratio{
		CountMutantDna: mutantCount,
		CountHumanDna:  humanCount,
		Ratio:          rat,
	}
}
