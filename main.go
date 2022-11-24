package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// crear una ruta
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/contacto", Contact)
	// nil variable vacia
	fmt.Println("el servidor esta arriba")
	server := http.ListenAndServe(":8080", router)
	log.Fatal(server)
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hola mundo desde mux")
}

func Contact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pagina de contacto")
}
