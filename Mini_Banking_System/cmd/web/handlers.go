package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"
)
	var balance = 70000.0


func transferFund(w http.ResponseWriter, r *http.Request) {
	amount := r.FormValue("transfer-amount")
	if amount == "" {
		http.Error(w, "Amount is required", 400)
		return
	}

	//I am converting the amount which is a string here to a number so I can subtract it from the balance.
	transferAmount, _ := strconv.ParseFloat(amount, 64)
	balance = balance - transferAmount;
	// fmt.Fprintf(w, "Transfer successful. New balance is: %.2f. Redirecting...", balance)
	// http.Redirect(w,r, "/", 303)
	 //statusSeeOther is for redirecting to another page after a POST request to prevent the user from resubmitting the form after refreshing the page.

	 tmpl, err := template.ParseFiles("./ui/html/pages/success.tmpl.html")
	 if err != nil {
		 log.Println(err)
		 http.Error(w, "Internal server error", 500)
		 return
	 }
	 data := struct {
		Balance float64
	}{
		Balance: balance,
	 }
	 
	err = tmpl.ExecuteTemplate(w, "success", data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", 500)
		return
	}
}
// func withdrawFund() {}
// func depositFund()  {}
// func checkBalance() {}

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
		"./ui/html/pages/partials/nav.tmpl.html",
		"./ui/html/pages/home.tmpl.html",
		"./ui/html/base.tmpl.html",
	}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", 500)
		return
	}

	today := time.Now().Format("02 January 2006")

	data := struct {
		Date    string
		Balance float64
	}{
		Date:    today,
		Balance: balance,
	}
	_ = tmpl.ExecuteTemplate(w, "base", data)
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
