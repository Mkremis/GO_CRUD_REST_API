package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Person struct {
	ID        string   `json:"id,omitempty"`
	FirstName string   `json:"firstname,omitempty"`
	LastName  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var people []Person

func GetPeopleEndpoint(res http.ResponseWriter, req *http.Request) {
	json.NewEncoder(res).Encode(people)
}
func GetPersonEndpoint(res http.ResponseWriter, req *http.Request)    {}
func CreatePersonEndpoint(res http.ResponseWriter, req *http.Request) {}
func DeletePersonEndpoint(res http.ResponseWriter, req *http.Request) {}

func main() {
	router := mux.NewRouter()

	people = append(people, Person{ID: "1", FirstName: "Martin", LastName: "Mengano", Address: &Address{
		City: "Doubling", State: "California",
	}})
	people = append(people, Person{ID: "2", FirstName: "Ignacio", LastName: "Sultano", Address: &Address{
		City: "Miami", State: "Florida",
	}})
	people = append(people, Person{ID: "2", FirstName: "Natalia", LastName: "Minovich"})

	//endpoints
	router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", GetPersonEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePersonEndpoint).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePersonEndpoint).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", router))
}
