---
layout: default
---

# Switch Statement

Example demonstrates working of Switch Statement 
using the sign to word program.
Click [here](https://tour.golang.org/flowcontrol/9) to learn more

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

### Output

```bash
Equal To
```

<a href='https://play.golang.com/p/2FMbm1koSBi' target='_blank'>Try It Out</a> | <a href='https://github.com/sagar-jadhav/go-examples/blob/master/src/switch.go' target='_blank'>Source Code</a>

### Contributors
- <a href='https://github.com/sagar-jadhav' target='_blank'>Sagar Jadhav</a>

[<< Home Page](./) | [Previous << If Else Statement](./if-elseif-else.html) | [Next >> For Loop](./for-loop.html)
