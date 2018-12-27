package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>hey!</h1>")
}

func main() {
	http.HandleFunc("/", index)
	fmt.Println("server is running")
	http.ListenAndServe(":3001", nil)

}
