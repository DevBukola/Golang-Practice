package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

var Port = ":8080"

type Ascii struct {
	input string
	data  string
}

func (f *Ascii) Generate() string {
	content := strings.ReplaceAll(f.data, "\r\n", "\n")
	lines := strings.Split(content, "\n")

	input := strings.ReplaceAll(f.input, "\r", "")
	words := strings.Split(input, "\n")

	var result strings.Builder

	for _, word := range words {
		if word == "" {
			result.WriteString("\n")
			continue
		}

		for i := 0; i < 8; i++ {
			for _, char := range word {
				if char < 32 || char > 126 {
					continue
				}
				ascii := int(char-32) * 9
				index := ascii + i
				if index >= 0 && index < len(lines) {
					result.WriteString(lines[index])
				} else {
					result.WriteString(" ")
				}
			}
			result.WriteString("\n")
		}
	}

	return result.String()
}

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		// http.NotFound(w, r)
		http.Error(w, "Status not found", 404)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	tmpl, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal error", http.StatusInternalServerError)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed.", http.StatusMethodNotAllowed)
		return
	}

	text := r.FormValue("text")
	banner := r.FormValue("banner")

	log.Println("TEXT:", text)
	log.Println("BANNER:", banner)
	if text == "" || banner == "" {
		http.Error(w, "Text or banner is missing - Bad request.", 400)
		return
	}

	path := "./assets/" + banner + ".txt"

	bannerContent, err := os.ReadFile(path)
	if err != nil {
		log.Println(err)
		http.Error(w, "Banner file not found", http.StatusInternalServerError)
		return
	}

	generator := &Ascii{
		input: text,
		data:  string(bannerContent),
	}

	output := generator.Generate()

	pageData := struct {
		Result string
	}{
		Result: output,
	}

	tmpl, _ := template.ParseFiles("./templates/index.html")
	// if err != nil {
	// 	log.Println(err)
	// 	http.Error(w, "Template error", http.StatusInternalServerError)
	// 	return
	// }

	err = tmpl.Execute(w, pageData)
	// if err != nil {
	// 	log.Println(err)
	// 	http.Error(w, "Error", http.StatusInternalServerError)
	// }
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/ascii-art", handler)

	fmt.Println("Server running on port", Port)
	log.Fatal(http.ListenAndServe(Port, mux))
}
