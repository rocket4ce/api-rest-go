package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
)

var collection = GetSession().DB("course-go").C("movies")

var movies = Movies{
	Movie{"Sin limite 2013", 2013, "123"},
	Movie{"Sin limite 2013", 2013, "123"},
	Movie{"Sin limite 2013", 2013, "123"},
	Movie{"Sin limite 2013", 2013, "123"},
	Movie{"Sin limite 2013", 2013, "123"},
}

func GetSession() *mgo.Session {
	session, err := mgo.Dial("mongodb://root:example@mongo:27017")
	if err != nil {
		panic(err)
	}
	return session
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
	// log.Println(movie_data)
	err = collection.Insert(movie_data)
	if err != nil {
		w.WriteHeader(500)
	}
	json.NewEncoder(w).Encode(movie_data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	// movies = append(movies, movie_data)

}
