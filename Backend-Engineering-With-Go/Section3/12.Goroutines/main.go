// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// func boilEggs() {
// 	time.Sleep(time.Second *5)
// 	fmt.Println("Boil eggs.")
// }

// func cookJollof() {
// 	time.Sleep(time.Second *4)
// 	fmt.Println("Cook jollof rice.")
// }

//  func main() {
// 	var wg sync.WaitGroup

// 	wg.Add(2)

// 	go func () {
// 		boilEggs()
// 		wg.Done()
// 	}()

// 	go func ()  {
// 		cookJollof()
// 		wg.Done()
// 	}()
// 	time.Sleep(time.Second * 6)
// 	fmt.Println("Watch movies.")
// 	wg.Wait()
// 	fmt.Println(time.Since(time.Now()))
//  }

package main

import (
	"fmt"
	"sync"
)

func doSomething(num string) {
	num = "20"
	fmt.Println(num)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func () {
		doSomething("")
		defer wg.Done()
	}()
	// go func () {
	// 	doSomething("4")
	// 	defer wg.Done()
	// }()
	wg.Wait()

	fmt.Println("Hello, there.")
}
