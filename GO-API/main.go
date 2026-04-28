package main

import (
	// "encoding/json"
	"fmt"
	// "io"
	// "net/http"
)

var URL = "https://v2.jokeapi.dev/joke/any"
type Detail struct {
	Category string
	Type string
	Id int
	Safe bool
	Flag Flags
	
}

type Flags struct {
	Religious bool
	Explicit bool
}

type Fellow struct {
	Gender string
	Religion string
}
// type fellows struct {
// 	F []Fellow
// }

func main() {

	// res, _ := http.Get(URL)
	// response, _ := io.ReadAll(res.Body)
	// defer res.Body.Close()
	// fmt.Println(string(response))

	// var res1 Detail
	// json.Unmarshal(response, &res1)
	// fmt.Printf("%+v", res1)
	// fmt.Printf("%+v\n",res1.Flag.Religious)
	// fmt.Printf("%+v\n",res1.Flag.Explicit)


	var fellows []Fellow 
	
	// {
	// 	{
	// 		Religion: "Muslim",
	// 		Gender: "Male",
	// 	},
	// 	{
	// 		Religion: "",
	// 		Gender: "Female",
	// 	},Christian
	// }

	fellows = append(fellows, generateFellow("Muslim","Male"))
	fellows = append(fellows, generateFellow("Christian","Female"))

	// for _, fellow := range fellows{
	// 	fmt.Println(fellow.Religion)
	// }

	fmt.Println(fellows)
}

func generateFellow(r,g string) Fellow{
	return Fellow{
			Religion: r,
			Gender: g,
		}
}
