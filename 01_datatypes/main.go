package main

import "fmt"

func main() {

	var num int = 5
	var floatNum float64 = 1234.5678

	fmt.Println(num, floatNum)

	var floatMultipled float64 = floatNum * float64(num)
	fmt.Println("In order to do operations on mixed types we need to cast", floatMultipled)

	var myString string = "Hello world! \n next line"
	fmt.Println(myString)

	var myBool bool = true
	fmt.Println(myBool)

	var count int
	fmt.Println("default value",count)

	myVar := "hey"
	fmt.Println("dropping the var keyword and using := ", myVar)

	var1 , var2 := 1 , 2
	fmt.Println(var1 , var2)

	const constString string =  "constant string"
	// constString = "can't change const"
	fmt.Println(constString)

}
