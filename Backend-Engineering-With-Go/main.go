package main

import "fmt"

func main() {
	fellows := map[int]string{
		1: "Stephen Asadu",
		2: "Enoch Adesiji",
		3: "Charlesmary Chukwuma",
	}

	delete(fellows, 3)
	fellows[4] = "Patience"
	clear(fellows)
	for _, fellow := range fellows {
		fmt.Println(fellow)
	}
}
