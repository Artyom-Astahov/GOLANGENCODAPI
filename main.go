package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Person struct {
	ID        string `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Phone     int    `json:"phone"`
	Email     string `json:"email"`
}

var persons []Person

func getPersons(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(persons)
}

func getPerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range persons {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}

func createPerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var person Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	person.ID = strconv.Itoa(rand.Intn(1000000))
	persons = append(persons, person)
	json.NewEncoder(w).Encode(person)
}

func updatePerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range persons {
		if item.ID == params["id"] {
			persons = append(persons[:index], persons[index+1:]...)
			var person Person
			_ = json.NewDecoder(r.Body).Decode(&person)
			person.ID = params["id"]
			persons = append(persons, person)
			json.NewEncoder(w).Encode(person)
			return
		}
	}
	json.NewEncoder(w).Encode(persons)
}

func deletePerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range persons {
		if item.ID == params["id"] {
			persons = append(persons[:index], persons[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(persons)
}

func main() {
	r := mux.NewRouter()
	persons = append(persons, Person{ID: "1", Firstname: "Артём", Lastname: "Астахов", Phone: 89156151929, Email: "astahov.artem.sh@gmail.com"})
	persons = append(persons, Person{ID: "2", Firstname: "Иван", Lastname: "Иванов", Phone: 89966011826, Email: "skatchanel@gmail.com"})
	r.HandleFunc("/person", getPersons).Methods("GET")
	r.HandleFunc("/person/{id}", getPerson).Methods("GET")
	r.HandleFunc("/persons", createPerson).Methods("POST")
	r.HandleFunc("/persons/{id}", updatePerson).Methods("PUT")
	r.HandleFunc("/persons/{id}", deletePerson).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", r))
}
