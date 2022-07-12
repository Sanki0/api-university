package main

import (
	"github.com/Sanki0/api-university/handlers"

	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
}

func main() {

	r := mux.NewRouter()

	//CREATE STUDENT
	r.HandleFunc("/student", handlers.CreatePage).Methods("POST")

	//READ ALL STUDENTS
	r.HandleFunc("/students", handlers.StudentPage).Methods("GET")

	//READ SINGLE STUDENT
	r.HandleFunc("/student", handlers.SingleStudentPage).Methods("GET")

	//UPDATE STUDENT
	r.HandleFunc("/student", handlers.UpdatePage).Methods("PUT")
	//DELETE STUDENT
	r.HandleFunc("/student", handlers.DeleteStudentPage).Methods("DELETE")

	//HOME PAGE

	r.HandleFunc("/", homePage)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
