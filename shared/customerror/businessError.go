package customerror

import "errors"

var (
	ErrMongo         = errors.New("Error en operacion mongo")
	DNASeqValidation = errors.New("Matriz incorrecta, revise que sea NxN, tamañi mínimo correcto y caracteres correctos")
)
