package rest

import (
	"encoding/json"
	"log"
	"ms-dna/pkg/dna/application"
	"ms-dna/pkg/dna/infrastructure/request"
	"ms-dna/shared/customerror"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	dnaApp application.DnaUseCaseInterface
)

/// se configura el ruteo de la aplicación
type DnaRoute interface {
	AddRoutes(router *mux.Router)
}

type dnaRoute struct {
	app application.DnaUseCaseInterface
}

// New ...
func New(
	app application.DnaUseCaseInterface,
) DnaRoute {
	dnaApp = app
	return &dnaRoute{app: dnaApp}
}

func (pRoute *dnaRoute) AddRoutes(router *mux.Router) {
	router.HandleFunc("/mutant", mutant).Methods(http.MethodPost)
	router.HandleFunc("/stats", stats).Methods(http.MethodGet)

}

func mutant(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var requestPayload request.DnaRequest
	err := decoder.Decode(&requestPayload)

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode("body erroneo")
		return
	}

	if err := dnaApp.ValidateMutant(requestPayload); err != nil {
		log.Println("error ValidateMutant", err)
		if err == customerror.ErrNoMutantDna {
			w.WriteHeader(http.StatusForbidden)
		}
		if err == customerror.ErrDNASeqValidation {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		_ = json.NewEncoder(w).Encode(err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func stats(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if result, error := dnaApp.Stats(); error != nil {
		log.Println(error)
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode("Error calculando las estadisticas")
		return
	} else {
		_ = json.NewEncoder(w).Encode(result)
	}

}
