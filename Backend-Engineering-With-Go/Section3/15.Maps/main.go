package main

import "fmt"

func main() {
	//map literal
	//use a map literal when you already know the value when creating the map.
	var menu = map[string]int{
		"coleslaw":          500,
		"rice with chicken": 3200,
		"shawarma":          4000,
		"beef":              1000,
		"extra chicken":     1500,
	}

	// fmt.Println(menu)
	// fmt.Println(menu["extra chicken"])

	for key, value := range menu {
		fmt.Println(key, "-", value)
	}

	//use make when you want an empty ma and will fill it later.
	position := make(map[int]string)
	position[1] = "Dara"
	position[2] = "Opeyemi"
	position[3] = "Oluwatoyin"
	position[4] = "Olajide"
	position[5] = "Olawale"

	// fmt.Println(position)
	// fmt.Println(position[1])

	//example of make in action.
	names := []string{"Darasimi", "Pearl", "Daniel"}

	userName := make(map[int]string)

	for i, name := range names {
		userName[i+1] = name
	}
	// fmt.Println(userName)

	for pos, post := range userName {
		fmt.Println(pos, post)
	}
}
