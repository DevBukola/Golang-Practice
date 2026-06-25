package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func transferFund() {}
func withdrawFund() {}
func depositFund()  {}
func checkBalance() {}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Page not found", 404)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}
	files := []string{
		"./ui/html/pages/home.tmpl.html",
		"./ui/html/base.tmpl.html",
	}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", 500)
		return
	}
	_ = tmpl.ExecuteTemplate(w, "base", nil)
}

func signupPageRendering(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Welcome to signup page!")
	filepath := []string{
		"./ui/html/pages/signup.tmpl.html",
	}

	tmpl, err := template.ParseFiles(filepath...)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", 500)
		return
	}
	err = tmpl.ExecuteTemplate(w, "signup", nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", 500)
		return
	}
}

func createAccountHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Bad request", 400)
		return
	}
	fmt.Fprintf(w, "Hello")
}

func signInPageRendering(w http.ResponseWriter, r *http.Request) {
	filepath := "./ui/html/pages/login.tmpl.html"

	tmpl, err := template.ParseFiles(filepath)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", 500)
		return
	}

	err = tmpl.ExecuteTemplate(w, "login", nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", 500)
		return
	}
}

func signInHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Bad request", 400)
		return
	}
	fmt.Fprintf(w,"Welcome to the sign in page")
}
