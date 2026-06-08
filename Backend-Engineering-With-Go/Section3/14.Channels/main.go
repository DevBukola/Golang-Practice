package main

import (
	"fmt"
	"math"
)

func main() {
	numbers := [5]int{1, 2, 3, 4, 5}
	numberChan := make(chan int)

	go func() {
		//this loop goes through the array and sends each number into the numberChan. Without it, I'd have to do:
		// numberChan <- 1
		// numberChan <- 2
		// numberChan <- 3
		// numberChan <- 4
		// numberChan <- 5
		for _, num := range numbers {
			numberChan <- num
		}
	}()

	// this loop is receiving value from the channel. Without it I'd have to do:
	//fmt.Println(<-numberChan)
	//fmt.Println(<-numberChan)
	//fmt.Println(<-numberChan)
	//fmt.Println(<-numberChan)
	//fmt.Println(<-numberChan)
	for i := 0; i < len(numbers); i++ {
		number := <-numberChan
		fmt.Println("num is:",number)

		//squaring the numbers:
		square := math.Pow(float64(number), 2)
		fmt.Println("square is:", square)
		// cube := math.Pow(float64(number), 3)
		cube := number * number * number
		fmt.Println("cube is:", cube)
	}
}
