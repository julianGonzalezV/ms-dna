package application

import (
	"ms-dna/pkg/dna/domain/service"
	"ms-dna/pkg/dna/infrastructure/mapper"
	"ms-dna/pkg/dna/infrastructure/request"
	"ms-dna/pkg/dna/infrastructure/response"
)

// DnaUseCaseInterface provee la firma de las operaciones
type DnaUseCaseInterface interface {
	ValidateMutant(requestData request.DnaRequest) error
	Stats() (*response.StatResponse, error)
}

type dnaUseCase struct {
	service service.DnaServiceInterface
}

// New crea la implementación de la capa deapplicación
func New(service service.DnaServiceInterface) DnaUseCaseInterface {
	return &dnaUseCase{service}
}

// ValidateMutant valida si un Dna  pertenece o no a un mutante
func (app *dnaUseCase) ValidateMutant(requestData request.DnaRequest) error {
	return app.service.ValidateDna(mapper.DnaRequestToDnaDomain(requestData))

}

// Stats devuelve las estadisticas de las verificaciones de DNAs
func (app *dnaUseCase) Stats() (*response.StatResponse, error) {
	ratio, error := app.service.GetStats()
	return mapper.RatioDomainToStatResponse(ratio), error
}
