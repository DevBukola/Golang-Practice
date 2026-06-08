package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg = sync.WaitGroup{} 
	nameChan := make(chan string) 
	
	go func () {
		fmt.Println(<-nameChan)
	}()
	nameChan <- getName()
	wg.Wait()
}

func getName()string {
	 fmt.Print("Oluwabukola")
	 return ""
}