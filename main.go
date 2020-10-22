package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func handleRequests() {
	port := os.Getenv("PORT")
	if port == ""{
		port = "8080"
	}
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.Use(commonMiddleware)
	myRouter.HandleFunc("/pokemon", getAllPokemon).Methods("GET")
	myRouter.HandleFunc("/pokemon/new", addPokemon).Methods("POST")
	myRouter.HandleFunc("/pokemon/{id}", getPokemon).Methods("GET")
	log.Fatal(http.ListenAndServe(":"+port, myRouter))
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
func main() {
	fmt.Println("Pokemon Rest API")
	handleRequests()
}
