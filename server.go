package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!\n")
}

func helloTwo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi\n")
}

func main() {

	link := http.FileServer(http.Dir("./static"))

	http.Handle("/", link)
	http.HandleFunc("/Hi", helloTwo)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
