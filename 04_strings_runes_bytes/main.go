package main

import (
	"fmt"
	"strings"
)

func main(){

	str :=  []rune("résumé")
	fmt.Println(len(str))
	fmt.Println(len("résumé"))
	// using rune gives us correct length of characters when dealing with special characters 


	// golang stores strings as a array of bytes ( utf-8 encoding ) 
	// strings are immutable in go which is inefficient when concatinating ( we should used string builder package)

	strSlice := []string{"s","t","r","i","n","g"}

	var strBuilder	strings.Builder
	
	for i:= range strSlice{
		strBuilder.WriteString(strSlice[i])
	}
	var concatStr = strBuilder.String()

	fmt.Println(concatStr)


}