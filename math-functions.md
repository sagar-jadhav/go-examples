---
layout: default
---

# Math Functions

Example demonstrates following Math Functions available in Go:
- Power: a^b

```go
func main() {
	fmt.Println("2^4 =", pow(2, 4))
}

func pow(base, exp float64) float64 {
	return math.Pow(base, exp)
}
```


<a href='https://play.golang.org/p/ZyhuQAlcnsX' target='_blank'>Try It Out</a>

[Source Code](https://github.com/sagar-jadhav/go-examples/blob/master/src/test/power.go)

### Output

```bash
2^4 = 16
```

[<< Back](./)
