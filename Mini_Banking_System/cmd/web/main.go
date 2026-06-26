package main

import (
	"fmt"
	"net/http"
)

var Port = ":3000"

func main() {

	mux := http.NewServeMux()

	fileserver := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileserver))

	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/signup", signupPageRendering)
	mux.HandleFunc("/logout", logOutHandler)
	mux.HandleFunc("/login", signInPageRendering)
	mux.HandleFunc("/sign-in", signInHandler)
	mux.HandleFunc("/transfer", transferFund)
	mux.HandleFunc("/deposit", depositFund)

	fmt.Println("Server running on port", Port)
	http.ListenAndServe(Port, mux)
}
