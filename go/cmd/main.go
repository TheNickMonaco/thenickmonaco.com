package main

import (
	"fmt"
	"net/http"
	"html/template"
)

func main() {
	http.HandleFunc("/", homeHandler)

	// Start the server on port 8080
	fmt.Println("Server started on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// Your home page or any response for htmx interaction
	tmpl, err := template.ParseFiles("htmx/html/home.html")
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}
