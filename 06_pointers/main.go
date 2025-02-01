package main
import "fmt"

func main(){

	var p *int32
	var i int32

	p = &i
	i = 32

	fmt.Println(i , p , &i)

	*p = 10

	fmt.Println(i , *p , &i)

}