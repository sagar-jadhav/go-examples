package main

import (
	"log"
	"net/http"
)

const bindPort = ":8080"

func helloJSON(w http.ResponseWriter, r *http.Request) {
	// set http content type header for the clients
	w.Header().Set("Content-Type", "application/json")

	// set http statuscode
	w.WriteHeader(http.StatusOK)

	// write json string in bytes to http output stream
	_, err := w.Write([]byte("{\"msg\": \"hello world\"}"))
	if err != nil {
		log.Println(err)
	}
}

func helloUser(w http.ResponseWriter, r *http.Request) {
	name, ok := r.URL.Query()["name"]
	if !ok {
		_, _ = w.Write([]byte("no name supplied"))
		return
	}

	if len(name) != 1 {
		_, _ = w.Write([]byte("too many arguments"))
		return
	}

	_, _ = w.Write([]byte("hello " + name[0]))

	log.Println("user ", name[0], " called this function")
}

// http_srv program demonstrates how to serve simple content over http
// this sample implements three different endpoints
// /hello-world
// /hello-json
// /hello-user?name=<USERNAME>
// call them with curl: curl localhost:8080/hello-world
func main() {

	// implement simple inline function
	http.HandleFunc("/hello-world", func(w http.ResponseWriter, r *http.Request) {

		// write to output stream
		_, err := w.Write([]byte("hello world"))
		if err != nil {
			log.Println(err)
		}
	})

	// bind http endpoints
	http.HandleFunc("/hello-json", helloJSON)
	http.HandleFunc("/hello-user", helloUser)

	log.Println("listening on ", bindPort)
	log.Fatal(http.ListenAndServe(bindPort, nil))
}
