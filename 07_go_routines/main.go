package main

import (
	"fmt"
	"sync"
	"time"
)

var m = sync.RWMutex{}
var wg = sync.WaitGroup{}
var dbData = []string{"id1","id2","id3","id4","id5"}
var results = []string{}

func main(){
	t0 := time.Now()
	for i :=0 ;  i < len(dbData) ; i++ {
		wg.Add(1)
		go dbCall(i)
	}
	wg.Wait()
	fmt.Println("\nTotal Execution time: " , time.Since(t0))
	fmt.Println("The results are :", results)
}

func dbCall(i int){
	// simulate db call delay
	var delay float32 = 2000 
	time.Sleep(time.Duration(delay)*time.Millisecond)
	// fmt.Println("The result from database is: ", dbData[i])

	save(dbData[i])
	log()

	wg.Done()
}

func save(result string) {
	m.Lock()
	// if we don't lock the write operation then the append function might 
	//try to overwrite the change due to picking previous length of slice
	results = append(results, result)
	m.Unlock()
}

func log( ){
	m.RLock()
	fmt.Println("The current results are:", results)
	m.RUnlock()
}