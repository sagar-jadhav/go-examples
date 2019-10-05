package main

import "fmt"

/*
 * slice program will demonstrate the working and syntax
 * of slices in the Go programming language along with use of
 * append & copy functions
 */
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
