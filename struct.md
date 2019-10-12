---
layout: default
---

# STRUCT

The example demonstrates structs. Structs are used to group one of more fields of logical types together.

```
//For example we could represent a Point struct as:
type Point struct {
	x int
	y int
}

func main() {

	//Then in order to use it we need to initialize it. There are 3 ways to initialize a struct in Go.
	var c Point
	c := Point{x: 1, y: 2}

	//Once you have a struct instance you can access its fields using the dot . operator:

	fmt.Println(c.x) // 1
	c.x = 10         //updating value
	fmt.Println(c.x) // 10

}

```

[Source Code](https://github.com/sagar-jadhav/go-examples/blob/master/src/struct.go)

### Output

```
1
10
```
[Back](./)
