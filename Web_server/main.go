package main

import (
	"fmt"
	"log"
	"net/http"
)

type UserData struct {
	name string
	age  string
}

func createHelloPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method is not supported", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hello from Earth!")
}

func createSavedDataPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/saved_data" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "POST" {
		http.Error(w, "Method is not supported", http.StatusNotFound)
		return
	}

	err := r.ParseForm()

	if err != nil {
		log.Fatal(err)
	}

	savedData := &UserData{
		name: r.FormValue("name"),
		age:  r.FormValue("age"),
	}

	fmt.Fprintf(w, "{\n\tname: %s\n\tage: %s\n}", savedData.name, savedData.age)

}

func main() {
	// createing FileServer
	fileServer := http.FileServer(http.Dir("./static")) // setting static folder for serving
	//creating main page
	http.Handle("/", fileServer)
	// creating form page
	http.Handle("/form", fileServer)
	// creating saved data page
	http.HandleFunc("/saved_data", createSavedDataPage)
	// creating hello page
	http.HandleFunc("/hello", createHelloPage)
	// Write that server is running
	fmt.Println("Stating server at 8000 port")

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}

}
