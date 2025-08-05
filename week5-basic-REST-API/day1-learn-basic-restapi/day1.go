package main

import (
	f "fmt"
	http "net/http"
	"log"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	f.Fprint(w, "This is Home Page")
}

func Hello(w http.ResponseWriter, r *http.Request) {
	f.Fprint(w, "Hello Quan. How are you?")
}

func Goodbye(w http.ResponseWriter, r *http.Request) {
	f.Fprint(w, "Bye Quan. See you again?")
}

func HandleHomePage() {
	f.Println("HomePage is listening on port 3101\nAccess to http://localhost:3101/")
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/hello", Hello)
	http.HandleFunc("/goodbye", Goodbye)
	log.Fatal(http.ListenAndServe(":3101", nil))
}


func main() {
	HandleHomePage()
}