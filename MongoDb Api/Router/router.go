package router

import (
	controller "Guruprasad/Controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/movies", controller.GetAllMovies).Methods("GET")
	router.HandleFunc("/api/create", controller.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movie/{id}", controller.MarkWatched).Methods("PUT")
	router.HandleFunc("/api/movie/delete/{id}", controller.DeleteMovie).Methods("DELETE")
	router.HandleFunc("/api/movie/delete", controller.DeleteAllMovie).Methods("DELETE")

	return router
}
