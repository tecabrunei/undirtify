package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)


func main() {
	fmt.Println("Starting server on https://localhost:8080...")

	db, err := gorm.Open(
		"postgres", 
		"host=db user=postgres dbname=undirtify password=password sslmode=disable"
	)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	router := mux.NewRouter()
	router.HandleFunc("/", homeHandler).Methods("GET")
	router.HandleFunc("/about", aboutHandler).Methods("GET")
	router.HandleFunc("/services", servicesHandler).Methods("GET")
	router.HandleFunc("/careers", careersHandler).Methods("GET")
	router.HandleFunc("/contact", contactHandler).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}


func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to Undirtify!")
}


func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Learn more about us here.")
}


func servicesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Our services are coming soon!")
}


func careersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Check out our career opportunities.")
}


func contactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Thanks for contacting us! We'll get back to you soon.")
}
