---
layout: default
---

# Math Functions

Example demonstrates following Math Functions available in Go:
- Power: a^b
Click [here](https://golang.org/pkg/math/) to learn more

```go
func main() {
	fmt.Println("2^4 =", pow(2, 4))
}

func pow(base, exp float64) float64 {
	return math.Pow(base, exp)
}
```

### Output

```bash
2^4 = 16
```

<a href='https://play.golang.org/p/ZyhuQAlcnsX' target='_blank'>Try It Out</a> | <a href='https://github.com/sagar-jadhav/go-examples/blob/master/src/test/power.go' target='_blank'>Source Code</a>

### Contributors
- <a href='https://github.com/rajkumarGosavi' target='_blank'>Rajkumar</a>
- <a href='https://github.com/sagar-jadhav' target='_blank'>Sagar Jadhav</a>

[<< Home Page](./) | [Previous << Comparison Operators](./comparison-operators.html) | [Next >> Format Verbs](./format-verbs.html)
