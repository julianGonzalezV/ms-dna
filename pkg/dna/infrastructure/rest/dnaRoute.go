package rest

import (
	"encoding/json"
	"log"
	"ms-dna/pkg/dna/application"
	"ms-dna/pkg/dna/infrastructure/request"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	dnaApp application.DnaUseCaseInterface
)

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
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(err)
		return
	}
	w.WriteHeader(http.StatusCreated)
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
