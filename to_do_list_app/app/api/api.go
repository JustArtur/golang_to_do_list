package api

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"to_do_list_app/app/controllers"
)

func RunServer() {
	server := http.Server{
		Handler: newRoute(),
		Addr:    ":8000",
	}

	log.Printf("Starting up server on port %s", server.Addr)
	log.Fatal(server.ListenAndServe())
}

func newRoute() *mux.Router {
	router := mux.NewRouter()

	apiRouter := router.PathPrefix("/api/v1").Subrouter()
	apiRouter.HandleFunc("/tasks", controllers.Create).Methods("POST")
	apiRouter.HandleFunc("/tasks/:id", controllers.Show).Methods("GET")
	apiRouter.HandleFunc("/tasks/:id", controllers.Update).Methods("PUT")
	apiRouter.HandleFunc("/tasks/:id", controllers.Destroy).Methods("DELETE")
	apiRouter.HandleFunc("/tasks", controllers.Index).Methods("GET")

	return router
}
