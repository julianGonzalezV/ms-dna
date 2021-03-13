package entity

// New function is used to create a new struct
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
