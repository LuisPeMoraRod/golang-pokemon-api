package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"pokemon-api/database"
	"testing"
)

const ok ="200 OK"

func TestGetAllPokemon(t *testing.T){
	go handleRequests()
	resp, err := http.Get("http://localhost:8080/pokemon")
	if err  != nil{
		t.Error(err)
	}
	defer resp.Body.Close()

	if resp.Status == ok{
		t.Log("Request made successfully")
	}else{
		t.Errorf("Connection with REST service failed, status expect %v, got %v,", ok, resp.Status)
	}
}

func TestNewPokemon(t *testing.T){
	requestBody := []byte(`{"ID":"4","Name":"Ditto","Type":"Normal"}`)
	req, err := http.NewRequest("POST", "http://localhost:8080/pokemon/new", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	if resp.Status == ok {
		t.Log("Pokemon added successfully to DB")
	}else{
		t.Errorf("addPokemon method failed, expected %v, got %v", ok , resp.Status)
	}

}

func TestGetPokemon(t *testing.T) {
	resp, err := http.Get("http://localhost:8080/pokemon/4")
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(nil)
	}
	var pokemon database.Pokemon
	json.Unmarshal(body, &pokemon)

	if (pokemon.ID == "4") && (pokemon.Name == "Ditto") && (pokemon.Type == "Normal") {
		t.Log("getPokemon method passed test successfully")
	} else {
		t.Errorf("getPokemon failed, expected %v, got %v", "Venusaur", pokemon.Name)
	}
}
