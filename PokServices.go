package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"pokemon-api/database"
)

func getAllPokemon(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(database.PokemonDb)
	w.WriteHeader(http.StatusOK)
}

func getPokemon(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	key := vars["id"]

	if !idExists(key){
		w.WriteHeader(http.StatusNotFound)
		return
	}

	for _, pokemon := range database.PokemonDb {
		if pokemon.ID == key {
			json.NewEncoder(w).Encode(pokemon)
		}
	}
	w.WriteHeader(http.StatusOK)
}

func addPokemon(w http.ResponseWriter, r *http.Request){
	reqBody, _ := ioutil.ReadAll(r.Body)

	var pokemon database.Pokemon
	json.Unmarshal(reqBody, &pokemon)

	if exists(pokemon){
		w.WriteHeader(http.StatusNotModified)
		return
	}
	if newPokemon(pokemon){
		w.WriteHeader(http.StatusOK)
		return
	}
	w.WriteHeader(http.StatusNotModified)
}

/*!
Checks if passed pokemon exists in DB
 */
func exists(pokemon database.Pokemon) bool {
	if idExists(pokemon.ID){
		return true
	}
	if _, ok := database.PokemonDb[pokemon.Name]; ok{
		if _, ok := database.PokemonDb[pokemon.Type];ok{
			return true
		}
	}
	return false
}

/*!
Checks if a pokemon is registered with passed id
 */
func idExists(id string) bool{
	if id == "" {
		return true
	}
	for _, article := range database.PokemonDb {
		if article.ID == id {
			return true
		}
	}
	return false
}

/*!
Add new pokemon to DB
 */
func newPokemon(pokemon database.Pokemon) bool{
	if (pokemon.Name != "") && (pokemon.Type != ""){
		database.PokemonDb[pokemon.ID] = pokemon
		return true
	}
	return false
}