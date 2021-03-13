package server

import (
	uRoute "ms-dna/pkg/dna/infrastructure/rest"
	"net/http"

	"github.com/gorilla/mux"
)

type Api struct {
	router http.Handler
}

// Server ...
type Server interface {
	Router() http.Handler
}

// New ...
func New(dnaRoute uRoute.DnaRoute) Server {
	api := &Api{}
	r := mux.NewRouter()
	dnaRoute.AddRoutes(r)
	api.router = r
	return api
}

func (a *Api) Router() http.Handler {
	return a.router
}
