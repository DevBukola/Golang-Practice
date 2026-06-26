package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"
)
// 	var balance = 70000.0

type User struct {
  ID int
  Username string
  Password string
  AccountNumber string
  Balance float64
}

var Users = map[string]*User {
  "devbukola": {
    ID: 1,
    Username: "devbukola",
    Password: "1234",
    AccountNumber: "8103664531",
    Balance: 500000.0,
  },
  "oudo": {
     ID: 2,
    Username: "oudo",
    Password: "4567",
    AccountNumber: "1111111111",
    Balance: 602000.0,
  },
  "sasadu": {
     ID: 3,
    Username: "sasadu",
    Password: "1111",
    AccountNumber: "2222222222",
    Balance: 522420.0,
  },
  "daniel": {
     ID: 4,
    Username: "daniel",
    Password: "2222",
    AccountNumber: "3333333333",
    Balance: 5784955.0,
  },
}

func transferFund(w http.ResponseWriter, r *http.Request) {
  cookie, err := r.Cookie("session")
  if err != nil {
    http.Redirect(w, r, "/login", http.StatusSeeOther)
    return
}
username := cookie.Value
user, valid := Users[username]
	if !valid {
	  http.Redirect(w, r, "/login", 303)
	  return
	}
	amount := r.FormValue("transfer-amount")
	if amount == "" {
		http.Error(w, "Amount is required", 400)
		return
	}

	//I am converting the amount which is a string here to a number so I can subtract it from the balance.
	transferAmount, _ := strconv.ParseFloat(amount, 64)
	user.Balance = user.Balance - transferAmount;
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
		Balance: user.Balance,
	 }
	 
	err = tmpl.ExecuteTemplate(w, "success", data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", 500)
		return
	}
}
func depositFund(w http.ResponseWriter, r *http.Request) {
  cookie, err := r.Cookie("session")
  if err != nil {
    http.Redirect(w, r, "/login", http.StatusSeeOther)
    return
}
username := cookie.Value
user, valid := Users[username]
	if !valid {
	  http.Redirect(w, r, "/login", 303)
	  return
	}
  amount := r.FormValue("deposit-amount")
  if amount == "" {
    http.Error(w, "Bad request", 400)
    return
  }
  depositAmount, _ := strconv.ParseFloat(amount, 64)
  user.Balance = user.Balance + depositAmount
  
  tmpl, _ := template.ParseFiles("./ui/html/pages/successDepo.tmpl.html")
  data := struct {
    Balance float64
  }{
    Balance: user.Balance,
  }
  tmpl.ExecuteTemplate(w, "successDepo", data)
}

// func depositFund()  {}
// func checkBalance() {}

func homeHandler(w http.ResponseWriter, r *http.Request) {
  cookie, err := r.Cookie("session")
  if err != nil {
    http.Redirect(w, r, "/login", http.StatusSeeOther)
    return
}
username := cookie.Value
user, valid := Users[username]
	if !valid {
	  http.Redirect(w, r, "/login", 303)
	  return
	}
	if r.URL.Path != "/" {
		http.Error(w, "Page not found", 404)
		return
	}
// 	if r.Method != "GET" {
// 		http.Error(w, "Method not allowed", 405)
// 		return
// 	}
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
		Username string
		AccountNumber string
	}{
		Date:    today,
		Balance: user.Balance,
		Username: user.Username,
		AccountNumber: user.AccountNumber,
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
	username := r.FormValue("username")
	password := r.FormValue("password")
	
	user, valid := Users[username]
	if !valid {
	  http.Error(w, "Username or Password not valid.", http.StatusUnauthorized)
	  return
	}
	if user.Password != password {
	  http.Error(w, "Password is incorrect", http.StatusUnauthorized)
	  return
	}
	cookie := &http.Cookie{
	Name:  "session",
	Value: user.Username,
	Path:  "/",
}

http.SetCookie(w, cookie)
http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	
func logOutHandler(w http.ResponseWriter, r *http.Request) {
  cookie := &http.Cookie {
    Name: "session",
    Value: "",
    Path: "/",
  }
  http.SetCookie(w, cookie)
  http.Redirect(w, r, "/login", 303)
}
