package main

import "fmt"

func main() {

	// arraySlicesBasics()
	mapsBasics()

}

func arraySlicesBasics() {

	// good array / slice resource --->>> 	https://dev.to/dawkaka/go-arrays-and-slices-a-deep-dive-dp8

	var intArr [3]int32
	intArr[1] = 123

	fmt.Println(intArr)
	fmt.Println((intArr[0:1]))

	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])

	// can also be initialized like
	var intArr2 [3]int = [3]int{1, 2, 3}
	fmt.Println(intArr2)

	intArr3 := []int{1, 2, 3, 4}
	fmt.Println(intArr3)

	intArr4 := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(intArr4)

	stringArr := []string{"abc", "def", "gh"}
	fmt.Println(stringArr)

	stringArr = append(stringArr, "xyz")
	fmt.Println(len(stringArr), cap(stringArr))

	// key note -> while appending to a slice if there is not engough capacity in slice , a new slice is made
	// good practice to follow is to know the estimate capacity in order to
	// prevent performance drop due to copying data to new memory location

	intArr3 = append(intArr3, intArr4...)
	fmt.Println(intArr3)

	var intArr5 [3][5]int
	fmt.Println(intArr5)

	// slice can also be created using make function

	// With capacity
	slice1 := make([]int, 5, 10)
	fmt.Println(len(slice1), cap(slice1)) // => 5 10

	// Without capacity
	slice2 := make([]int, 5)
	fmt.Println(len(slice2), cap(slice2)) // => 5 5

	for i , v := range intArr3 {
		fmt.Printf("index is %v , and value is %v \n", i , v)
	}

}
