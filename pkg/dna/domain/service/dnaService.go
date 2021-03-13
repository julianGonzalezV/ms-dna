package service

// This component is in charge of handle domain rules
import (
	"ms-dna/pkg/dna/domain/entity"
	"ms-dna/pkg/dna/domain/repository"
	"ms-dna/shared/customerror"
)

// DnaServiceInterface interface that establishes functions to be implemented
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

// Create a new record
func (service *dnaService) ValidateDna(entity *entity.Dna) error {
	if sequenceValidation(entity.Sequence) {
		isM := isAMutantDna(entity.Sequence)
		if isM {
			entity.IsMutant = 1
		}
		return service.dnaRepository.SaveDna(entity)

	} else {
		return customerror.DNASeqValidation
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
