package main

import (
	"pokemon-api/database"
	"testing"
)

func TestExists1(t *testing.T){
	//go handleRequests() //run REST asynchronous
	pok := database.Pokemon{ID: "1", Name: "Pikachu", Type: "Electric"}
	if exists(pok){
		t.Logf("exists method passed test with succes")
	}else{
		t.Errorf("exists method failed, expected %v, got %v", true, exists(pok))
	}
}


func TestExists2(t *testing.T){
	//go handleRequests() //run REST asynchronous
	pok := database.Pokemon{ID: "3", Name: "Pikachu", Type: "Electric"}
	if !exists(pok){
		t.Logf("exists method passed test successfully")
	}else{
		t.Errorf("exists method failed, expected %v, got %v", false, exists(pok))
	}
}

func TestAddPokemon1(t *testing.T){
	pok := database.Pokemon{ID : "3", Name : "Venusaur", Type: "Grass"}
	newPokemon(pok)
	if exists(pok){
		t.Logf("newPokemon method passed test successfully")
	}else{
		t.Errorf("newPokemon method failed, no pokemon was add to DB")
	}
}
func TestAddPokemon2(t *testing.T){
	pok := database.Pokemon{ID : "3", Name : "", Type: ""}
	if !newPokemon(pok){
		t.Logf("newPokemon method passed test successfully")
	}else{
		t.Errorf("newPokemon method failed, expected %v, got %v", false, newPokemon(pok))
	}
}