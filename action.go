package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var movies = Movies{
	Movie{"Sin limite 2013", 2013, "123"},
	Movie{"Sin limite 2013", 2013, "123"},
	Movie{"Sin limite 2013", 2013, "123"},
	Movie{"Sin limite 2013", 2013, "123"},
	Movie{"Sin limite 2013", 2013, "123"},
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hola mundo desde mux")
}

func MovieList(w http.ResponseWriter, r *http.Request) {
	movies := movies
	json.NewEncoder(w).Encode(movies)
}
func GetMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movie_id := params["id"]
	fmt.Fprintf(w, "Movie id %s", movie_id)
}

func AddMovie(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var movie_data Movie
	err := decoder.Decode(&movie_data)
	if err != nil {
		panic(err)
	}

	defer r.Body.Close()
	log.Println(movie_data)
	movies = append(movies, movie_data)
}
