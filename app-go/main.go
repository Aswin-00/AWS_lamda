package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World! You've hit %s\n", r.URL.Path)
}

func main() {
	http.HandleFunc("/", handler)

	port := "8080"
	fmt.Println("Server is running on port " + port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
