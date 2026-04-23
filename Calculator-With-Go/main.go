package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	for {
	var firstNumber, secondNumber string
	var operator string

	fmt.Print("Kindly enter the first number or type 'exit' to quit: ")
	fmt.Scanln(&firstNumber)

	if firstNumber == "exit" {
			fmt.Println("Ouch, it's sad to see you go.😢 Goodbye!👋")
			break
	}

	// fmt.Println(firstNumber,n,err)

	firstInput, err:= strconv.ParseFloat(firstNumber, 64)
	if err != nil {
		fmt.Printf("%s is not a valid number.\n", firstNumber)
		return
	}

	fmt.Print("Thank you. Now, enter the second number: ")
	fmt.Scanln(&secondNumber)
	
	secondInput, err:= strconv.ParseFloat(secondNumber, 64)
	if err != nil {
		fmt.Printf("%s is an invalid number.", secondNumber)
		return
	}

	fmt.Print("Please enter the operator (+, -, *, /, %): ")
	fmt.Scanln(&operator)

	switch operator {
	case "+" :
		fmt.Printf("%.0f %s %.0f = %.0f\n", firstInput, operator, secondInput, firstInput + secondInput)
	case "-" :
		fmt.Printf("%.0f %s %.0f = %.0f\n", firstInput, operator, secondInput, firstInput - secondInput)
	case "*" :
		fmt.Printf("%.0f %s %.0f = %.0f\n", firstInput, operator, secondInput, firstInput * secondInput)
	case "/" :
		if secondInput == 0.0 {
			fmt.Println("You've divided by zero.")
			// return
		}else {
			fmt.Printf("%.0f %s %.0f = %.0f\n", firstInput, operator, secondInput, firstInput / secondInput)
		}
	case "%" :
		fmt.Printf("%.0f %s %.0f = %.0f\n", firstInput, operator, secondInput, math.Mod(firstInput, secondInput))
	default: 
		fmt.Println("You've entered an invalid operator.")
	}
}
}