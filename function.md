
---

layout: default

---

  

# Function

  

The function encloses several sentences. In Go, the function is func. func but I wrote down the function name and passed parentheses `()` and that function. The performance of ints, strings, etc. has been improved. The return value of the function is parentheses `()`. You can define a function as a package.

  

The example below is a function called say. Replace the say `()` function.

  

```

package main

func main() {

	msg := "Hello"

	say(msg)

}

func say(msg string) {

	println(msg)

}

```

  

# 2.pass by reference / value

  

The method of passing parameters in Go is largely divided into Pass By Value and Pass By Reference.

  

### Pass by value

Above [1. Function], the string "Hello" of msg is copied and passed to the function say (). That is, even if the value of the parameter msg is changed in the say () function, the msg variable in the calling function main () is not changed.

  

### Pass by reference

As in the example below, prefixing the msg variable with & indicates the address of the msg variable. This technique, commonly referred to as a pointer, lets you pass the address of a msg variable to a function without copying the value of the msg variable. The called function say () indicates that the parameter is a pointer, such as `*` string. In this case, the msg of the say function has the address of the memory area that contains the string, not the string. To write data to msg addresses, prefix `*` with `msg = ""`, which is commonly referred to as dereferencing.

In the example below, the value of the msg variable in the main function has been changed from Change in the say function, so the changed value is printed in println () of the main function. The return value of the function is parentheses (). You can define a function as a package.

  

The example below is a function called say. Replace the say () function.

  

```

package main

func main() {

	msg := "Hello"

	say(&msg)

	println(msg) // print the message has been changed

}

func say(msg *string) {
	println(*msg)
	*msg = "Changed" // Change a message

```

  

# 3. Variadic Function

When you want to pass a variable number of parameters without passing a fixed number of parameters to the function, use the ... (three periods) to represent the variable parameters. That is, to represent a string variable parameter, it is expressed as ` ... ` string. When calling a function with variable parameters, you can pass zero, one, two, or n identical type parameters. This function that accepts a variable parameter is called a Variadic Function.

The example below shows that you can pass four strings to the say function and one string to pass.

```
package main

func main() {

	say("This", "is", "a", "book")

	say("Hi")

}

func say(msg ...string) {

	for _, s := range msg {

		println(s)

	}

}
```
 
# 4. Function return value

In the Go programming language, a function may have no return value, one return value, or multiple return values. In contrast to returning void or only one value from the C language, the Go language can return multiple values.
The Go language also provides a feature called Named Return Parameter, which allows you to assign returned values ​​to return parameters (defined in a function).

If there is a return value from the function, the type of the return value is defined at the end of the func statement (after the parentheses). Then use the return keyword in the function to return a value. For example, the following example shows that the return type of the sum () function is int, and it returns the value of the integer s at the end of the sum function, such as return s.

```
package main
 
func main() {
    total := sum(1, 7, 3, 5, 9)
    println(total)
}
 
func sum(nums ...int) int {
    s := 0
    for _, n := range nums {
        s += n
    }
    return s
}

```

Even, you can also return several values. you can use like `function(args) (int, string)`

```
package main
 
func main() {
    count, total := sum(1, 7, 3, 5, 9)
    println(count, total)   
}
 
func sum(nums ...int) (int, int) {
    s := 0    
    count := 0  
    for _, n := range nums {
        s += n
        count++
    }
    return count, s
}
```


[Back](./)
