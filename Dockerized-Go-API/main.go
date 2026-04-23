package main

import (
	// "encoding/json"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
)

// var quotes = []string {
// 	"Quote A",
// 	"Quote B",
// 	"Quote C",
// 	"Quote D",
// 	"Quote C",

// }
type Quote struct {
	Quote string
	Author string 
	Experience int `json:"Years of Experience"`
}

type ResponseHealth struct {
	Status string `json:"status"`
	Service string `json:"service"`
	
}
var quotes = []Quote{
	{
	Quote: "The best time to start was yesterday. The next best time is now.",
	Author: "Karen Lamb",
	Experience: 12,
},
{
	Quote: "Success is not final, failure is not fatal: it is the courage to continue that counts.",
	Author: "Winston Churchill",
	Experience: 40,
},
{
	Quote: "Do what you can, with what you have, where you are.",
	Author: "Theodore Roosevelt",
	Experience: 35,
},
{
	Quote: "Hardships often prepare ordinary people for an extraordinary destiny.",
	Author: "C.S. Lewis",
	Experience: 25,
},
{
	Quote: "It always seems impossible until it's done.",
	Author: "Nelson Mandela",
	Experience: 50,
},
{
	Quote: "Don't watch the clock; do what it does. Keep going.",
	Author: "Sam Levenson",
	Experience: 20,
},
{
	Quote: "Your time is limited, so don't waste it living someone else's life.",
	Author: "Steve Jobs",
	Experience: 30,
},
{
	Quote: "Discipline is the bridge between goals and accomplishment.",
	Author: "Jim Rohn",
	Experience: 22,
},
{
	Quote: "Act as if what you do makes a difference. It does.",
	Author: "William James",
	Experience: 28,
},
{
	Quote: "Believe you can and you're halfway there.",
	Author: "Theodore Roosevelt",
	Experience: 35,
},
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Better person, welcome to my Go Quote API!😊✈️")
}

func handleQuote(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, quotes[0])
	randomQuoteIdx := rand.Intn(len(quotes))
	randomQuoteValue := quotes[randomQuoteIdx]
	// fmt.Fprintln(w, randomQuoteValue)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(randomQuoteValue)

	// w.Header().Set("content-type","application/json")
	// json.NewEncoder(w).Encode(data)
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
w.Header().Set("Content-Type", "application/json")
if len(quotes) == 0 {
	json.NewEncoder(w).Encode(map[string]string{
		"Status": "Error! There is no quote available.",
	})
} else {
json.NewEncoder(w).Encode(&ResponseHealth{
	Status: "ok!",
	Service: "go-quote-api",
})
}
}

func main() {
	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/quote", handleQuote)
	http.HandleFunc("/health", handleHealth)

	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}