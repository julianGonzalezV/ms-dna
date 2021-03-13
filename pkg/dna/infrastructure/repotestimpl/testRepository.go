package repotestimpl

import (
	"ms-dna/pkg/dna/domain/entity"
	"ms-dna/pkg/dna/domain/repository"
)

type testRepository struct {
}

// New  create a Mongo repo
func New() repository.DnaRepository {
	return testRepository{}
}

// SaveDna saves a dna record
func (repo testRepository) SaveDna(entity *entity.Dna) error {
	return nil
}

// Fetch return all records saved in storage
func (repo testRepository) GetDnas() ([]*entity.Dna, error) {
	return nil, nil
}
