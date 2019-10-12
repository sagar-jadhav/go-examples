package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

const bindPort = ":8080"

var people = map[int]Person{}
var lastId = 0

func updatePerson(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("bad request"))
		return
	}

	person, ok := people[intId]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("not found"))
		return
	}

	person.FirstName = r.FormValue("first_name")
	person.LastName = r.FormValue("last_name")

	people[intId] = person
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("updated"))
}

func createPerson(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		fn := r.FormValue("first_name")
		ln := r.FormValue("last_name")

		p := Person{
			FirstName: fn,
			LastName:  ln,
		}

		people[lastId] = p
		lastId++

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("created"))

	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("bad request"))
	}
}

func readPeople(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(people)
	if err != nil {
		log.Println(err)
	}
}

func deletePerson(w http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" {
		id := r.FormValue("id")

		idInt, err := strconv.Atoi(id)
		if err != nil {
			log.Println(err)
		}

		_, ok := people[idInt]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		delete(people, idInt)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("deleted"))
	} else {
		w.Write([]byte("bad request"))
	}
}

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

// http_srv program demonstrates how to serve simple content over http
// this sample implements three different endpoints
// /hello-world
// /hello-json
// /hello-user?name=<USERNAME>
// call them with curl: curl localhost:8080/hello-world
func main() {
	people = make(map[int]Person)

	// updatePerson people
	http.HandleFunc("/people", readPeople)
	http.HandleFunc("/people/create", createPerson)
	http.HandleFunc("/people/update", updatePerson)
	http.HandleFunc("/people/delete", deletePerson)

	log.Println("listening on ", bindPort)
	log.Fatal(http.ListenAndServe(bindPort, nil))
}
