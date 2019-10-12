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
The attributes fo the person, namely first- and lastname, are specified in the POST body
URL-encoded. In this example `John Doe` is used.

**List all people**
```bash
curl http://localhost:8080/people

### Output: {"0":{"first_name":"John","last_name":"Doe"}}
```
This command calls the `readPeople()` function.
It returns all the people currently stored in the data map in a JSON format.


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

The `people` map should now be empty.
Check this by calling the list function again

```bash
curl http://localhost:8080/people

### Output: {}
```
