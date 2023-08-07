package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Pattern     string
	Method      string
	HandlerFunc http.HandlerFunc
}

func (s *Server) NewRouter() *mux.Router {
	router := mux.NewRouter()
	return router
}
