# Simple HTTP CRUD in Go
This example shows how to utilize the Go programming language to write an HTTP server.
It implements methods for create, read, update and delete on a map of people.
The data is stored in-memory and is therefore lost after stopping the server.

## Usage
Start the server (from src directory)
```bash
go run http_srv.go
```

The following commands show how to interface with the server using the popular cURL program.
Open another terminal instance to execute these commands.


**Create a new person**
```bash
curl -X POST http://localhost:8080/people/create \
    -H "Content-Type: application/x-www-form-urlencoded" \
    -d 'first_name=John&last_name=Doe'

### Output: created
```
When this command is executed, a new person gets stored in the `people` map.
The attributes fo the person, namely firstname and lastname, are specified in the POST body
URL-encoded. In this example `John Doe` is used.

The following snippet is responsible for creating new a ner person.
```go
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
```

**List all people**
```bash
curl http://localhost:8080/people

### Output: {"0":{"first_name":"John","last_name":"Doe"}}
```
This command calls the `readPeople()` function.
It returns all the people currently stored in the data map in a JSON format.

```go
func readPeople(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(people)
	if err != nil {
		log.Println(err)
	}
}
```

**Update a person**

When updating a person, an identifying key must be specified.
for this we just use an incrementing integer value starting at zero.
This key can be found in the output from the list function.
For the first person inserted, this is (of course) 0.

```bash
curl -X PUT http://localhost:8080/people/update?id=0 \
     -H "Content-Type: application/x-www-form-urlencoded" \
     -d 'first_name=Alice&last_name=Smith'

### Output: updated
```

```go
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
```

Now check again the list function

```bash
curl http://localhost:8080/people

### Output: {"0":{"first_name":"Alice","last_name":"Smith"}}
```

**Delete a person**

Deleting the only person in the map by calling the following HTTP endpoint with the ID 0.
```bash
curl -X DELETE http://localhost:8080/people/delete?id=0

### Output: deleted
```

```go
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
```


The `people` map should now be empty.
Check this by calling the list function again

```bash
curl http://localhost:8080/people

### Output: {}
```

<a href='https://github.com/sagar-jadhav/go-examples/blob/master/src/http_srv.go' target='_blank'>Source Code</a>

### Contributors
- <a href='https://github.com/davidkroell' target='_blank'>David Kröll</a>

[<< Home Page](./) | [Previous << Concurrency](./concurrency.html) | [Next >> HTTP GET](./http_get.html)