package controller

import (
	model "Guruprasad/Model"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetAllMovies(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/x-www-form-urlencode")

	allMovies := getAll()

	json.NewEncoder(w).Encode(allMovies)

}

func CreateMovie(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie model.Netflix

	_ = json.NewDecoder(r.Body).Decode(&movie)

	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)

}

func MarkWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)

	updateMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])

}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)

	deleteRecord(params["id"])
	json.NewEncoder(w).Encode(params["id"])

}

func DeleteAllMovie(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	deleteAll()
	json.NewEncoder(w).Encode("Deleted the records")

}
