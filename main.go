package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"pokemon-api/database"
)

func getAllPokemons(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(database.PokemonDb)
}

func getPokemon(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	key := vars["id"]

	for _, article := range database.PokemonDb {
		if article.ID == key {
			json.NewEncoder(w).Encode(article)
		}
	}
}

func addPokemon(w http.ResponseWriter, r *http.Request){
	reqBody, _ := ioutil.ReadAll(r.Body)

	var pokemon database.Pokemon
	json.Unmarshal(reqBody, &pokemon)

	if exists(pokemon){
		w.WriteHeader(http.StatusNotModified)
		return
	}
	database.PokemonDb[pokemon.ID] = pokemon
	w.WriteHeader(http.StatusOK)
}

func exists(pokemon database.Pokemon) bool {
	if _, ok := database.PokemonDb[pokemon.ID]; ok{
		return true
	}
	return false
}

//func createPokemon()

func handleRequests() {
	port := os.Getenv("PORT")
	if port == ""{
		port = "80"
	}
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.Use(commonMiddleware)
	myRouter.HandleFunc("/pokemons", getAllPokemons).Methods("GET")
	myRouter.HandleFunc("/pokemons/new", addPokemon).Methods("POST")
	myRouter.HandleFunc("/pokemons/{id}", getPokemon).Methods("GET")

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
