---
layout: default
---

# Map

Example demonstrates usage of map in go.

```go
func main() {
	/* Uncomment this code block and verify that program will give
	   * runtime error as Map is not initialized
		   var alphabets map[string]string
		   alphabets["A"] = "Apple"
		   fmt.Println(alphabets)
	*/

	//Initializing Map

	//variation #1
	var alphabets map[string]string
	alphabets = make(map[string]string)
	alphabets["A"] = "Apple"
	alphabets["B"] = "Ball"
	alphabets["C"] = "Cat"
	fmt.Println(alphabets)

	//variation #2
	players := make(map[int]string)
	players[1] = "Sachin"
	players[2] = "Virat"
	players[3] = "Dhoni"
	fmt.Println(players)

	//variation #3
	signs := map[string]string{
		"=": "equalto",
		"+": "plus",
	}
	fmt.Println(signs)

	//If provided key is not present then map will return zero value
	fmt.Println(signs["-"])

	//There is another way which is provided by Go which returns two
	//values one is value & other is whether operation is successful
	//or not
	value, isSuccessful := signs["-"]
	fmt.Println(value, isSuccessful)

	if value, isSuccessful := signs["+"]; isSuccessful {
		fmt.Println(value)
	}

	//Iterating Over Map
	//Variation #1
	for key, value := range alphabets {
		fmt.Println(key, value)
	}

	//Variation #2
	for key := range players {
		fmt.Println(players[key])
	}
}
```
<a href='https://play.golang.org/p/HYTVd0XK6ms' target='_blank'>Try It Out</a>

[Source Code](https://github.com/sagar-jadhav/go-examples/blob/master/src/map.go)

# Output

```bash
map[B:Ball C:Cat A:Apple]
map[1:Sachin 2:Virat 3:Dhoni]
map[=:equalto +:plus]

false
plus
A Apple
B Ball
C Cat
Sachin
Virat
Dhoni
```

[<< Back](./)
