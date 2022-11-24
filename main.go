package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// crear una ruta
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/movies", MovieList)
	router.HandleFunc("/movies/{id}", GetMovie)
	// nil variable vacia
	fmt.Println("el servidor esta arriba")
	server := http.ListenAndServe(":8080", router)
	log.Fatal(server)
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hola mundo desde mux")
}

func MovieList(w http.ResponseWriter, r *http.Request) {
	movies := Movies{
		Movie{"Sin limite 2013", 2013, "123"},
		Movie{"Sin limite 2013", 2013, "123"},
		Movie{"Sin limite 2013", 2013, "123"},
		Movie{"Sin limite 2013", 2013, "123"},
		Movie{"Sin limite 2013", 2013, "123"},
	}
	json.NewEncoder(w).Encode(movies)
}
func GetMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movie_id := params["id"]
	fmt.Fprintf(w, "Movie id %s", movie_id)
}
