// mongodb://root:example@mongo:27017/
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// crear una ruta
	router := NewRouter()
	// nil variable vacia
	fmt.Println("el servidor esta arriba")
	server := http.ListenAndServe(":8080", router)
	log.Fatal(server)
}
