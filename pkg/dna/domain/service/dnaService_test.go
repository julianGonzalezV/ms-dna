package service

import (
	"ms-dna/pkg/dna/domain/entity"
	"ms-dna/pkg/dna/infrastructure/repotestimpl"
	"ms-dna/shared/customerror"
	"testing"
)

var (
	dnaMutant = []string{
		"ATGCGA",
		"CAGTGC",
		"TTATGT",
		"AGAAGG",
		"CCCCTA",
		"TCACTG"}

	dnaNOMutant = []string{
		"ATGCGA",
		"CAGTGC",
		"TTATTT",
		"AGACGG",
		"GCGTCA",
		"TCACTG"}

	dnaMutantCols = []string{
		"ATGCGA",
		"CAGTGA",
		"TTGTTA",
		"AGTCGA",
		"GCGTCA",
		"TCACGG"}

	dnaMutantDiag = []string{
		"ACTCGA",
		"CATTGA",
		"TTATTG",
		"AGTAGA",
		"GCCGCA",
		"TCGCGG"}
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

///TestValidateDna valida un AND mutante y otro NO mutante, note que para el No mutante se espera
/// un error tipo customerror.ErrNoMutantDna
func TestValidateDna(t *testing.T) {
	serviceI := New(repotestimpl.New())
	error := serviceI.ValidateDna(&entity.Dna{IsMutant: 1, Sequence: dnaMutant})
	if error != nil {
		t.Errorf(error.Error())
	}

	error = serviceI.ValidateDna(&entity.Dna{IsMutant: 1, Sequence: dnaNOMutant})
	if error != customerror.ErrNoMutantDna {
		t.Errorf(error.Error())
	}

	error = serviceI.ValidateDna(&entity.Dna{IsMutant: 1, Sequence: dnaMutantCols})
	if error != customerror.ErrNoMutantDna {
		t.Errorf(error.Error())
	}

	error = serviceI.ValidateDna(&entity.Dna{IsMutant: 1, Sequence: dnaMutantDiag})
	if error != nil {
		t.Errorf(error.Error())
	}
}

/// TestGetStats se valida que Stats funcione correctamente
func TestGetStats(t *testing.T) {
	serviceI := New(repotestimpl.New())
	_, error := serviceI.GetStats()
	if error != nil {
		t.Errorf("Falla test de getStats")
	}
}
