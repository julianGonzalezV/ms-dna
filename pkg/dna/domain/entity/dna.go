package entity

// New devuelve la entidad de domina DNA
func New(sequence []string) *Dna {
	return &Dna{
		Sequence: sequence,
	}
}

type Dna struct {
	IsMutant int
	Sequence []string
}

type Ratio struct {
	CountMutantDna int
	CountHumanDna  int
	Ratio          float64
}
