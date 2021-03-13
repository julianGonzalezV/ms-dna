package mapper

import (
	"ms-dna/pkg/dna/domain/entity"
	"ms-dna/pkg/dna/infrastructure/request"
)

func DnaRequestToDnaDomain(request request.DnaRequest) *entity.Dna {
	return &entity.Dna{
		IsMutant: 0,
		Sequence: request.Sequence,
	}
}
