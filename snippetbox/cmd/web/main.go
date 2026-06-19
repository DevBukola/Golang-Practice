package main

import (
	"log"
	"net/http"
)

// func snippetView(w http.ResponseWriter, r *http.Request) {
// 	id, err := strconv.Atoi(r.URL.Query().Get("id"))
// 	if err != nil || id < 1 {
// 		http.NotFound(w, r)
// 		return
// 	}

// 	// w.Write([]byte("Display a specific snippet...\n"))
// 	fmt.Fprintf(w, "Display a specific snippet with ID %d\n", id)
// }

// func snippetCreate(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodPost {
// 		// w.Header().Set("Allow", "POST")
// 		w.Header().Set("Allow", http.MethodPost)
// 		w.Header().Set("Cache-Control", "public, max-age=209")
// 		w.Header().Add("Cache-control", "public")
// 		w.Header().Set("Cache-Control", "public, max-age=209")
// 		// w.Header().Del("Cache-Control")
// 		cache := w.Header().Get("Cache-Control")
// 		fmt.Println("cache is:", cache)
// 		value := w.Header().Values("Cache-Control")
// 		fmt.Println("value is:", value)

// 		// w.WriteHeader(405)
// 		// w.Write([]byte("Method Not Allowed.\n"))

// 		// http.Error(w, "Method not allowed", 405)
// 		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")

// 	w.Write([]byte(`{"name": "Darasimi"}`))
// }

func main() {
	Port := ":8000"
	mux := http.NewServeMux()

	fileserver := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("/static/", http.StripPrefix("/static", fileserver))

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Println("Starting server on", Port)
	err := http.ListenAndServe(Port, mux)

	if err != nil {
		log.Fatal(err)
	}
}
