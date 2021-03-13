package request

type DnaRequest struct {
	Sequence []string `json:"dna,omitempty"`
}
