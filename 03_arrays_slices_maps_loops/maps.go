package main

import "fmt"

func mapsBasics() {
	
	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"akshat":23, "name":18}
	fmt.Println(myMap2)
 
	var age = myMap2["akshat"]
	fmt.Println(age)

	// if a key doen't exists then the map return a default value of value type
	var age2 , ok = myMap2["codespicer"]
	if( ok){
		fmt.Println(age2)
	}else{
		fmt.Printf("age doesn't exit for key %v \n", "codespicer")
	}

	for key , value := range myMap2{
		fmt.Println(key , value)
	}

	
}
