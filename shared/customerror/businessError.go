package customerror

import "errors"

var (
	ErrMongo            = errors.New("Error en operacion mongo")
	ErrDNASeqValidation = errors.New("Matriz incorrecta, revise que sea NxN, tamañi mínimo correcto y caracteres correctos")
	ErrNoMutantDna      = errors.New("No es un mutante")
)
