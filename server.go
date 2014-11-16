package main

import (
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

const CONTENT_DIR = "content"

// main starts the web server
func main() {
	// negroni + mux
	n := negroni.New()

	// setup middleware
	n.Use(negroni.NewRecovery())
	n.Use(negroni.NewLogger())
	n.Use(negroni.NewStatic(http.Dir("./client/dist/")))
	n.Use(negroni.HandlerFunc(respondWithJson))

	// listen to handlers
	n.UseHandler(Handlers())

	// start server
	n.Run(":8080")
}

// Handlers connects the handlers to a new router
func Handlers() *mux.Router {
	r := mux.NewRouter()
	api := r.PathPrefix("/api").Subrouter()

	// directories
	api.HandleFunc("/dir/{path:.*}", handleReadDir).Methods("GET")
	api.HandleFunc("/dir/{path:.*}", handleCreateDir).Methods("POST")
	api.HandleFunc("/dir/{path:.*}", handleUpdateDir).Methods("PUT")
	api.HandleFunc("/dir/{path:.*}", handleDeleteDir).Methods("DELETE")

	// pages
	api.HandleFunc("/page/{path:.*}", handleReadPage).Methods("GET")
	api.HandleFunc("/page/{path:.*}", handleCreatePage).Methods("POST")
	api.HandleFunc("/page/{path:.*}", handleUpdatePage).Methods("PUT")
	api.HandleFunc("/page/{path:.*}", handleDeletePage).Methods("DELETE")

	// config
	api.HandleFunc("/config", handleReadConfig).Methods("GET")
	api.HandleFunc("/config", handleUpdateConfig).Methods("PUT")

	// files
	api.HandleFunc("/copy/{path:.*}", handleCopy).Methods("POST")

	return r
}

// respondWithJson sets the content-type header to application/json
func respondWithJson(w http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	w.Header().Set("Content-Type", "application/json")
	next(w, req)
}
