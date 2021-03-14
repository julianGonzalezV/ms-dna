package server

import (
	uRoute "ms-dna/pkg/dna/infrastructure/rest"
	"net/http"

	"github.com/gorilla/mux"
)

// Server
type Server interface {
	Router() http.Handler
}

type Api struct {
	router http.Handler
}

// New  inicializa el servidor, pasandole las rutas por las que escuchará
func New(dnaRoute uRoute.DnaRoute) Server {
	api := &Api{}
	r := mux.NewRouter()
	dnaRoute.AddRoutes(r)
	api.router = r
	return api
}

/// Router() se implementa para el struct Api, para que cumpla con la firma de la Interface Service
func (a *Api) Router() http.Handler {
	return a.router
}
