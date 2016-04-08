package server

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"path/filepath"
)

func NewRouter() *mux.Router {

	apidir := "api"
	router := mux.NewRouter().StrictSlash(true)
	var r Routes
	err := filepath.Walk(apidir, r.findRoute)
	if err != nil {
		log.Fatal(err)
	}

	generateSwaggerFile()
	for _, route := range r {
		log.Printf("Adding route %v with method %v", route.Pattern, route.Method)
		var handler http.Handler

		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}
	router.
		Methods("GET").
		Path("/v0/services").
		Name("Service List").
		HandlerFunc(GetServices)

	log.Println("Adding route to /apidocs/")
	router.
		Methods("GET").
		PathPrefix("/apidocs").
		Name("Apidocs").
		Handler(http.StripPrefix("/apidocs", http.FileServer(http.Dir("./dist"))))
	return router

}
