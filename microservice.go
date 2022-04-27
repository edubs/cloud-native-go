package main

import (
	"cloud-native-go/api"
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/api/echo", api.Echo)
	http.HandleFunc("/api/books", api.BooksHandleFunc)

	// Paths below this call will fail
	http.ListenAndServe(port(), nil)

}

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	return ":" + port

}

func index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello Cloud Native Go from %v.", port())
}
