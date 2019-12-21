---
layout: default
---

# Switch Statement

Example demonstrates working of Switch Statement 
using the sign to word program

```go
func main() {
	var sign string = "="
	switch sign {
	case "=":
		fmt.Println("Equal To")
	case "<":
		fmt.Println("Less Than")
	case ">":
		fmt.Println("Greater Than")
	default:
		fmt.Println("Unknown")
	}
}
```


<a href='https://play.golang.com/p/2FMbm1koSBi' target='_blank'>Try It Out</a>

[Source Code](https://github.com/sagar-jadhav/go-examples/blob/master/src/switch.go)

### Output

```bash
Equal To
```

[<< Back](./)
