package mapper

import (
	"ms-dna/pkg/dna/domain/entity"
	"ms-dna/pkg/dna/infrastructure/response"
)

func RatioDomainToStatResponse(ratio *entity.Ratio) *response.StatResponse {
	return &response.StatResponse{
		CountMutantDna: ratio.CountMutantDna,
		CountHumanDna:  ratio.CountHumanDna,
		Ratio:          ratio.Ratio,
	}
}
