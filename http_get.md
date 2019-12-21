---
layout: default
---

# HTTP GET
Example demonstates How to call REST API which supports GET Method? & How to convert response JSON to Object?


```go
func main() {

	// mapping json keys to the struct fields. 
	// Make sure first character of each field value is uppercase, 
	// otherwise mapping will not work
	type User struct {
		Id        int    `json:"id"`
		Email     string `json:"email"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Avatar    string `json:"avatar"`
	}

	// mapping json keys to the struct fields. 
	// Make sure first character of each field value is uppercase, 
	// otherwise mapping will not work. 
	// We are taking this struct as json is nested object 
	// where "data" is the key for inner object.
	type Data struct {
		User User `json:"data"`
	}

	var data Data

	/*response from dummy API: {
		"data":{
			"id":7,
			"email":"michael.lawson@reqres.in",
			"first_name":"Michael",
			"last_name":"Lawson",
			"avatar":"https://s3.amazonaws.com/uifaces/faces/twitter/follettkyle/128.jpg"
		}
	}
	*/
	
	response, err := http.Get("http://reqres.in/api/users/7")
	if err != nil {
		fmt.Printf("Request Failed %s\n", err) // print error if err is not empty
	} else {
	    // json.Decode can also be used to do this conversion. 
		// Declared blank identifier("_") to avoid unused variables 
		// from the returns values.
		jsonData, _ := ioutil.ReadAll(response.Body)

		// print response as json string	
		fmt.Printf("%s\n", jsonData)

		// there are many ways to do this	
		json.Unmarshal(jsonData, &data)

		// print struct values with field name	
		fmt.Printf("%+v", data.User)                 
	}
}
```

[Source code](https://github.com/sagar-jadhav/go-examples/blob/master/src/http_get.go)

### Output

```bash
{
  "data": {
    "id": 7,
    "email": "michael.lawson@reqres.in",
    "first_name": "Michael",
    "last_name": "Lawson",
    "avatar": "https://s3.amazonaws.com/uifaces/faces/twitter/follettkyle/128.jpg"
  }
}
{
	Id:7 Email:michael.lawson@reqres.in 
	FirstName:Michael 
	LastName:Lawson 
	Avatar:https://s3.amazonaws.com/uifaces/faces/twitter/follettkyle/128.jpg
}
```

[<< Back](./)
