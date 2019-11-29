---
layout: default
---

# Slice

The example slice program will demonstrate the working and syntax of slices in the Go programming language along with use of append & copy functions

```
import "fmt"

func main() {
	//Creating Slice of length 5
	sliceA := make([]int, 5)
	sliceA[0] = 1
	sliceA[1] = 2
	sliceA[2] = 3
	print(sliceA, "sliceA")

	//Creating Slice of length 2 which is backed by array of length 5
	sliceB := make([]int, 2, 5)
	sliceB[0] = 1
	sliceB[1] = 2
	print(sliceB, "sliceB")

	//Creating Slice along with providing values to it
	slice := []int{1, 2}
	print(slice, "slice")

	//Creating Slice with low & high expressions
	array := [9]int{
		1, 2, 3,
		4, 5, 6,
		7, 8, 9,
	}

	//[low:high], Where low is the index where to start &
	//high is the index where to end but not including
	//the index itself
	sliceC := array[1:4]
	print(sliceC, "sliceC")

	//[:high]
	sliceD := array[:5]
	print(sliceD, "sliceD")

	//[low:]
	sliceE := array[0:]
	print(sliceE, "sliceE")

	//[:]
	sliceF := array[:]
	print(sliceF, "sliceF")

	//Append function
	appendedSlice := append(sliceF, 10, 11)
	print(appendedSlice, "appendedSlice")

	//Copy (Destination, Source) function
	arrayX := [6]int{
		1, 2,
		3, 4,
		5, 6,
	}

	sliceX := arrayX[:2]
	print(sliceX, "sliceX")

	sliceY := arrayX[2:5]
	print(sliceY, "sliceY")

	copiedValuesCount := copy(sliceX, sliceY)
	fmt.Printf("Number of values copied is %d", copiedValuesCount)
	fmt.Println()
	print(sliceX, "sliceX")
}

func print(arr []int, arrName string) {
	fmt.Printf("Value of Slice '%s' with length %d is %d", arrName, len(arr), arr)
	fmt.Println()
}
```
<a href='https://play.golang.org/p/TTiY5rd1h5r' target='_blank'>Try It Out</a>

[Source Code](https://github.com/sagar-jadhav/go-examples/blob/master/src/slice.go)

### Output

```
Value of Slice 'sliceA' with length 5 is [1 2 3 0 0]
Value of Slice 'sliceB' with length 2 is [1 2]
Value of Slice 'slice' with length 2 is [1 2]
Value of Slice 'sliceC' with length 3 is [2 3 4]
Value of Slice 'sliceD' with length 5 is [1 2 3 4 5]
Value of Slice 'sliceE' with length 9 is [1 2 3 4 5 6 7 8 9]
Value of Slice 'sliceF' with length 9 is [1 2 3 4 5 6 7 8 9]
Value of Slice 'appendedSlice' with length 11 is [1 2 3 4 5 6 7 8 9 10 11]
Value of Slice 'sliceX' with length 2 is [1 2]
Value of Slice 'sliceY' with length 3 is [3 4 5]
Number of values copied is 2
Value of Slice 'sliceX' with length 2 is [3 4]
```

[Back](./)
