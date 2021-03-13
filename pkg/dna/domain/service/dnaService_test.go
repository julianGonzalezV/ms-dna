package service

import (
	"ms-dna/pkg/dna/domain/entity"
	"ms-dna/pkg/dna/infrastructure/repotestimpl"
	"testing"
)

/// Validando todos los posibles valores de Ratio
func TestGetRatio(t *testing.T) {
	dnas := []*entity.Dna{
		{IsMutant: 1, Sequence: []string{}},
		{IsMutant: 0, Sequence: []string{}},
		{IsMutant: 0, Sequence: []string{}}}
	ratio := getRatio(dnas)
	if ratio.Ratio != 0.5 {
		t.Errorf("Ratio es  %f pero se espera %f \n", ratio.Ratio, 0.5)
	}
	if ratio.CountMutantDna != 1 {
		t.Errorf("CountMutantDna es  %d pero se espera %d \n", ratio.CountMutantDna, 1)
	}

	if ratio.CountHumanDna != 2 {
		t.Errorf("CountHumanDna es  %d pero se espera %d \n", ratio.CountHumanDna, 2)
	}
}

func TestSaveDna(t *testing.T) {
	serviceI := New(repotestimpl.New())
	result := serviceI.ValidateDna(&entity.Dna{IsMutant: 1, Sequence: []string{}})
	if result != nil {
		t.Errorf("Error registrando DNA")
	}
}

func TestGetStats(t *testing.T) {
	serviceI := New(repotestimpl.New())
	_, error := serviceI.GetStats()
	if error != nil {
		t.Errorf("Falla test de getStats")
	}
}
