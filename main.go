package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello canary world")
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "OK")
	})

	// port := os.Getenv("PORT")
	// if len(port) == 0 {
	// 	panic(port)
	// }
	http.ListenAndServe(":8080", nil)
}
