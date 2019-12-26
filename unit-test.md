---
layout: default
---

# Unit Test

Test is very needed in todays development. In go, there is a powerful tool to do a unit Test. let's say we have a file like below.

```go
package calc

func Sum(a int, b int) int {
	return a + b
}
```

save the file as `sum.go`, and make your own test file as `sum_test.go`. you can easily make the test file to put `_test.go`. you should follow the naming convetion for mighty go.

```go
package calc

import "testing" // go basic test

func TestSum(t *testing.T) { // Test + Function Name
// arg would be basic t *testing.T
	r := Sum(1, 2) // we are adding 1+2.
	if r != 3 {    // when is not 3? it means it is failed!
		t.Error("WRONG. r=", r) // error
	}
}
```

Functions that perform unit tests must adhere to the following rules: Otherwise, the compiler will not recognize the test code.

Test functions always start with Test, like TestSum.
Test is followed by the function name, and the first letter of the function name should always be capitalized. E.g. TestSum, TestReadData, TestConvert
It always takes a parameter of type * testing.T.
Writing test code is simple. You only need to determine if it fits with the conclusions you need to put in the function and then come out. In this case, since the Sum function adds two numbers, it is normal to get 3, which is the sum of 1 and 2. If you don't get 3, the function is implemented incorrectly, so use the t.Error function to print an error.

In the console (terminal), go to the `GOPATH / src / calc` directory and run the following command (on Windows, run a command prompt or PowerShell): I set `GOPATH` to  this project.
make sure always you have `$GOPATH` settings!

```bash
~$ cd $GOPATH/src/calc
~/hello_project/src/calc$ go test
PASS
ok      calc    0.088s
```

with `go test`, you can test your all `_test.go` function with `Testblabla....` functions.
There are many `test` command options, like `-v` and `-coverage`, and you can easily find it on the go [site](https://golang.org/cmd/go/#hdr-Testing_flags)

### Contributors
- <a href='https://github.com/rajkumarGosavi' target='_blank'>Rajkumar</a>

[<< Home Page](./) | [Previous << Stack](./stack.html) 
