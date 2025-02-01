package main

import "fmt"

type engine struct{
	mpg uint
	gallons uint 
	owner
}

func ( e engine ) milesLeft() uint8 {
	return uint8(e.gallons) * uint8(e.mpg)
}

type owner struct {
	name string
}

func main(){

	var myEngine engine = engine{ mpg:10 , gallons:20 };
	myEngine.mpg = 20;
	var user owner = owner{"akshat"}
	myEngine.owner = user
	fmt.Println((myEngine))

	fmt.Println(myEngine.milesLeft())

}